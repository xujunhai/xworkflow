package xworkflow

import "k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/intstr"

// Workflow 定义整个执行流资源
type Workflow struct {
	TypeMeta `json:",inline"`
	ObjectMeta `json:"metadata"`
	Spec WorkflowSpec `json:"spec"`
	Status            WorkflowStatus `json:"status,omitempty"`
}

// WorkflowSpec workflow详细
type WorkflowSpec struct {
	// 工作流起点
	Entrypoint string `json:"entrypoint,omitempty"`

	// 工作步骤模板（Steps,DAG,SCRIPT）
	Templates []Template `json:"templates,omitempty"`

	// 入口点的参数
	// e.g. {{workflow.parameters.param}}
	Arguments Arguments `json:"arguments,omitempty"`

	// 执行器配置
	Executor *ExecutorConfig `json:"executor,omitempty"`

	// 是一个模板引用，退出时执行。无论工作流成功、失败或错误。
	OnExit string `json:"onExit,omitempty"`

	// 指标收集
	Metrics *Metrics `json:"metrics,omitempty" protobuf:"bytes,32,opt,name=metrics"`

	// 所有模板的重试执行策略
	RetryStrategy *RetryStrategy `json:"retryStrategy,omitempty" protobuf:"bytes,37,opt,name=retryStrategy"`
}

//Template 工作流重用和组合执行单元
type Template struct {
	// Name 模板内容
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// Inputs 描述模板输入内容
	Inputs Inputs `json:"inputs,omitempty" protobuf:"bytes,5,opt,name=inputs"`

	// Outputs 描述模板输出
	Outputs Outputs `json:"outputs,omitempty" protobuf:"bytes,6,opt,name=outputs"`

	// DAG 模板
	DAG *DAGTemplate `json:"dag,omitempty" protobuf:"bytes,15,opt,name=dag"`

	// Data 数据模板
	Data *Data `json:"data,omitempty" protobuf:"bytes,39,opt,name=data"`

	// HTTP makes a HTTP request
	HTTP *HTTP `json:"http,omitempty" protobuf:"bytes,42,opt,name=http"`

	// Plugin is a plugin template
	Plugin *Plugin `json:"plugin,omitempty" protobuf:"bytes,43,opt,name=plugin"`

	// 模板任务重试策略
	RetryStrategy *RetryStrategy `json:"retryStrategy,omitempty" protobuf:"bytes,22,opt,name=retryStrategy"`

	// FailFast 快速失败
	FailFast *bool `json:"failFast,omitempty" protobuf:"varint,41,opt,name=failFast"`

	// Executor 执行器配置
	Executor *ExecutorConfig `json:"executor,omitempty" protobuf:"bytes,33,opt,name=executor"`

	// Metrics are a list of metrics emitted from this template
	Metrics *Metrics `json:"metrics,omitempty" protobuf:"bytes,35,opt,name=metrics"`

	// Memoize allows templates to use outputs generated from already executed templates
	Memoize *Memoize `json:"memoize,omitempty" protobuf:"bytes,37,opt,name=memoize"`

	// Timeout allows to set the total node execution timeout duration counting from the node's start time.
	// This duration also includes time in which the node spends in Pending state. This duration may not be applied to Step or DAG templates.
	Timeout string `json:"timeout,omitempty" protobuf:"bytes,38,opt,name=timeout"`
}

// Memoization enables caching for the Outputs of the template
type Memoize struct {
	// Key is the key to use as the caching key
	Key string `json:"key" protobuf:"bytes,1,opt,name=key"`
	// Cache sets and configures the kind of cache
	Cache *Cache `json:"cache" protobuf:"bytes,2,opt,name=cache"`
	// MaxAge is the maximum age (e.g. "180s", "24h") of an entry that is still considered valid. If an entry is older
	// than the MaxAge, it will be ignored.
	MaxAge string `json:"maxAge" protobuf:"bytes,3,opt,name=maxAge"`
}

// Cache is the configuration for the type of cache to be used
type Cache struct {
	// ConfigMap
}

type Arguments struct {
	Parameters []Parameter `json:"parameters,omitempty"`
}

// ExecutorConfig executor
type ExecutorConfig struct {
	// Parallel 并行执行最大值
	Parallel *int64 `json:"Parallel,omitempty"`
}

// RetryStrategy provides controls on how to retry a workflow step
type RetryStrategy struct {
	// Limit is the maximum number of attempts when retrying a container
	Limit *intstr.IntOrString `json:"limit,omitempty" protobuf:"varint,1,opt,name=limit"`

	// RetryPolicy is a policy of NodePhase statuses that will be retried
	RetryPolicy RetryPolicy `json:"retryPolicy,omitempty" protobuf:"bytes,2,opt,name=retryPolicy,casttype=RetryPolicy"`

	// Backoff is a backoff strategy
	Backoff *Backoff `json:"backoff,omitempty" protobuf:"bytes,3,opt,name=backoff,casttype=Backoff"`

	// Expression is a condition expression for when a node will be retried. If it evaluates to false, the node will not
	// be retried and the retry strategy will be ignored
	Expression string `json:"expression,omitempty" protobuf:"bytes,5,opt,name=expression"`
}

type RetryPolicy string

const (
	RetryPolicyAlways           RetryPolicy = "Always"
	RetryPolicyOnFailure        RetryPolicy = "OnFailure"
	RetryPolicyOnError          RetryPolicy = "OnError"
	RetryPolicyOnTransientError RetryPolicy = "OnTransientError"
)

// Backoff is a backoff strategy to use within retryStrategy
type Backoff struct {
	// Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. "2m", "1h")
	Duration string `json:"duration,omitempty" protobuf:"varint,1,opt,name=duration"`
	// Factor is a factor to multiply the base duration after each failed retry
	Factor *intstr.IntOrString `json:"factor,omitempty" protobuf:"varint,2,opt,name=factor"`
	// MaxDuration is the maximum amount of time allowed for the backoff strategy
	MaxDuration string `json:"maxDuration,omitempty" protobuf:"varint,3,opt,name=maxDuration"`
}