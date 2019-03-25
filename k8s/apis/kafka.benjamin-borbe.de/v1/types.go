package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Topic describes a database.
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TopicSpec `json:"spec"`
}

// TopicSpec is the spec for a Foo resource
type TopicSpec struct {
	Name              string `json:"name"`
	NumPartitions     int32  `json:"partitions,omitempty"`
	ReplicationFactor int16  `json:"replicas,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TopicList is a list of Topic resources
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Topic `json:"items"`
}
