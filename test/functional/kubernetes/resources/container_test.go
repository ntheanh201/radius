// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package container_test

import (
	"testing"

	"github.com/Azure/radius/pkg/workloads"
	"github.com/Azure/radius/test/kubernetestest"
	"github.com/Azure/radius/test/validation"
)

func Test_ContainerHttpBinding(t *testing.T) {
	template := "testdata/kubernetes-resources-container-httpbinding.bicep"
	application := "kubernetes-resources-container-httpbinding"
	test := kubernetestest.NewApplicationTest(t, application, []kubernetestest.Step{
		{
			Executor: kubernetestest.NewDeployStepExecutor(template),
			Components: &validation.ComponentSet{
				Components: []validation.Component{
					{
						ApplicationName: application,
						ComponentName:   "frontend",
						OutputResources: map[string]validation.ExpectedOutputResource{
							workloads.LocalIDDeployment: validation.NewOutputResource(workloads.LocalIDDeployment, workloads.OutputResourceTypeKubernetes, workloads.ResourceKindKubernetes, true),
							workloads.LocalIDService:    validation.NewOutputResource(workloads.LocalIDService, workloads.OutputResourceTypeKubernetes, workloads.ResourceKindKubernetes, true),
						},
					},
					{
						ApplicationName: application,
						ComponentName:   "backend",
						OutputResources: map[string]validation.ExpectedOutputResource{
							workloads.LocalIDDeployment: validation.NewOutputResource(workloads.LocalIDDeployment, workloads.OutputResourceTypeKubernetes, workloads.ResourceKindKubernetes, true),
							workloads.LocalIDService:    validation.NewOutputResource(workloads.LocalIDService, workloads.OutputResourceTypeKubernetes, workloads.ResourceKindKubernetes, true),
						},
					},
				},
			},
			Pods: &validation.K8sObjectSet{
				Namespaces: map[string][]validation.K8sObject{
					"default": {
						validation.NewK8sObjectForComponent(application, "frontend"),
						validation.NewK8sObjectForComponent(application, "backend"),
					},
				},
			},
		},
	})

	test.Test(t)
}
