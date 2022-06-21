package xworkflow

import (
	"fmt"
	"sort"
)

type MetricType string

const (
	MetricTypeGauge     MetricType = "Gauge"
	MetricTypeHistogram MetricType = "Histogram"
	MetricTypeCounter   MetricType = "Counter"
	MetricTypeUnknown   MetricType = "Unknown"
)

// Metrics are a list of metrics emitted from a Workflow/Template
type Metrics struct {
	// Prometheus is a list of prometheus metrics to be emitted
	Prometheus []*Prometheus `json:"prometheus" protobuf:"bytes,1,rep,name=prometheus"`
}

// Prometheus is a prometheus metric to be emitted
type Prometheus struct {
	// Name is the name of the metric
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Labels is a list of metric labels
	Labels []*MetricLabel `json:"labels,omitempty" protobuf:"bytes,2,rep,name=labels"`
	// Help is a string that describes the metric
	Help string `json:"help" protobuf:"bytes,3,opt,name=help"`
	// When is a conditional statement that decides when to emit the metric
	When string `json:"when,omitempty" protobuf:"bytes,4,opt,name=when"`
	// Gauge is a gauge metric
	Gauge *Gauge `json:"gauge,omitempty" protobuf:"bytes,5,opt,name=gauge"`
	// Histogram is a histogram metric
	Histogram *Histogram `json:"histogram,omitempty" protobuf:"bytes,6,opt,name=histogram"`
	// Counter is a counter metric
	Counter *Counter `json:"counter,omitempty" protobuf:"bytes,7,opt,name=counter"`
}

func (p *Prometheus) GetMetricLabels() map[string]string {
	labels := make(map[string]string)
	for _, label := range p.Labels {
		labels[label.Key] = label.Value
	}
	return labels
}

func (p *Prometheus) GetMetricType() MetricType {
	if p.Gauge != nil {
		return MetricTypeGauge
	}
	if p.Histogram != nil {
		return MetricTypeHistogram
	}
	if p.Counter != nil {
		return MetricTypeCounter
	}
	return MetricTypeUnknown
}

func (p *Prometheus) GetValueString() string {
	switch p.GetMetricType() {
	case MetricTypeGauge:
		return p.Gauge.Value
	case MetricTypeCounter:
		return p.Counter.Value
	case MetricTypeHistogram:
		return p.Histogram.Value
	default:
		return ""
	}
}

func (p *Prometheus) SetValueString(val string) {
	switch p.GetMetricType() {
	case MetricTypeGauge:
		p.Gauge.Value = val
	case MetricTypeCounter:
		p.Counter.Value = val
	case MetricTypeHistogram:
		p.Histogram.Value = val
	}
}

func (p *Prometheus) GetDesc() string {
	// This serves as a hash for the metric
	// TODO: Make sure this is what we want to use as the hash
	labels := p.GetMetricLabels()
	desc := p.Name + "{"
	for _, key := range sortedMapStringStringKeys(labels) {
		desc += key + "=" + labels[key] + ","
	}
	if p.Histogram != nil {
		sortedBuckets := p.Histogram.GetBuckets()
		sort.Float64s(sortedBuckets)
		for _, bucket := range sortedBuckets {
			desc += "bucket=" + fmt.Sprint(bucket) + ","
		}
	}
	desc += "}"
	return desc
}

func sortedMapStringStringKeys(in map[string]string) []string {
	var stringList []string
	for key := range in {
		stringList = append(stringList, key)
	}
	sort.Strings(stringList)
	return stringList
}

func (p *Prometheus) IsRealtime() bool {
	return p.GetMetricType() == MetricTypeGauge && p.Gauge.Realtime != nil && *p.Gauge.Realtime
}

// MetricLabel is a single label for a prometheus metric
type MetricLabel struct {
	Key   string `json:"key" protobuf:"bytes,1,opt,name=key"`
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

// Gauge is a Gauge prometheus metric
type Gauge struct {
	// Value is the value of the metric
	Value string `json:"value" protobuf:"bytes,1,opt,name=value"`
	// Realtime emits this metric in real time if applicable
	Realtime *bool `json:"realtime" protobuf:"varint,2,opt,name=realtime"`
}

// Histogram is a Histogram prometheus metric
type Histogram struct {
	// Value is the value of the metric
	Value string `json:"value" protobuf:"bytes,3,opt,name=value"`
	// Buckets is a list of bucket divisors for the histogram
	Buckets []Amount `json:"buckets" protobuf:"bytes,4,rep,name=buckets"`
}

func (in *Histogram) GetBuckets() []float64 {
	buckets := make([]float64, len(in.Buckets))
	for i, bucket := range in.Buckets {
		buckets[i], _ = bucket.Float64()
	}
	return buckets
}

// Counter is a Counter prometheus metric
type Counter struct {
	// Value is the value of the metric
	Value string `json:"value" protobuf:"bytes,1,opt,name=value"`
}