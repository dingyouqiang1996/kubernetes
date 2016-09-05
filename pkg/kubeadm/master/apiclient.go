/*
Copyright 2016 The Kubernetes Authors.

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

package kubemaster

import (
	"encoding/json"
	"fmt"
	"time"

	"k8s.io/kubernetes/pkg/api"
	unversionedapi "k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/extensions"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	clientcmdapi "k8s.io/kubernetes/pkg/client/unversioned/clientcmd/api"
	"k8s.io/kubernetes/pkg/util/wait"
)

func CreateClientAndWaitForAPI(adminConfig *clientcmdapi.Config) (*clientset.Clientset, error) {
	adminClientConfig, err := clientcmd.NewDefaultClientConfig(
		*adminConfig,
		&clientcmd.ConfigOverrides{},
	).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("<master/apiclient> failed to create API client configuration [%s]", err)
	}

	fmt.Println("<master/apiclient> created API client configuration")

	client, err := clientset.NewForConfig(adminClientConfig)
	if err != nil {
		return nil, fmt.Errorf("<master/apiclient> failed to create API client [%s]", err)
	}

	fmt.Println("<master/apiclient> created API client, waiting for the control plane to become ready")

	start := time.Now()
	wait.PollInfinite(500*time.Millisecond, func() (bool, error) {
		cs, err := client.ComponentStatuses().List(api.ListOptions{})
		if err != nil {
			return false, nil
		}
		// TODO revisit this when we implement HA
		if len(cs.Items) < 3 {
			fmt.Println("<master/apiclient> not all control plane components are ready yet")
			return false, nil
		}
		for _, item := range cs.Items {
			for _, condition := range item.Conditions {
				if condition.Type != api.ComponentHealthy {
					fmt.Printf("<master/apiclient> control plane component %q is still unhealthy: %#v\n", item.ObjectMeta.Name, item.Conditions)
					return false, nil
				}
			}
		}

		fmt.Printf("<master/apiclient> all control plane components are healthy after %f seconds\n", time.Since(start).Seconds())
		return true, nil
	})

	fmt.Println("<master/apiclient> waiting for at least one node to register and become ready")
	start = time.Now()
	wait.PollInfinite(500*time.Millisecond, func() (bool, error) {
		nodeList, err := client.Nodes().List(api.ListOptions{})
		if err != nil {
			fmt.Println("<master/apiclient> not able to list nodes (will retry)")
			return false, nil
		}
		if len(nodeList.Items) < 1 {
			//fmt.Printf("<master/apiclient> %d nodes have registered so far", len(nodeList.Items))
			return false, nil
		}
		n := &nodeList.Items[0]
		if !api.IsNodeReady(n) {
			fmt.Println("<master/apiclient> first node has registered, but is not ready yet")
			return false, nil
		}

		fmt.Printf("<master/apiclient> first node is ready after %f seconds\n", time.Since(start).Seconds())
		return true, nil
	})

	return client, nil
}

func NewDaemonSet(daemonName string, podSpec api.PodSpec) *extensions.DaemonSet {
	l := map[string]string{"component": daemonName, "tier": "node"}
	return &extensions.DaemonSet{
		ObjectMeta: api.ObjectMeta{Name: daemonName},
		Spec: extensions.DaemonSetSpec{
			Selector: &unversionedapi.LabelSelector{MatchLabels: l},
			Template: api.PodTemplateSpec{
				ObjectMeta: api.ObjectMeta{Labels: l},
				Spec:       podSpec,
			},
		},
	}
}

func NewDeployment(deploymentName string, replicas int32, podSpec api.PodSpec) *extensions.Deployment {
	l := map[string]string{"name": deploymentName}
	return &extensions.Deployment{
		ObjectMeta: api.ObjectMeta{Name: deploymentName},
		Spec: extensions.DeploymentSpec{
			Replicas: replicas,
			Selector: &unversionedapi.LabelSelector{MatchLabels: l},
			Template: api.PodTemplateSpec{
				ObjectMeta: api.ObjectMeta{Labels: l},
				Spec:       podSpec,
			},
		},
	}
}

// It's safe to do this for alpha, as we don't have HA and there is no way we can get
// more then one node here (TODO find a way to determine owr own node name)
func findMyself(client *clientset.Clientset) (*api.Node, error) {
	nodeList, err := client.Nodes().List(api.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to list nodes [%s]", err)
	}
	if len(nodeList.Items) < 1 {
		return nil, fmt.Errorf("no nodes found")
	}
	node := &nodeList.Items[0]
	return node, nil
}

func UpdateMasterRoleLabelsAndTaints(client *clientset.Clientset) error {
	n, err := findMyself(client)
	if err != nil {
		return fmt.Errorf("<master/apiclient> failed to update master node - %s", err)
	}

	n.ObjectMeta.Labels["kubeadm.alpha.kubernetes.io/role"] = "master"
	taintsAnnotation, _ := json.Marshal([]api.Taint{{Key: "dedicated", Value: "master", Effect: "NoSchedule"}})
	n.ObjectMeta.Annotations[api.TaintsAnnotationKey] = string(taintsAnnotation)

	if _, err := client.Nodes().Update(n); err != nil {
		return fmt.Errorf("<master/apiclient> failed to update master node - %s", err)
	}

	return nil
}

func SetMasterTaintTolerations(meta *api.ObjectMeta) {
	tolerationsAnnotation, _ := json.Marshal([]api.Toleration{{Key: "dedicated", Value: "master", Effect: "NoSchedule"}})
	if meta.Annotations == nil {
		meta.Annotations = map[string]string{}
	}
	meta.Annotations[api.TolerationsAnnotationKey] = string(tolerationsAnnotation)
}

func SetMasterNodeAffinity(meta *api.ObjectMeta) {
	nodeAffinity := &api.NodeAffinity{
		RequiredDuringSchedulingIgnoredDuringExecution: &api.NodeSelector{
			NodeSelectorTerms: []api.NodeSelectorTerm{{
				MatchExpressions: []api.NodeSelectorRequirement{{
					Key: "kubeadm.alpha.kubernetes.io/role", Operator: api.NodeSelectorOpIn, Values: []string{"master"},
				}},
			}},
		},
	}
	affinityAnnotation, _ := json.Marshal(api.Affinity{NodeAffinity: nodeAffinity})
	if meta.Annotations == nil {
		meta.Annotations = map[string]string{}
	}
	meta.Annotations[api.AffinityAnnotationKey] = string(affinityAnnotation)
}
