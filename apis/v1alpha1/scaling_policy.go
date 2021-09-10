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

// ScalingPolicySpec defines the desired state of ScalingPolicy.
//
// Represents a scaling policy to use with Application Auto Scaling.
//
// For more information about configuring scaling policies for a specific service,
// see Getting started with Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html)
// in the Application Auto Scaling User Guide.
type ScalingPolicySpec struct {
	// The name of the scaling policy.
	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName"`
	// The policy type. This parameter is required if you are creating a scaling
	// policy.
	//
	// The following policy types are supported:
	//
	// TargetTrackingScaling—Not supported for Amazon EMR
	//
	// StepScaling—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon
	// Keyspaces (for Apache Cassandra), or Amazon MSK.
	//
	// For more information, see Target tracking scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html)
	// and Step scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html)
	// in the Application Auto Scaling User Guide.
	PolicyType *string `json:"policyType,omitempty"`
	// The identifier of the resource associated with the scaling policy. This string
	// consists of the resource type and unique identifier.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot Fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the table name. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the index name. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variant - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	//    * Amazon Comprehend document classification endpoint - The resource type
	//    and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.
	//
	//    * Amazon Comprehend entity recognizer endpoint - The resource type and
	//    unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.
	//
	//    * Lambda provisioned concurrency - The resource type is function and the
	//    unique identifier is the function name with a function version or alias
	//    name suffix that is not $LATEST. Example: function:my-function:prod or
	//    function:my-function:1.
	//
	//    * Amazon Keyspaces table - The resource type is table and the unique identifier
	//    is the table name. Example: keyspace/mykeyspace/table/mytable.
	//
	//    * Amazon MSK cluster - The resource type and unique identifier are specified
	//    using the cluster ARN. Example: arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.
	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceID"`
	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    Fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	//
	//    * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The
	//    number of inference units for an Amazon Comprehend document classification
	//    endpoint.
	//
	//    * comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number
	//    of inference units for an Amazon Comprehend entity recognizer endpoint.
	//
	//    * lambda:function:ProvisionedConcurrency - The provisioned concurrency
	//    for a Lambda function.
	//
	//    * cassandra:table:ReadCapacityUnits - The provisioned read capacity for
	//    an Amazon Keyspaces table.
	//
	//    * cassandra:table:WriteCapacityUnits - The provisioned write capacity
	//    for an Amazon Keyspaces table.
	//
	//    * kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB)
	//    for brokers in an Amazon MSK cluster.
	// +kubebuilder:validation:Required
	ScalableDimension *string `json:"scalableDimension"`
	// The namespace of the AWS service that provides the resource. For a resource
	// provided by your own application or service, use custom-resource instead.
	// +kubebuilder:validation:Required
	ServiceNamespace *string `json:"serviceNamespace"`
	// A step scaling policy.
	//
	// This parameter is required if you are creating a policy and the policy type
	// is StepScaling.
	StepScalingPolicyConfiguration *StepScalingPolicyConfiguration `json:"stepScalingPolicyConfiguration,omitempty"`
	// A target tracking scaling policy. Includes support for predefined or customized
	// metrics.
	//
	// This parameter is required if you are creating a policy and the policy type
	// is TargetTrackingScaling.
	TargetTrackingScalingPolicyConfiguration *TargetTrackingScalingPolicyConfiguration `json:"targetTrackingScalingPolicyConfiguration,omitempty"`
}

// ScalingPolicyStatus defines the observed state of ScalingPolicy
type ScalingPolicyStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The CloudWatch alarms created for the target tracking scaling policy.
	// +kubebuilder:validation:Optional
	Alarms []*Alarm `json:"alarms,omitempty"`
	// The Unix timestamp for when the scaling policy was created.
	// +kubebuilder:validation:Optional
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// The Unix timestamp for when the scaling policy was created.
	// +kubebuilder:validation:Optional
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

// ScalingPolicy is the Schema for the ScalingPolicies API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ScalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScalingPolicySpec   `json:"spec,omitempty"`
	Status            ScalingPolicyStatus `json:"status,omitempty"`
}

// ScalingPolicyList contains a list of ScalingPolicy
// +kubebuilder:object:root=true
type ScalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScalingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ScalingPolicy{}, &ScalingPolicyList{})
}
