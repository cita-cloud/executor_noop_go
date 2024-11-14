package main

import (
	"github.com/BurntSushi/toml"
)

func readToml(path string, name string) (ExecutorConfig, error) {
	var result map[string]toml.Primitive
	md, err := toml.DecodeFile(path, &result)
	if err != nil {
		return ExecutorConfig{}, err
	}
	var config ExecutorConfig
	err = md.PrimitiveDecode(result[name], &config)
	if err != nil {
		return ExecutorConfig{}, err
	}
	return config, nil
}
