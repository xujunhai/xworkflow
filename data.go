package xworkflow

// Data is a data template
type Data struct {
	// Source sources external data into a data template
	Source string `json:"source" protobuf:"bytes,1,opt,name=source"`

	// Transformation applies a set of transformations
	Transformation Transformation `json:"transformation" protobuf:"bytes,2,rep,name=transformation"`
}

type Transformation []TransformationStep

type TransformationStep struct {
	// Expression defines an expr expression to apply
	Expression string `json:"expression" protobuf:"bytes,1,opt,name=expression"`
}
