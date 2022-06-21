package ex

import (
	"dictStore/task"
)

type Task interface {
	Id() string
	Run(ctx task.WorkflowContext,input task.Input) task.Output
}

type TaskA struct {

}

type TaskAReq struct {
	Uid string `json:"uid"`
	Imei string `json:"imei"`
	OrderId string `json:"order_id"`
}

type TaskAResp struct {
}

func (t TaskA) Id() string  {
	return "taskA"
}

func (t TaskA) Run(ctx task.WorkflowContext,input task.Input)  {
	//do something
}

