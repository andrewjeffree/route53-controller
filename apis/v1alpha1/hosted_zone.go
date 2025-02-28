// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HostedZoneSpec defines the desired state of HostedZone.
//
// A complex type that contains general information about the hosted zone.
type HostedZoneSpec struct {

	// If you want to associate a reusable delegation set with this hosted zone,
	// the ID that Amazon Route 53 assigned to the reusable delegation set when
	// you created it. For more information about reusable delegation sets, see
	// CreateReusableDelegationSet (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateReusableDelegationSet.html).
	//
	// If you are using a reusable delegation set to create a public hosted zone
	// for a subdomain, make sure that the parent hosted zone doesn't use one or
	// more of the same name servers. If you have overlapping nameservers, the operation
	// will cause a ConflictingDomainsExist error.
	DelegationSetID *string `json:"delegationSetID,omitempty"`
	// (Optional) A complex type that contains the following optional values:
	//
	//   - For public and private hosted zones, an optional comment
	//
	//   - For private hosted zones, an optional PrivateZone element
	//
	// If you don't specify a comment or the PrivateZone element, omit HostedZoneConfig
	// and the other elements.
	HostedZoneConfig *HostedZoneConfig `json:"hostedZoneConfig,omitempty"`
	// The name of the domain. Specify a fully qualified domain name, for example,
	// www.example.com. The trailing dot is optional; Amazon Route 53 assumes that
	// the domain name is fully qualified. This means that Route 53 treats www.example.com
	// (without a trailing dot) and www.example.com. (with a trailing dot) as identical.
	//
	// If you're creating a public hosted zone, this is the name you have registered
	// with your DNS registrar. If your domain name is registered with a registrar
	// other than Route 53, change the name servers for your domain to the set of
	// NameServers that CreateHostedZone returns in DelegationSet.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// A complex type that contains a list of the tags that you want to add to the
	// specified health check or hosted zone and/or the tags that you want to edit
	// Value for.
	//
	// You can add a maximum of 10 tags to a health check or a hosted zone.
	Tags []*Tag `json:"tags,omitempty"`
	// (Private hosted zones only) A complex type that contains information about
	// the Amazon VPC that you're associating with this hosted zone.
	//
	// You can specify only one Amazon VPC when you create a private hosted zone.
	// If you are associating a VPC with a hosted zone with this request, the paramaters
	// VPCId and VPCRegion are also required.
	//
	// To associate additional Amazon VPCs with the hosted zone, use AssociateVPCWithHostedZone
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_AssociateVPCWithHostedZone.html)
	// after you create a hosted zone.
	VPC *VPC `json:"vpc,omitempty"`
}

// HostedZoneStatus defines the observed state of HostedZone
type HostedZoneStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The value that you specified for CallerReference when you created the hosted
	// zone.
	// +kubebuilder:validation:Optional
	CallerReference *string `json:"callerReference,omitempty"`
	// A complex type that includes the Comment and PrivateZone elements. If you
	// omitted the HostedZoneConfig and Comment elements from the request, the Config
	// and Comment elements don't appear in the response.
	// +kubebuilder:validation:Optional
	Config *HostedZoneConfig `json:"config,omitempty"`
	// A complex type that describes the name servers for this hosted zone.
	// +kubebuilder:validation:Optional
	DelegationSet *DelegationSet `json:"delegationSet,omitempty"`
	// The ID that Amazon Route 53 assigned to the hosted zone when you created
	// it.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty"`
	// If the hosted zone was created by another service, the service that created
	// the hosted zone. When a hosted zone is created by another service, you can't
	// edit or delete it using Route 53.
	// +kubebuilder:validation:Optional
	LinkedService *LinkedService `json:"linkedService,omitempty"`
	// The number of resource record sets in the hosted zone.
	// +kubebuilder:validation:Optional
	ResourceRecordSetCount *int64 `json:"resourceRecordSetCount,omitempty"`
}

// HostedZone is the Schema for the HostedZones API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type HostedZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedZoneSpec   `json:"spec,omitempty"`
	Status            HostedZoneStatus `json:"status,omitempty"`
}

// HostedZoneList contains a list of HostedZone
// +kubebuilder:object:root=true
type HostedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedZone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HostedZone{}, &HostedZoneList{})
}
