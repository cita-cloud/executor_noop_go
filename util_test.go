package main

import "testing"

func TestReadToml(t *testing.T) {
	config, err := readToml("example/test-chain-0/config.toml", "executor_evm")
	if err != nil {
		t.Fatal(err)
	}
	if config.ExecutorPort != 50002 {
		t.Errorf("config.ExecutorPort is %d, want 5002", config.ExecutorPort)
	}
}
