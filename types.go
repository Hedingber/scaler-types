package scaler_types

import "time"

type AutoScalerOptions struct {
	Namespace     string
	ScaleInterval time.Duration
	ScaleWindow   time.Duration
	MetricName    string
	Threshold     int64
}

type PollerOptions struct {
	MetricInterval time.Duration
	MetricName     string
	Namespace      string
	GroupKind      string
}

type ResourceScalerConfig struct {
	KubeconfigPath    string
	AutoScalerOptions AutoScalerOptions
	PollerOptions     PollerOptions
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

type Resource string
