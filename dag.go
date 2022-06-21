package xworkflow

// DAGTemplate is a template subtype for directed acyclic graph templates
type DAGTemplate struct {
	// Target are one or more names of targets to execute in a DAG
	Target string `json:"target,omitempty" protobuf:"bytes,1,opt,name=target"`

	// Tasks are a list of DAG tasks
	// +patchStrategy=merge
	// +patchMergeKey=name
	Tasks []DAGTask `json:"tasks" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=tasks"`

	// This flag is for DAG logic. The DAG logic has a built-in "fail fast" feature to stop scheduling new steps,
	// as soon as it detects that one of the DAG nodes is failed. Then it waits until all DAG nodes are completed
	// before failing the DAG itself.
	// The FailFast flag default is true,  if set to false, it will allow a DAG to run all branches of the DAG to
	// completion (either success or failure), regardless of the failed outcomes of branches in the DAG.
	// More info and example about this feature at https://github.com/argoproj/argo-workflows/issues/1442
	FailFast *bool `json:"failFast,omitempty" protobuf:"varint,3,opt,name=failFast"`
}

// DAGTask represents a node in the graph during DAG execution
type DAGTask struct {
	// Name is the name of the target
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// Name of template to execute
	Template string `json:"template,omitempty" protobuf:"bytes,2,opt,name=template"`

	// Inline is the template. Template must be empty if this is declared (and vice-versa).
	Inline *Template `json:"inline,omitempty" protobuf:"bytes,14,opt,name=inline"`

	// Arguments are the parameter and artifact arguments to the template
	Arguments Arguments `json:"arguments,omitempty" protobuf:"bytes,3,opt,name=arguments"`

	// TemplateRef is the reference to the template resource to execute.
	TemplateRef *TemplateRef `json:"templateRef,omitempty" protobuf:"bytes,4,opt,name=templateRef"`

	// Dependencies are name of other targets which this depends on
	Dependencies []string `json:"dependencies,omitempty" protobuf:"bytes,5,rep,name=dependencies"`

	// WithItems expands a task into multiple parallel tasks from the items in the list
	WithItems []Item `json:"withItems,omitempty" protobuf:"bytes,6,rep,name=withItems"`

	// WithParam expands a task into multiple parallel tasks from the value in the parameter,
	// which is expected to be a JSON list.
	WithParam string `json:"withParam,omitempty" protobuf:"bytes,7,opt,name=withParam"`

	// WithSequence expands a task into a numeric sequence
	WithSequence *Sequence `json:"withSequence,omitempty" protobuf:"bytes,8,opt,name=withSequence"`

	// When is an expression in which the task should conditionally execute
	When string `json:"when,omitempty" protobuf:"bytes,9,opt,name=when"`

	// ContinueOn makes argo to proceed with the following step even if this step fails.
	// Errors and Failed states can be specified
	ContinueOn *ContinueOn `json:"continueOn,omitempty" protobuf:"bytes,10,opt,name=continueOn"`

	// OnExit is a template reference which is invoked at the end of the
	// template, irrespective of the success, failure, or error of the
	// primary template.
	// DEPRECATED: Use Hooks[exit].Template instead.
	OnExit string `json:"onExit,omitempty" protobuf:"bytes,11,opt,name=onExit"`

	// Depends are name of other targets which this depends on
	Depends string `json:"depends,omitempty" protobuf:"bytes,12,opt,name=depends"`

	// Hooks hold the lifecycle hook which is invoked at lifecycle of
	// task, irrespective of the success, failure, or error status of the primary task
	Hooks LifecycleHooks `json:"hooks,omitempty" protobuf:"bytes,13,opt,name=hooks"`
}

// ContinueOn defines if a workflow should continue even if a task or step fails/errors.
// It can be specified if the workflow should continue when the pod errors, fails or both.
type ContinueOn struct {
	// +optional
	Error bool `json:"error,omitempty" protobuf:"varint,1,opt,name=error"`
	// +optional
	Failed bool `json:"failed,omitempty" protobuf:"varint,2,opt,name=failed"`
}