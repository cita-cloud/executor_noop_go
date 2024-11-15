package main

type ExecutorConfig struct {
	ExecutorPort     uint16    `toml:"executor_port" json:"executor_port,omitempty"`
	EthCompatibility bool      `toml:"eth_compatibility" json:"eth_compatibility,omitempty"`
	DbPath           string    `toml:"db_path" json:"db_path" json:"db_path,omitempty"`
	SyncMode         string    `toml:"sync_mode" json:"sync_mode,omitempty"`
	EnableMetrics    bool      `toml:"enable_metrics" json:"enable_metrics,omitempty"`
	MetricsPort      uint16    `toml:"metrics_port" json:"metrics_port,omitempty"`
	MetricsBuckets   []float64 `toml:"metrics_buckets" json:"metrics_buckets,omitempty"`
	Domain           string    `toml:"domain" json:"domain,omitempty"`
	LogConfig        LogConfig `toml:"log_config" json:"log_config,omitempty"`
}

type LogConfig struct {
	Filter          string `json:"filter,omitempty"`
	MaxLevel        string `toml:"max_level" json:"max_level,omitempty"`
	RollingFilePath string `toml:"rolling_file_path" json:"rolling_file_path,omitempty"`
	ServiceName     string `toml:"service_name" json:"service_name,omitempty"`
}
