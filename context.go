package xworkflow

import (
	"context"
	"sync"
)

// WorkflowContext 支持2级实体
type WorkflowContext struct {
	ctx context.Context //支持外部传入ctx 做时间控制
	mu sync.RWMutex
	data map[string]interface{} //支持2级对象存储 a.b
}


