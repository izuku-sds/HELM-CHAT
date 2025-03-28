// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	trainerv1alpha1 "github.com/kubeflow/trainer/pkg/apis/trainer/v1alpha1"
	applyconfigurationtrainerv1alpha1 "github.com/kubeflow/trainer/pkg/client/applyconfiguration/trainer/v1alpha1"
	scheme "github.com/kubeflow/trainer/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// TrainingRuntimesGetter has a method to return a TrainingRuntimeInterface.
// A group's client should implement this interface.
type TrainingRuntimesGetter interface {
	TrainingRuntimes(namespace string) TrainingRuntimeInterface
}

// TrainingRuntimeInterface has methods to work with TrainingRuntime resources.
type TrainingRuntimeInterface interface {
	Create(ctx context.Context, trainingRuntime *trainerv1alpha1.TrainingRuntime, opts v1.CreateOptions) (*trainerv1alpha1.TrainingRuntime, error)
	Update(ctx context.Context, trainingRuntime *trainerv1alpha1.TrainingRuntime, opts v1.UpdateOptions) (*trainerv1alpha1.TrainingRuntime, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*trainerv1alpha1.TrainingRuntime, error)
	List(ctx context.Context, opts v1.ListOptions) (*trainerv1alpha1.TrainingRuntimeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *trainerv1alpha1.TrainingRuntime, err error)
	Apply(ctx context.Context, trainingRuntime *applyconfigurationtrainerv1alpha1.TrainingRuntimeApplyConfiguration, opts v1.ApplyOptions) (result *trainerv1alpha1.TrainingRuntime, err error)
	TrainingRuntimeExpansion
}

// trainingRuntimes implements TrainingRuntimeInterface
type trainingRuntimes struct {
	*gentype.ClientWithListAndApply[*trainerv1alpha1.TrainingRuntime, *trainerv1alpha1.TrainingRuntimeList, *applyconfigurationtrainerv1alpha1.TrainingRuntimeApplyConfiguration]
}

// newTrainingRuntimes returns a TrainingRuntimes
func newTrainingRuntimes(c *TrainerV1alpha1Client, namespace string) *trainingRuntimes {
	return &trainingRuntimes{
		gentype.NewClientWithListAndApply[*trainerv1alpha1.TrainingRuntime, *trainerv1alpha1.TrainingRuntimeList, *applyconfigurationtrainerv1alpha1.TrainingRuntimeApplyConfiguration](
			"trainingruntimes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *trainerv1alpha1.TrainingRuntime { return &trainerv1alpha1.TrainingRuntime{} },
			func() *trainerv1alpha1.TrainingRuntimeList { return &trainerv1alpha1.TrainingRuntimeList{} },
		),
	}
}
