/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pool

import (
	"fmt"
	math_rand "math/rand"
	"sync"

	"os"
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/tools"
	"strings"
)

// A PoolAllocator is a helper interface that supports a pool of resources that can be allocated
// TODO: Should an allocation with the same owner succeed?  Probably yes...
type PoolAllocator interface {
	// Allocate allocates a specific entry.
	Allocate(key string, owner string) (bool, error)

	// AllocateNext allocates and returns a new entry.
	AllocateNext(owner string) (string, error)

	// Release de-allocates an entry.
	// TODO: We could take an owner, for safety
	Release(key string) (bool, error)

	// List all allocations, for repair.
	ListAllocations() ([]Allocation, error)

	// Checks if key is locked, for repair.
	ReadAllocation(key string) (*Allocation, uint64, error)

	// Releases allocation, if still current, for repair.
	ReleaseForRepair(allocation Allocation) (bool, error)

	// For tests
	DisableRandomAllocation()
}

// A very simple string iterator
type StringIterator interface {
	Next() (string, bool)
}

type Allocation struct {
	Key     string
	Owner   string
	Version uint64
}

// A PoolDriver provides the PoolAllocator with the set from which it allocates
type PoolDriver interface {
	// Choose an item from the set
	PickRandom(random *math_rand.Rand) string

	// Iterate across all items in the set
	Iterate() StringIterator
}

// MemoryPoolAllocator is PoolAllocator that is backed by an in-memory collection,
// can't be used if there are multiple processes allocating from the pool (cf EtcdPoolAllocator)
type MemoryPoolAllocator struct {
	driver PoolDriver

	lock sync.Mutex // protects 'used' and 'random'

	used           map[string]Allocation
	randomAttempts int

	random *math_rand.Rand
}

// EtcdPoolAllocator is PoolAllocator that is backed by etcd, so can be used by multiple processes
type EtcdPoolAllocator struct {
	driver PoolDriver

	lock sync.Mutex // protects 'used' and 'random'

	usedCached     map[string]bool
	randomAttempts int

	random *math_rand.Rand

	etcd    *tools.EtcdHelper
	baseKey string
}

// For tests
func (a *MemoryPoolAllocator) DisableRandomAllocation() {
	a.randomAttempts = 0
}

// For tests
func (a *EtcdPoolAllocator) DisableRandomAllocation() {
	a.randomAttempts = 0
}

// Init initializes a MemoryPoolAllocator.
func (a *MemoryPoolAllocator) Init(driver PoolDriver) {
	seed := time.Now().UTC().UnixNano()
	a.random = math_rand.New(math_rand.NewSource(seed))

	a.randomAttempts = 1000
	a.used = make(map[string]Allocation)

	a.driver = driver
}

// Init initializes a EtcdPoolAllocator.
func (a *EtcdPoolAllocator) Init(driver PoolDriver, etcd *tools.EtcdHelper, baseKey string) {
	seed := time.Now().UTC().UnixNano()
	seed ^= int64(os.Getpid())
	a.random = math_rand.New(math_rand.NewSource(seed))

	a.randomAttempts = 1000
	a.usedCached = make(map[string]bool)

	a.driver = driver

	a.etcd = etcd
	if !strings.HasSuffix(baseKey, "/") {
		baseKey += "/"
	}
	a.baseKey = baseKey
}

// Allocate allocates a specific entry.
func (a *MemoryPoolAllocator) Allocate(key string, owner string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	allocation, found := a.used[key]
	if found {
		return false, nil
	}
	allocation.Owner = owner
	allocation.Key = key
	a.used[key] = allocation
	return true, nil
}

// Try to lock the resource identified by s.
// Checks to see if we have a cached allocation first if useCache is true.
// Always updates the cache when it learns information.
func (a *EtcdPoolAllocator) tryLock(key string, owner string, useCache bool) (bool, error) {
	if useCache {
		a.lock.Lock()
		usedCached := a.usedCached[key]
		a.lock.Unlock()

		if usedCached {
			return false, nil
		}
	}

	etcdKey := a.baseKey + key
	nodeValue := owner

	_, err := a.etcd.Client.Create(etcdKey, nodeValue, 0)

	created := true
	if err != nil {
		if tools.IsEtcdNodeExist(err) {
			// We failed to obtain the lock
			created = false
		} else {
			return false, fmt.Errorf("Error communicating with etcd: %v", err)
		}
	}

	// Whether or not we locked it, it is locked
	a.lock.Lock()
	a.usedCached[key] = true
	a.lock.Unlock()

	return created, nil
}

// Allocate allocates a specific entry.
func (a *EtcdPoolAllocator) Allocate(key string, owner string) (bool, error) {
	useCache := false
	locked, err := a.tryLock(key, owner, useCache)
	if err != nil {
		return false, err
	}
	return locked, nil
}

// AllocateNext allocates and returns a new entry.
func (a *MemoryPoolAllocator) AllocateNext(owner string) (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	// Try randomly first
	for i := 0; i < a.randomAttempts; i++ {
		s := a.driver.PickRandom(a.random)

		allocation, found := a.used[s]
		if !found {
			allocation.Owner = owner
			allocation.Key = s
			a.used[s] = allocation
			return s, nil
		}
	}

	// If that doesn't work, try a linear search
	iterator := a.driver.Iterate()
	for {
		s, found := iterator.Next()
		if !found {
			break
		}
		allocation, found := a.used[s]
		if !found {
			allocation.Owner = owner
			allocation.Key = s
			a.used[s] = allocation
			return s, nil
		}
	}
	return "", nil
}

// AllocateNext allocates and returns a new entry.
func (a *EtcdPoolAllocator) AllocateNext(owner string) (string, error) {
	// Try randomly first
	for _, useCache := range []bool{true, false} {
		for i := 0; i < a.randomAttempts; i++ {
			// random is not documented to be thread-safe
			a.lock.Lock()
			s := a.driver.PickRandom(a.random)
			a.lock.Unlock()

			locked, err := a.tryLock(s, owner, useCache)
			if err != nil {
				return "", err
			}
			if locked {
				return s, nil
			}
		}
	}

	// If that doesn't work, try a linear search
	useCache := false
	iterator := a.driver.Iterate()
	for {
		s, found := iterator.Next()
		if !found {
			break
		}
		locked, err := a.tryLock(s, owner, useCache)
		if err != nil {
			return "", err
		}
		if locked {
			return s, nil
		}
	}
	return "", nil
}

// Release de-allocates an entry.
func (a *MemoryPoolAllocator) Release(key string) (bool, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	_, found := a.used[key]
	if !found {
		return false, nil
	}
	delete(a.used, key)
	return true, nil
}

// Release de-allocates an entry.
func (a *EtcdPoolAllocator) Release(key string) (bool, error) {
	etcdKey := a.baseKey + key

	recursive := false
	_, err := a.etcd.Client.Delete(etcdKey, recursive)

	deleted := true
	if err != nil {
		if tools.IsEtcdNotFound(err) {
			deleted = false
		} else {
			return false, fmt.Errorf("Error communicating with etcd: %v", err)
		}
	}

	// Whether or not we deleted it, it is deleted
	a.lock.Lock()
	a.usedCached[key] = false
	a.lock.Unlock()

	return deleted, nil
}

// Count the number of items allocated.  Intended for tests, because an etcd implementation is likely to be slow.
func (a *MemoryPoolAllocator) size() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	n := len(a.used)
	return n
}

// Checks if an entry is allocated.
func (a *MemoryPoolAllocator) ReadAllocation(s string) (*Allocation, uint64, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	allocation, found := a.used[s]
	if !found {
		return nil, 0, nil
	}

	return &allocation, 0, nil
}

// Checks if an entry is allocated.
func (a *EtcdPoolAllocator) ReadAllocation(s string) (*Allocation, uint64, error) {
	etcdKey := a.baseKey + s

	allocations, t, err := a.listAllocations(etcdKey)

	var allocation *Allocation
	if err == nil {
		a.lock.Lock()

		if len(allocations) == 0 {
			a.usedCached[s] = false
		} else if len(allocations) == 1 {
			a.usedCached[s] = true
			allocation = &allocations[0]
		}

		a.lock.Unlock()
	}

	return allocation, t, err

}

// Release an entry, if still current.
func (a *MemoryPoolAllocator) ReleaseForRepair(allocation Allocation) (bool, error) {
	// We shouldn't need a repair for an in-memory implementation
	panic("Repair functions not implemented for MemoryPoolAllocator")
}

// Release an entry, if still current.
func (a *EtcdPoolAllocator) ReleaseForRepair(allocation Allocation) (bool, error) {
	etcdKey := a.baseKey + allocation.Key

	nodeValue := allocation.Owner
	_, err := a.etcd.Client.CompareAndDelete(etcdKey, nodeValue, allocation.Version)

	deleted := true
	newUsedCached := false
	if err != nil {
		if tools.IsEtcdNotFound(err) {
			deleted = false
			newUsedCached = false
		} else if tools.IsEtcdTestFailed(err) {
			deleted = false
			newUsedCached = true
		} else {
			return false, fmt.Errorf("Error communicating with etcd: %v", err)
		}
	}

	// Whether or not we deleted it, it is deleted
	a.lock.Lock()
	a.usedCached[allocation.Key] = newUsedCached
	a.lock.Unlock()

	return deleted, nil
}

// List all allocations
func (a *MemoryPoolAllocator) ListAllocations() ([]Allocation, error) {
	panic("Not implemented")
}

// List all allocations
func (a *EtcdPoolAllocator) ListAllocations() ([]Allocation, error) {
	etcdKey := a.baseKey

	allocations, _, err := a.listAllocations(etcdKey)

	if err == nil {
		a.lock.Lock()
		a.usedCached = map[string]bool{}
		for i := range allocations {
			a.usedCached[allocations[i].Key] = true
		}
		a.lock.Unlock()
	}
	return allocations, err
}

// List all allocations
func (a *EtcdPoolAllocator) listAllocations(etcdKey string) ([]Allocation, uint64, error) {
	sort := false
	recursive := false
	response, err := a.etcd.Client.Get(etcdKey, sort, recursive)

	if err != nil {
		if tools.IsEtcdNotFound(err) {
			return nil, response.EtcdIndex, nil
		}

		return nil, 0, fmt.Errorf("Error communicating with etcd: %v", err)
	}

	allocations := []Allocation{}

	keyPrefix := response.Node.Key + "/"
	for _, node := range response.Node.Nodes {
		var allocation Allocation
		allocation.Key = node.Key[len(keyPrefix):]
		allocation.Owner = node.Value
		allocation.Version = node.ModifiedIndex
		allocations = append(allocations, allocation)
	}

	return allocations, response.EtcdIndex, nil
}
