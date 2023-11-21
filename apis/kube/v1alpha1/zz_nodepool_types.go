// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MetadataInitParameters struct {

	// Annotations to apply to each node
	// annotations
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Finalizers to apply to each node. A finalizer name must be fully qualified, e.g. kubernetes.io/pv-protection , where you prefix it with hostname of your service which is related to the controller responsible for the finalizer.
	// finalizers
	Finalizers []*string `json:"finalizers,omitempty" tf:"finalizers,omitempty"`

	// Labels to apply to each node
	// labels
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type MetadataObservation struct {

	// Annotations to apply to each node
	// annotations
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Finalizers to apply to each node. A finalizer name must be fully qualified, e.g. kubernetes.io/pv-protection , where you prefix it with hostname of your service which is related to the controller responsible for the finalizer.
	// finalizers
	Finalizers []*string `json:"finalizers,omitempty" tf:"finalizers,omitempty"`

	// Labels to apply to each node
	// labels
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type MetadataParameters struct {

	// Annotations to apply to each node
	// annotations
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations" tf:"annotations,omitempty"`

	// Finalizers to apply to each node. A finalizer name must be fully qualified, e.g. kubernetes.io/pv-protection , where you prefix it with hostname of your service which is related to the controller responsible for the finalizer.
	// finalizers
	// +kubebuilder:validation:Optional
	Finalizers []*string `json:"finalizers" tf:"finalizers,omitempty"`

	// Labels to apply to each node
	// labels
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels" tf:"labels,omitempty"`
}

type NodePoolInitParameters struct {

	// should the pool use the anti-affinity feature. Default to false. Changing this value recreates the resource.
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity *bool `json:"antiAffinity,omitempty" tf:"anti_affinity,omitempty"`

	// Enable auto-scaling for the pool. Default to false.
	// Enable auto-scaling for the pool
	Autoscale *bool `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// number of nodes to start.
	// Number of nodes you desire in the pool
	DesiredNodes *float64 `json:"desiredNodes,omitempty" tf:"desired_nodes,omitempty"`

	// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
	// Changing this value recreates the resource.
	// Flavor name
	FlavorName *string `json:"flavorName,omitempty" tf:"flavor_name,omitempty"`

	// maximum number of nodes allowed in the pool. Setting desired_nodes over this value will raise an error.
	// Number of nodes you desire in the pool
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// minimum number of nodes allowed in the pool. Setting desired_nodes under this value will raise an error.
	// Number of nodes you desire in the pool
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// should the nodes be billed on a monthly basis. Default to false. Changing this value recreates the resource.
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled *bool `json:"monthlyBilled,omitempty" tf:"monthly_billed,omitempty"`

	// The name of the nodepool. Warning: _ char is not allowed! Changing this value recreates the resource.
	// NodePool resource name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the public cloud project. If omitted, the OVH_CLOUD_PROJECT_SERVICE environment variable is used. Changing this value recreates the resource.
	// Service name
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Node pool template
	Template []TemplateInitParameters `json:"template,omitempty" tf:"template,omitempty"`
}

type NodePoolObservation struct {

	// should the pool use the anti-affinity feature. Default to false. Changing this value recreates the resource.
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity *bool `json:"antiAffinity,omitempty" tf:"anti_affinity,omitempty"`

	// Enable auto-scaling for the pool. Default to false.
	// Enable auto-scaling for the pool
	Autoscale *bool `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// Number of nodes which are actually ready in the pool
	// Number of nodes which are actually ready in the pool
	AvailableNodes *float64 `json:"availableNodes,omitempty" tf:"available_nodes,omitempty"`

	// Creation date
	// Creation date
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Number of nodes present in the pool
	// Number of nodes present in the pool
	CurrentNodes *float64 `json:"currentNodes,omitempty" tf:"current_nodes,omitempty"`

	// number of nodes to start.
	// Number of nodes you desire in the pool
	DesiredNodes *float64 `json:"desiredNodes,omitempty" tf:"desired_nodes,omitempty"`

	// Flavor name
	// Flavor name
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
	// Changing this value recreates the resource.
	// Flavor name
	FlavorName *string `json:"flavorName,omitempty" tf:"flavor_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The id of the managed kubernetes cluster. Changing this value recreates the resource.
	// Kube ID
	KubeID *string `json:"kubeId,omitempty" tf:"kube_id,omitempty"`

	// maximum number of nodes allowed in the pool. Setting desired_nodes over this value will raise an error.
	// Number of nodes you desire in the pool
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// minimum number of nodes allowed in the pool. Setting desired_nodes under this value will raise an error.
	// Number of nodes you desire in the pool
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// should the nodes be billed on a monthly basis. Default to false. Changing this value recreates the resource.
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled *bool `json:"monthlyBilled,omitempty" tf:"monthly_billed,omitempty"`

	// The name of the nodepool. Warning: _ char is not allowed! Changing this value recreates the resource.
	// NodePool resource name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Project id
	// Project id
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The id of the public cloud project. If omitted, the OVH_CLOUD_PROJECT_SERVICE environment variable is used. Changing this value recreates the resource.
	// Service name
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Status describing the state between number of nodes wanted and available ones
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus *string `json:"sizeStatus,omitempty" tf:"size_status,omitempty"`

	// Current status
	// Current status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Node pool template
	Template []TemplateObservation `json:"template,omitempty" tf:"template,omitempty"`

	// Number of nodes with the latest version installed in the pool
	// Number of nodes with latest version installed in the pool
	UpToDateNodes *float64 `json:"upToDateNodes,omitempty" tf:"up_to_date_nodes,omitempty"`

	// Last update date
	// Last update date
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type NodePoolParameters struct {

	// should the pool use the anti-affinity feature. Default to false. Changing this value recreates the resource.
	// Enable anti affinity groups for nodes in the pool
	// +kubebuilder:validation:Optional
	AntiAffinity *bool `json:"antiAffinity,omitempty" tf:"anti_affinity,omitempty"`

	// Enable auto-scaling for the pool. Default to false.
	// Enable auto-scaling for the pool
	// +kubebuilder:validation:Optional
	Autoscale *bool `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// number of nodes to start.
	// Number of nodes you desire in the pool
	// +kubebuilder:validation:Optional
	DesiredNodes *float64 `json:"desiredNodes,omitempty" tf:"desired_nodes,omitempty"`

	// a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
	// Changing this value recreates the resource.
	// Flavor name
	// +kubebuilder:validation:Optional
	FlavorName *string `json:"flavorName,omitempty" tf:"flavor_name,omitempty"`

	// The id of the managed kubernetes cluster. Changing this value recreates the resource.
	// Kube ID
	// +crossplane:generate:reference:type=saagie.io/provider-ovh/apis/kube/v1alpha1.Kube
	// +kubebuilder:validation:Optional
	KubeID *string `json:"kubeId,omitempty" tf:"kube_id,omitempty"`

	// Reference to a Kube in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDRef *v1.Reference `json:"kubeIdRef,omitempty" tf:"-"`

	// Selector for a Kube in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDSelector *v1.Selector `json:"kubeIdSelector,omitempty" tf:"-"`

	// maximum number of nodes allowed in the pool. Setting desired_nodes over this value will raise an error.
	// Number of nodes you desire in the pool
	// +kubebuilder:validation:Optional
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// minimum number of nodes allowed in the pool. Setting desired_nodes under this value will raise an error.
	// Number of nodes you desire in the pool
	// +kubebuilder:validation:Optional
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// should the nodes be billed on a monthly basis. Default to false. Changing this value recreates the resource.
	// Enable monthly billing on all nodes in the pool
	// +kubebuilder:validation:Optional
	MonthlyBilled *bool `json:"monthlyBilled,omitempty" tf:"monthly_billed,omitempty"`

	// The name of the nodepool. Warning: _ char is not allowed! Changing this value recreates the resource.
	// NodePool resource name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the public cloud project. If omitted, the OVH_CLOUD_PROJECT_SERVICE environment variable is used. Changing this value recreates the resource.
	// Service name
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Node pool template
	// +kubebuilder:validation:Optional
	Template []TemplateParameters `json:"template,omitempty" tf:"template,omitempty"`
}

type SpecInitParameters struct {

	// Taints to apply to each node
	// taints
	Taints []map[string]*string `json:"taints,omitempty" tf:"taints,omitempty"`

	// If true, set nodes as un-schedulable
	// unschedulable
	Unschedulable *bool `json:"unschedulable,omitempty" tf:"unschedulable,omitempty"`
}

type SpecObservation struct {

	// Taints to apply to each node
	// taints
	Taints []map[string]*string `json:"taints,omitempty" tf:"taints,omitempty"`

	// If true, set nodes as un-schedulable
	// unschedulable
	Unschedulable *bool `json:"unschedulable,omitempty" tf:"unschedulable,omitempty"`
}

type SpecParameters struct {

	// Taints to apply to each node
	// taints
	// +kubebuilder:validation:Optional
	Taints []map[string]*string `json:"taints" tf:"taints,omitempty"`

	// If true, set nodes as un-schedulable
	// unschedulable
	// +kubebuilder:validation:Optional
	Unschedulable *bool `json:"unschedulable" tf:"unschedulable,omitempty"`
}

type TemplateInitParameters struct {

	// Metadata of each node in the pool
	// metadata
	Metadata []MetadataInitParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Spec of each node in the pool
	// spec
	Spec []SpecInitParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type TemplateObservation struct {

	// Metadata of each node in the pool
	// metadata
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Spec of each node in the pool
	// spec
	Spec []SpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`
}

type TemplateParameters struct {

	// Metadata of each node in the pool
	// metadata
	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata" tf:"metadata,omitempty"`

	// Spec of each node in the pool
	// spec
	// +kubebuilder:validation:Optional
	Spec []SpecParameters `json:"spec" tf:"spec,omitempty"`
}

// NodePoolSpec defines the desired state of NodePool
type NodePoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodePoolParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NodePoolInitParameters `json:"initProvider,omitempty"`
}

// NodePoolStatus defines the observed state of NodePool.
type NodePoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodePoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodePool is the Schema for the NodePools API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type NodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.flavorName) || (has(self.initProvider) && has(self.initProvider.flavorName))",message="spec.forProvider.flavorName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   NodePoolSpec   `json:"spec"`
	Status NodePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolList contains a list of NodePools
type NodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePool `json:"items"`
}

// Repository type metadata.
var (
	NodePool_Kind             = "NodePool"
	NodePool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodePool_Kind}.String()
	NodePool_KindAPIVersion   = NodePool_Kind + "." + CRDGroupVersion.String()
	NodePool_GroupVersionKind = CRDGroupVersion.WithKind(NodePool_Kind)
)

func init() {
	SchemeBuilder.Register(&NodePool{}, &NodePoolList{})
}
