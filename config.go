package main

type ExecutorConfig struct {
	ExecutorPort     uint16      `toml:"executor_port" json:"executor_port,omitempty"`
	EthCompatibility bool        `toml:"eth_compatibility" json:"eth_compatibility,omitempty"`
	DbPath           string      `json:"db_path" json:"db_path,omitempty"`
	SyncMode         string      `json:"sync_mode" json:"sync_mode,omitempty"`
	EnableMetrics    bool        `json:"enable_metrics" json:"enable_metrics,omitempty"`
	MetricsPort      uint16      `json:"metrics_port,omitempty"`
	MetricsBuckets   []float64   `json:"metrics_buckets,omitempty"`
	Domain           string      `json:"domain,omitempty"`
	LogConfig        interface{} `json:"log_config,omitempty"`
}
