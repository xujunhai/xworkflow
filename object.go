package xworkflow

import "encoding/json"

type Object struct {
	Value json.RawMessage `json:"-"`
}

func (i *Object) UnmarshalJSON(value []byte) error {
	return i.Value.UnmarshalJSON(value)
}

func (i Object) MarshalJSON() ([]byte, error) {
	return i.Value.MarshalJSON()
}

func (i Object) OpenAPISchemaType() []string {
	return []string{"object"}
}

func (i Object) OpenAPISchemaFormat() string { return "" }
