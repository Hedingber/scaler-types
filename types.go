package scaler_types

import (
	"encoding/json"
	"time"
)

type AutoScalerOptions struct {
	Namespace     string
	ScaleInterval time.Duration
	GroupKind     string
}

type ResourceScalerConfig struct {
	KubeconfigPath    string
	AutoScalerOptions AutoScalerOptions
	DLXOptions        DLXOptions
}

type DLXOptions struct {
	Namespace        string
	TargetNameHeader string
	TargetPathHeader string
	TargetPort       int
	ListenAddress    string
}

type ResourceScaler interface {
	SetScale(Resource, int) error
	GetResources() ([]Resource, error)
	GetConfig() (*ResourceScalerConfig, error)
}

type Resource struct {
	Name               string          `json:"name,omitempty"`
	ScaleResources     []ScaleResource `json:"scale_resources,omitempty"`
	LastScaleState     ScaleState      `json:"last_scale_state,omitempty"`
	LastScaleStateTime time.Time       `json:"last_scale_state_time,omitempty"`
}

func (r Resource) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(out)
}

type ScaleResource struct {
	MetricName string        `json:"metric_name,omitempty"`
	WindowSize time.Duration `json:"windows_size,omitempty"`
	Threshold  int           `json:"threshold,omitempty"`
}

func (sr ScaleResource) String() string {
	out, err := json.Marshal(sr)
	if err != nil {
		panic(err)
	}
	return string(out)
}

type ScaleState string

const (
	StartingScaleState ScaleState = "starting"
	FromZeroScaleState ScaleState = "fromZero"
	ToZeroScaleState   ScaleState = "toZero"
)
