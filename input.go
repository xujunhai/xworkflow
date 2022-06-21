package xworkflow

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Inputs are the mechanism for passing parameters, artifacts from one template to another
type Inputs struct {
	// 参数是作为输入的参数列表
	Parameters []Parameter `json:"parameters,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,1,opt,name=parameters"`
}

// Parameter indicate a passed string parameter to a service template with an optional default value
type Parameter struct {
	// Name is the parameter name
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// Default is the default value to use for an input parameter if a value was not supplied
	Default *AnyString `json:"default,omitempty" protobuf:"bytes,2,opt,name=default"`

	// Value is the literal value to use for the parameter.
	// If specified in the context of an input parameter, the value takes precedence over any passed values
	Value *AnyString `json:"value,omitempty" protobuf:"bytes,3,opt,name=value"`

	// ValueFrom is the source for the output parameter's value
	ValueFrom *ValueFrom `json:"valueFrom,omitempty" protobuf:"bytes,4,opt,name=valueFrom"`

	// GlobalName exports an output parameter to the global scope, making it available as
	// '{{workflow.outputs.parameters.XXXX}} and in workflow.status.outputs.parameters
	GlobalName string `json:"globalName,omitempty" protobuf:"bytes,5,opt,name=globalName"`

	// Enum holds a list of string values to choose from, for the actual value of the parameter
	Enum []AnyString `json:"enum,omitempty" protobuf:"bytes,6,rep,name=enum"`

	// Description is the parameter description
	Description *AnyString `json:"description,omitempty" protobuf:"bytes,7,opt,name=description"`
}

/*
AnyString It's JSON type is just string.
* It will unmarshall int64, int32, float64, float32, boolean, a plain string and represents it as string.
* It will marshall back to string - marshalling is not symmetric.
*/
type AnyString string

func ParseAnyString(val interface{}) AnyString {
	return AnyString(fmt.Sprintf("%v", val))
}

func AnyStringPtr(val interface{}) *AnyString {
	i := ParseAnyString(val)
	return &i
}

func (i *AnyString) UnmarshalJSON(value []byte) error {
	var v interface{}
	err := json.Unmarshal(value, &v)
	if err != nil {
		return err
	}
	switch v := v.(type) {
	case float64:
		*i = AnyString(strconv.FormatFloat(v, 'f', -1, 64))
	case float32:
		*i = AnyString(strconv.FormatFloat(float64(v), 'f', -1, 32))
	case int64:
		*i = AnyString(strconv.FormatInt(v, 10))
	case int32:
		*i = AnyString(strconv.FormatInt(int64(v), 10))
	case bool:
		*i = AnyString(strconv.FormatBool(v))
	case string:
		*i = AnyString(v)
	}
	return nil
}

func (i AnyString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(i))
}

func (i AnyString) String() string {
	return string(i)
}

// ValueFrom describes a location in which to obtain the value to a parameter
type ValueFrom struct {

	// JSONPath of a resource to retrieve an output parameter value from in resource templates
	JSONPath string `json:"jsonPath,omitempty" protobuf:"bytes,2,opt,name=jsonPath"`

	// JQFilter expression against the resource object in resource templates
	JQFilter string `json:"jqFilter,omitempty" protobuf:"bytes,3,opt,name=jqFilter"`

	// Selector (https://github.com/antonmedv/expr) that is evaluated against the event to get the value of the parameter. E.g. `payload.message`
	Event string `json:"event,omitempty" protobuf:"bytes,7,opt,name=event"`

	// Parameter reference to a dag task in which to retrieve an output parameter value from
	// (e.g. '{{steps.step.outputs.param}}')
	Parameter string `json:"parameter,omitempty" protobuf:"bytes,4,opt,name=parameter"`

	// Default specifies a value to be used if retrieving the value from the specified source fails
	Default *AnyString `json:"default,omitempty" protobuf:"bytes,5,opt,name=default"`

	// Expression, if defined, is evaluated to specify the value for the parameter
	Expression string `json:"expression,omitempty" protobuf:"bytes,8,rep,name=expression"`
}

func (p *Parameter) HasValue() bool {
	return p.Value != nil || p.Default != nil || p.ValueFrom != nil
}

func (p *Parameter) GetValue() string {
	if p.Value != nil {
		return p.Value.String()
	}
	if p.Default != nil {
		return p.Default.String()
	}
	return ""
}