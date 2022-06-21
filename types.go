package xworkflow

type TypeMeta struct {
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
}

type ObjectMeta struct {
	Name string `json:"name"`
}