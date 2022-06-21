package xworkflow

type Task struct {
	Id string `json:"id"`
	Operator Operator `json:"operator"`
	State int `json:"state"`
}