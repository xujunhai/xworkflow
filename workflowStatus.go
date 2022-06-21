package xworkflow

// WorkflowStatus contains overall status information about a workflow
type WorkflowStatus struct {
	// Phase a simple, high-level summary of where the workflow is in its lifecycle.
	Phase WorkflowPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=WorkflowPhase"`

	// Time at which this workflow started
	StartedAt metav1.Time `json:"startedAt,omitempty" protobuf:"bytes,2,opt,name=startedAt"`

	// Time at which this workflow completed
	FinishedAt metav1.Time `json:"finishedAt,omitempty" protobuf:"bytes,3,opt,name=finishedAt"`

	// EstimatedDuration in seconds.
	EstimatedDuration EstimatedDuration `json:"estimatedDuration,omitempty" protobuf:"varint,16,opt,name=estimatedDuration,casttype=EstimatedDuration"`

	// Progress to completion
	Progress Progress `json:"progress,omitempty" protobuf:"bytes,17,opt,name=progress,casttype=Progress"`

	// A human readable message indicating details about why the workflow is in this condition.
	Message string `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`

	// Compressed and base64 decoded Nodes map
	CompressedNodes string `json:"compressedNodes,omitempty" protobuf:"bytes,5,opt,name=compressedNodes"`

	// Nodes is a mapping between a node ID and the node's status.
	Nodes Nodes `json:"nodes,omitempty" protobuf:"bytes,6,rep,name=nodes"`

	// Whether on not node status has been offloaded to a database. If exists, then Nodes and CompressedNodes will be empty.
	// This will actually be populated with a hash of the offloaded data.
	OffloadNodeStatusVersion string `json:"offloadNodeStatusVersion,omitempty" protobuf:"bytes,10,rep,name=offloadNodeStatusVersion"`

	// StoredTemplates is a mapping between a template ref and the node's status.
	StoredTemplates map[string]Template `json:"storedTemplates,omitempty" protobuf:"bytes,9,rep,name=storedTemplates"`

	// PersistentVolumeClaims tracks all PVCs that were created as part of the workflow.
	// The contents of this list are drained at the end of the workflow.
	PersistentVolumeClaims []apiv1.Volume `json:"persistentVolumeClaims,omitempty" protobuf:"bytes,7,rep,name=persistentVolumeClaims"`

	// Outputs captures output values and artifact locations produced by the workflow via global outputs
	Outputs *Outputs `json:"outputs,omitempty" protobuf:"bytes,8,opt,name=outputs"`

	// Conditions is a list of conditions the Workflow may have
	Conditions Conditions `json:"conditions,omitempty" protobuf:"bytes,13,rep,name=conditions"`

	// ResourcesDuration is the total for the workflow
	ResourcesDuration ResourcesDuration `json:"resourcesDuration,omitempty" protobuf:"bytes,12,opt,name=resourcesDuration"`

	// StoredWorkflowSpec stores the WorkflowTemplate spec for future execution.
	StoredWorkflowSpec *WorkflowSpec `json:"storedWorkflowTemplateSpec,omitempty" protobuf:"bytes,14,opt,name=storedWorkflowTemplateSpec"`

	// Synchronization stores the status of synchronization locks
	Synchronization *SynchronizationStatus `json:"synchronization,omitempty" protobuf:"bytes,15,opt,name=synchronization"`

	// ArtifactRepositoryRef is used to cache the repository to use so we do not need to determine it everytime we reconcile.
	ArtifactRepositoryRef *ArtifactRepositoryRefStatus `json:"artifactRepositoryRef,omitempty" protobuf:"bytes,18,opt,name=artifactRepositoryRef"`
}

func (ws *WorkflowStatus) IsOffloadNodeStatus() bool {
	return ws.OffloadNodeStatusVersion != ""
}

func (ws *WorkflowStatus) GetOffloadNodeStatusVersion() string {
	return ws.OffloadNodeStatusVersion
}

func (wf *Workflow) GetOffloadNodeStatusVersion() string {
	return wf.Status.GetOffloadNodeStatusVersion()
}
