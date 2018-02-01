/*
Copyright 2017 The Kubernetes Authors.

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

package e2e_node

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/uuid"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig"
	"k8s.io/kubernetes/test/e2e/framework"

	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha"
	dm "k8s.io/kubernetes/pkg/kubelet/cm/devicemanager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	// fake resource name
	resourceName = "fake.com/resource"
)

// Serial because the test restarts Kubelet
var _ = framework.KubeDescribe("Device Plugin [Feature:DevicePlugin] [Serial] [Disruptive]", func() {
	f := framework.NewDefaultFramework("device-plugin-errors")

	Context("DevicePlugin", func() {
		By("Enabling support for Device Plugin")
		tempSetCurrentKubeletConfig(f, func(initialConfig *kubeletconfig.KubeletConfiguration) {
			initialConfig.FeatureGates[string(features.DevicePlugins)] = true
		})

		It("Verifies the Kubelet device plugin functionality.", func() {

			By("Wait for node is ready")
			framework.WaitForAllNodesSchedulable(f.ClientSet, framework.TestContext.NodeSchedulableTimeout)

			By("Start stub device plugin")
			// fake devices for e2e test
			devs := []*pluginapi.Device{
				{ID: "Dev-1", Health: pluginapi.Healthy},
				{ID: "Dev-2", Health: pluginapi.Healthy},
			}

			socketPath := pluginapi.DevicePluginPath + "dp." + fmt.Sprintf("%d", time.Now().Unix())

			dp1 := dm.NewDevicePluginStub(devs, socketPath)
			dp1.SetAllocFunc(stubAllocFunc)
			err := dp1.Start()
			framework.ExpectNoError(err)

			By("Register resources")
			err = dp1.Register(pluginapi.KubeletSocket, resourceName)
			framework.ExpectNoError(err)

			By("Waiting for the resource exported by the stub device plugin to become available on the local node")
			devsLen := int64(len(devs))
			Eventually(func() int64 {
				node, err := f.ClientSet.CoreV1().Nodes().Get(framework.TestContext.NodeName, metav1.GetOptions{})
				framework.ExpectNoError(err)
				return numberOfDevices(node, resourceName)
			}, 30*time.Second, framework.Poll).Should(Equal(devsLen))

			By("Creating one pod on node with at least one fake-device")
			podRECMD := "devs=$(ls /tmp/ | egrep '^Dev-[0-9]+$') && echo stub devices: $devs"
			pod1 := f.PodClient().CreateSync(makeBusyboxPod(resourceName, podRECMD))
			deviceIDRE := "stub devices: (Dev-[0-9]+)"
			count1, devId1 := parseLogFromNRuns(f, pod1.Name, pod1.Name, 0, deviceIDRE)
			Expect(devId1).To(Not(Equal("")))

			pod1, err = f.PodClient().Get(pod1.Name, metav1.GetOptions{})
			framework.ExpectNoError(err)

			By("Restarting Kubelet and waiting for the current running pod to restart")
			restartKubelet()

			By("Confirming that after a kubelet and pod restart, fake-device assignement is kept")
			count1, devIdRestart1 := parseLogFromNRuns(f, pod1.Name, pod1.Name, count1+1, deviceIDRE)
			Expect(devIdRestart1).To(Equal(devId1))

			By("Wait for node is ready")
			framework.WaitForAllNodesSchedulable(f.ClientSet, framework.TestContext.NodeSchedulableTimeout)

			By("Re-Register resources")
			dp1 = dm.NewDevicePluginStub(devs, socketPath)
			dp1.SetAllocFunc(stubAllocFunc)
			err = dp1.Start()
			framework.ExpectNoError(err)

			err = dp1.Register(pluginapi.KubeletSocket, resourceName)
			framework.ExpectNoError(err)

			By("Waiting for resource to become available on the local node after re-registration")
			Eventually(func() int64 {
				node, err := f.ClientSet.CoreV1().Nodes().Get(framework.TestContext.NodeName, metav1.GetOptions{})
				framework.ExpectNoError(err)
				return numberOfDevices(node, resourceName)
			}, 30*time.Second, framework.Poll).Should(Equal(devsLen))

			By("Creating another pod")
			pod2 := f.PodClient().CreateSync(makeBusyboxPod(resourceName, podRECMD))

			By("Checking that pods got a different GPU")
			count2, devId2 := parseLogFromNRuns(f, pod2.Name, pod2.Name, 1, deviceIDRE)

			Expect(devId1).To(Not(Equal(devId2)))

			By("Deleting device plugin.")
			err = dp1.Stop()
			framework.ExpectNoError(err)

			By("Waiting for stub device plugin to become unavailable on the local node")
			Eventually(func() bool {
				node, err := f.ClientSet.CoreV1().Nodes().Get(framework.TestContext.NodeName, metav1.GetOptions{})
				framework.ExpectNoError(err)
				return numberOfDevices(node, resourceName) <= 0
			}, 10*time.Minute, framework.Poll).Should(BeTrue())

			By("Checking that scheduled pods can continue to run even after we delete device plugin.")
			count1, devIdRestart1 = parseLogFromNRuns(f, pod1.Name, pod1.Name, count1+1, deviceIDRE)
			Expect(devIdRestart1).To(Equal(devId1))
			count2, devIdRestart2 := parseLogFromNRuns(f, pod2.Name, pod2.Name, count2+1, deviceIDRE)
			Expect(devIdRestart2).To(Equal(devId2))

			By("Restarting Kubelet.")
			restartKubelet()

			By("Checking that scheduled pods can continue to run even after we delete device plugin and restart Kubelet.")
			count1, devIdRestart1 = parseLogFromNRuns(f, pod1.Name, pod1.Name, count1+2, deviceIDRE)
			Expect(devIdRestart1).To(Equal(devId1))
			count2, devIdRestart2 = parseLogFromNRuns(f, pod2.Name, pod2.Name, count2+2, deviceIDRE)
			Expect(devIdRestart2).To(Equal(devId2))

			// Cleanup
			f.PodClient().DeleteSync(pod1.Name, &metav1.DeleteOptions{}, framework.DefaultPodDeletionTimeout)
			f.PodClient().DeleteSync(pod2.Name, &metav1.DeleteOptions{}, framework.DefaultPodDeletionTimeout)
		})
	})
})

// Serial because the test restarts Kubelet
// Slow because this is a soak test
var _ = framework.KubeDescribe("Device Plugin Soak [Feature:DevicePluginSoak] [Soak] [Slow] [Serial] [Disruptive]", func() {
	f := framework.NewDefaultFramework("device-plugin-errors")

	Context("DevicePlugin", func() {
		By("Enabling support for Device Plugin")
		tempSetCurrentKubeletConfig(f, func(initialConfig *kubeletconfig.KubeletConfiguration) {
			initialConfig.FeatureGates[string(features.DevicePlugins)] = true
		})

		const (
			// how many waves of device plugins
			scale = 10
			// how many milliseconds to wait before dumping the next wave of device plugins
			millisecondsBetweenWaves = 500 * time.Millisecond
			// ping to check device plugin health periodicity and continuously
			pingPeriodInSeconds = 15 * time.Second
			// timeout of the whole soak test
			soakTimeout = 30 * time.Minute
		)

		description := fmt.Sprintf("Kubelet should survive after handling multi device plugins for a duration of %v, "+
			"scaling up to %v device plugins per node.", soakTimeout, scale)

		It(description, func() {
			defer GinkgoRecover()
			var wg sync.WaitGroup
			wg.Add(scale)
			for i := 0; i < scale; i++ {
				go func() {
					wave := fmt.Sprintf("wave-%v", strconv.Itoa(i))
					framework.Logf("Starting device plugin soak test, wave = %v", wave)
					setUpAndConsumeResourceWithSleepOf(f, pingPeriodInSeconds, wave, soakTimeout)
					framework.Logf("Completed device plugin soak test, wave = %v", i)
					wg.Done()
				}()
				// gab between waves
				time.Sleep(millisecondsBetweenWaves)
			}
			framework.Logf("Waiting for all %v device plugin soak waves to complete", scale)
			wg.Wait()
		})
	})
})

func setUpAndConsumeResourceWithSleepOf(f *framework.Framework, sleep time.Duration, wave string, timeout time.Duration) {
	framework.Logf("Start stub device plugin")
	// setup fake device to be consumed
	// use wave as deviceID here
	devs := []*pluginapi.Device{
		{ID: wave, Health: pluginapi.Healthy},
	}

	socketPath := pluginapi.DevicePluginPath + "dp." + wave

	dp1 := dm.NewDevicePluginStub(devs, socketPath)
	dp1.SetAllocFunc(stubAllocFunc)
	err := dp1.Start()
	framework.ExpectNoError(err)

	framework.Logf("Register resources")
	resourceName := wave + ".com/resources"
	err = dp1.Register(pluginapi.KubeletSocket, resourceName)
	framework.ExpectNoError(err)

	framework.Logf("Waiting for the resource exported by the stub device plugin to become available on the local node")
	devsLen := int64(len(devs))
	Eventually(func() int64 {
		node, err := f.ClientSet.CoreV1().Nodes().Get(framework.TestContext.NodeName, metav1.GetOptions{})
		framework.ExpectNoError(err)
		return numberOfDevices(node, resourceName)
	}, 30*time.Second, framework.Poll).Should(Equal(devsLen))

	framework.Logf("Creating one pod on node with at least one fake-device")
	reCmd := fmt.Sprintf("devs=$(ls /tmp/ | egrep '^%v+$') && echo stub devices: $devs", wave)

	// new timer for break out this soak test
	t := time.NewTimer(timeout)
	for {
		select {
		case <-t.C:
			// break if soak timeout
			framework.Logf("Soak timeout, break wave: %s", wave)
			return
		default:
			// create pod to consume this resource periodicity
			pod := f.PodClient().CreateSync(makeBusyboxPod(resourceName, reCmd))

			deviceIDRE := fmt.Sprintf("stub devices: (%v+)", wave)
			_, devId := parseLogFromNRuns(f, pod.Name, pod.Name, 0, deviceIDRE)

			Expect(devId).To(Equal(wave))

			// cleanup pod
			f.PodClient().DeleteSync(pod.Name, &metav1.DeleteOptions{}, framework.DefaultPodDeletionTimeout)

			// sleep to do the next iteration
			time.Sleep(sleep)
		}
	}
}

// makeBusyboxPod returns a simple Pod spec with a busybox container
// that requests resourceName and runs the specified command.
func makeBusyboxPod(resourceName, cmd string) *v1.Pod {
	podName := "device-plugin-test-" + string(uuid.NewUUID())
	rl := v1.ResourceList{v1.ResourceName(resourceName): *resource.NewQuantity(1, resource.DecimalSI)}

	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: podName},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyAlways,
			Containers: []v1.Container{{
				Image: busyboxImage,
				Name:  podName,
				// Runs the specified command in the test pod.
				Command: []string{"sh", "-c", cmd},
				Resources: v1.ResourceRequirements{
					Limits:   rl,
					Requests: rl,
				},
			}},
		},
	}
}

// parseLogFromNRuns returns restart count of the specified container
// after it has been restarted at least restartCount times,
// and the matching string for the specified regular expression parsed from the container logs.
func parseLogFromNRuns(f *framework.Framework, podName string, contName string, restartCount int32, re string) (int32, string) {
	var count int32
	// Wait till pod has been restarted at least restartCount times.
	Eventually(func() bool {
		p, err := f.PodClient().Get(podName, metav1.GetOptions{})
		if err != nil || len(p.Status.ContainerStatuses) < 1 {
			return false
		}
		count = p.Status.ContainerStatuses[0].RestartCount
		return count >= restartCount
	}, 5*time.Minute, framework.Poll).Should(BeTrue())

	logs, err := framework.GetPodLogs(f.ClientSet, f.Namespace.Name, podName, contName)
	if err != nil {
		framework.Failf("GetPodLogs for pod %q failed: %v", podName, err)
	}

	framework.Logf("got pod logs: %v", logs)
	regex := regexp.MustCompile(re)
	matches := regex.FindStringSubmatch(logs)
	if len(matches) < 2 {
		return count, ""
	}

	return count, matches[1]
}

// numberOfDevices returns the number of devices of resourceName advertised by a node
func numberOfDevices(node *v1.Node, resourceName string) int64 {
	val, ok := node.Status.Capacity[v1.ResourceName(resourceName)]
	if !ok {
		return 0
	}

	return val.Value()
}

// stubAllocFunc will pass to stub device plugin
func stubAllocFunc(r *pluginapi.AllocateRequest, devs map[string]pluginapi.Device) (*pluginapi.AllocateResponse, error) {
	var response pluginapi.AllocateResponse
	for _, requestID := range r.DevicesIDs {
		dev, ok := devs[requestID]
		if !ok {
			return nil, fmt.Errorf("invalid allocation request with non-existing device %s", requestID)
		}

		if dev.Health != pluginapi.Healthy {
			return nil, fmt.Errorf("invalid allocation request with unhealthy device: %s", requestID)
		}

		// create fake device file
		fpath := filepath.Join("/tmp", dev.ID)

		// clean first
		os.RemoveAll(fpath)
		f, err := os.Create(fpath)
		if err != nil && !os.IsExist(err) {
			return nil, fmt.Errorf("failed to create fake device file: %s", err)
		}

		f.Close()

		response.Mounts = append(response.Mounts, &pluginapi.Mount{
			ContainerPath: fpath,
			HostPath:      fpath,
		})
	}

	return &response, nil
}
