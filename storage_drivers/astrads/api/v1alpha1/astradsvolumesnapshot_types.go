/*
Copyright 2021 NetApp Inc..

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetAppVolumeConditionType indicates the type of condition occurred on a AstraDSVolumeSnapshot
type AstraDSVolumeSnapshotConditionType string

const (
	// Is the snapshot ready to be used / done
	AstraDSVolumeSnapshotReady AstraDSVolumeSnapshotConditionType = "Ready"
)

// AstraDSVolumeSnapshot contains the condition information for a AstraDSVolumeSnapshot
type AstraDSVolumeSnapshotCondition struct {
	// Type of NetAppVolume condition.
	Type AstraDSVolumeSnapshotConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// AstraDSVolumeSnapshotSpec defines the desired state of AstraDSVolumeSnapshot
type AstraDSVolumeSnapshotSpec struct {
	VolumeName string `json:"volumeName"`
	Cluster    string `json:"cluster"`
}

// AstraDSVolumeSnapshotStatus defines the observed state of AstraDSVolumeSnapshot
type AstraDSVolumeSnapshotStatus struct {
	Conditions   []AstraDSVolumeSnapshotCondition `json:"conditions,omitempty"`
	VolumeName   string                           `json:"volumeName,omitempty"`
	CreationTime *metav1.Time                     `json:"creationTime,omitempty"`
	RestoreSize  int64                            `json:"restoreSize,omitempty"`
	VolumeUUID   string                           `json:"volumeUUID,omitempty"`
	ReadyToUse   bool                             `json:"readyToUse,omitempty"`
	Cluster      string                           `json:"cluster,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=adsvs,categories={ads,all}
// AstraDSVolumeSnapshot is the Schema for the astradsvolumesnapshots API
type AstraDSVolumeSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AstraDSVolumeSnapshotSpec   `json:"spec,omitempty"`
	Status AstraDSVolumeSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AstraDSVolumeSnapshotList contains a list of AstraDSVolumeSnapshot
type AstraDSVolumeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AstraDSVolumeSnapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AstraDSVolumeSnapshot{}, &AstraDSVolumeSnapshotList{})
}
