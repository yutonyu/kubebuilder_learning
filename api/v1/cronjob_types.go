/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1  "k8s.io/api/core/v1"
)

// ConcurrencyPolicy describes how the job will be handled.
// Only one of the following concurrent policies may be specified.
// If none of the following policies is specified, the default one
// is AllowConcurrent.
// +kubebuilder:validation:Enum=Allow;Forbid;Replace
type ConcurrencyPolicy string

const (
	// AllowConcurrent allows Cronjobs to run concurrently.
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent forbids concurrent runs, skiping net run if previous
	// hasn't finished yet.
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent cancels currently running job and replaces it with a new one.
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// Next we define the types corrspodnng to actual kinds. CronJob and Cronjob list.
// 
// CronJobSpec defines the desired state of CronJob
// A deadline for starting jobs ( if we miss the deadline, we will just wait till the next scheduled time)
// What to do if multiple jobs would run at once
// A way to pause the running of a cornjob, incase something is wrong with it
// Limits on old job history
type CronJobSpec struct {
	//+kubebuilder:validation:MinLength=0
	// The scheduler in Cron format
	Schedule string `json:"schedule"`

	// +kubebuilder:validaiton:Minimum=0
	// Optional deadlien in seconds for startingthe job if it misses scheduled
	// time for any reason. Missed jobs executions will be counted as failed ones.
	// + optional
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds, omitempty"`

	// Specifies how to treat concurrent executions of a job.
	// Valid valiues are:
	// - "allow" default: allows controjbs ot run concurrently
	// - "forbid": formibds concurrent runs.
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

  // This flag tells the controller to suspend subsequent executions, it does
  // not apply to already started executions.  Defaults to false.
  // +optional
  Suspend *bool `json:"suspend,omitempty"`

  // Specifies the job that will be created when executing a CronJob.
	// TODO  where is JobTempalteSpec defined?
  // JobTemplate batchv1.JobTemplateSpec `json:"jobTemplate"`

  //+kubebuilder:validation:Minimum=0

  // The number of successful finished jobs to retain.
  // This is a pointer to distinguish between explicit zero and not specified.
  // +optional
  SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

  //+kubebuilder:validation:Minimum=0

  // The number of failed finished jobs to retain.
  // This is a pointer to distinguish between explicit zero and not specified.
  // +optional
  FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// We will keep a list of actively running jobs, as well as the last time
// that we successfully ran our job. Notice that we use metav1.Time
// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// A list of pointers to currently running jobs.
    // +optional
    Active []corev1.ObjectReference `json:"active,omitempty"`

    // Information when was the last time the job was successfully scheduled.
    // +optional
    LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`

}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`  // describes api version and kind
	metav1.ObjectMeta `json:"metadata,omitempty"` // holds things like name, namespce and labels.

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct { // A container for multiple CronJobs. It is the kind used in bulk operations, like List
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
