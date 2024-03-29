package utils

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCalculatePodRequests(t *testing.T) {
	pods := []v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod1",
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{{
					Name: "container1",
					Resources: v1.ResourceRequirements{
						Requests: map[v1.ResourceName]resource.Quantity{
							v1.ResourceCPU:    *resource.NewQuantity(1, resource.DecimalSI),
							v1.ResourceMemory: *resource.NewQuantity(10, resource.DecimalSI),
						},
					},
				}, {
					Name: "container2",
					Resources: v1.ResourceRequirements{
						Requests: map[v1.ResourceName]resource.Quantity{
							v1.ResourceCPU:    *resource.NewQuantity(1, resource.DecimalSI),
							v1.ResourceMemory: *resource.NewQuantity(10, resource.DecimalSI),
						},
					},
				}},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod2",
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{{
					Name: "container1",
					Resources: v1.ResourceRequirements{
						Requests: map[v1.ResourceName]resource.Quantity{
							v1.ResourceCPU:    *resource.NewQuantity(2, resource.DecimalSI),
							v1.ResourceMemory: *resource.NewQuantity(20, resource.DecimalSI),
						},
					},
				}, {
					Name: "container2",
					Resources: v1.ResourceRequirements{
						Requests: map[v1.ResourceName]resource.Quantity{
							v1.ResourceCPU:    *resource.NewQuantity(2, resource.DecimalSI),
							v1.ResourceMemory: *resource.NewQuantity(20, resource.DecimalSI),
						},
					},
				}},
			},
		},
	}

	tests := []struct {
		description   string
		resource      v1.ResourceName
		containerName string
		expect        int64
	}{
		{
			description: "calculate cpu request total",
			resource:    v1.ResourceCPU,
			expect:      6000,
		},
		{
			description: "calculate memory request total",
			resource:    v1.ResourceMemory,
			expect:      60000,
		},
		{
			description:   "calculate cpu request total of container1",
			resource:      v1.ResourceCPU,
			containerName: "container1",
			expect:        3000,
		},
		{
			description:   "calculate memory request total of container1",
			resource:      v1.ResourceMemory,
			containerName: "container1",
			expect:        30000,
		},
	}

	for _, test := range tests {
		requests, err := CalculatePodRequests(pods, test.resource, test.containerName)
		if err != nil {
			t.Errorf("Failed to calculatePodRequests: %v", err)
		}
		if requests != test.expect {
			t.Errorf("expect requests %d actual requests %d", test.expect, requests)
		}
	}

}

func TestCalculatePodTemplateRequests(t *testing.T) {
	podTemplate := &v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name: "podTemplateTest",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{{
				Name: "container1",
				Resources: v1.ResourceRequirements{
					Requests: map[v1.ResourceName]resource.Quantity{
						v1.ResourceCPU:    *resource.NewQuantity(1, resource.DecimalSI),
						v1.ResourceMemory: *resource.NewQuantity(10, resource.DecimalSI),
					},
				},
			}, {
				Name: "container2",
				Resources: v1.ResourceRequirements{
					Requests: map[v1.ResourceName]resource.Quantity{
						v1.ResourceCPU:    *resource.NewQuantity(1, resource.DecimalSI),
						v1.ResourceMemory: *resource.NewQuantity(10, resource.DecimalSI),
					},
				},
			}},
		},
	}

	tests := []struct {
		description string
		resource    v1.ResourceName
		expect      int64
	}{
		{
			description: "calculate cpu request total",
			resource:    v1.ResourceCPU,
			expect:      2000,
		},
		{
			description: "calculate memory request total",
			resource:    v1.ResourceMemory,
			expect:      20000,
		},
	}

	for _, test := range tests {
		requests, err := CalculatePodTemplateRequests(podTemplate, test.resource)
		if err != nil {
			t.Errorf("Failed to CalculatePodTemplateRequests: %v", err)
		}
		if requests != test.expect {
			t.Errorf("expect requests %d actual requests %d", test.expect, requests)
		}
	}

}
