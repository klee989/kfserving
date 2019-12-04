/*
Copyright 2019 kubeflow.org.
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

package v1alpha2

import (
	"github.com/kubeflow/kfserving/pkg/constants"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
)

type InferenceRouterSpec {
	// List of edges to route traffic to
	Edges *v1.ObjectReference[] `json:"edges"`
	// Routing configuration
	ModelTesting ModelTestingSpec `json:"modelTesting"`
}

type ModelTestingSpec {
	// A/B testing configuration
	ABTest ABTestSpec `json:"ABTest,omitempty"`
}

// ABTestSpec defines parameters required for an A/B test.
type ABTestSpec {
	// Name of numeric metric we are A/B testing for improvement.
	// +required
	MetricName string `json:"metricName"`
	// Absolute minimum improvement in the metric that should yield a positive result from the A/B test.
	// +required
	MinimumDetectableEffect float `json:"minimumDetectableEffect"`
	// Percent chance that if the minimum detectable effect threshold is not met, the test yields a negative result.
	// +optional
	Confidence int `json:"confidence,omitempty"`
	// Percent chance that if the minimum detectable effect threshold is met, the test yields a poistive result.
	// +optional
	Power int `json:"power,omitempty"`
	// Base group's metric value, if known a priori. Ignores MaximumPercentError and EstimationConfidence if set.
	// +optional
	BaseRate float `json:"baseRate,omitempty"`
	// Largest allowed percent deviation from the true metric value in base rate metric estimation.
	// +optional
	MaximumPercentError int `json:"maximumPercentError,omitempty"`
	// Confidence that the metric estimate deviates from its true value by no more than MaximumPercentError.
	// +optional
	EstimationConfidence `json:"estimationConfidence,omitempty"`
	// Random seed for assigning users to a group.
	// +optional
	Seed int `json:"seed,omitempty"`
	// TrafficPercent defines the percentage of users routed to the B group, if not 50%.
	// +optional
	TrafficPercent int `json:"trafficPercent,omitempty"`
}

// InferenceRouterStatus defines the observed state of InferenceRouter
type InferenceRouterStatus struct {
	duckv1beta1.Status `json:",inline"`
	// URL of the InferenceRouter
	URL string `json:"url,omitempty"`
	// Statuses for the default endpoints of the InferenceRouter
	EdgeStatusMap map[string]*StatusConfigurationSpec `json:"edgeStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InferenceService is the Schema for the services API
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="URL",type="string",JSONPath=".status.url"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:path=inferenceservices,shortName=inferenceservice
type InferenceRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InferenceRouterSpec   `json:"spec,omitempty"`
	Status InferenceRouterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InferenceRouterList contains a list of InferenceRouter
type InferenceRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InferenceRouter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InferenceRouter{}, &InferenceRouterList{})
}
