package xworkflow

// Outputs 抽象任务执行后的输出 $.out.{字段值}
type Outputs struct {
	// Parameters holds the list of output parameters produced by a step
	Parameters []Parameter `json:"parameters,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,1,rep,name=parameters"`

	// Result holds the result (stdout) of a script template
	Result *string `json:"result,omitempty" protobuf:"bytes,3,opt,name=result"`

	// ExitCode holds the exit code of a script template
	ExitCode *string `json:"exitCode,omitempty" protobuf:"bytes,4,opt,name=exitCode"`
}