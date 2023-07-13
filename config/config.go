package config

import (
	"fmt"
)

type Reader interface {
	ReadSrun(filePath string) error
	ReadSystem(filePath string) error
}

type Config struct {
	SrunConfig
	SystemConfig
}

const (
	SrunConf   = "/srun3/etc/srun.conf"
	SystemConf = "/srun3/etc/system.conf"
)

var (
	globalConfig *Config
	srunConf     = SrunConf
	systemConf   = SystemConf
)

// SetConfigPath 设置配置目录
func SetConfigPath(path string) {
	srunConf = fmt.Sprintf("%s/srun.conf", path)
	systemConf = fmt.Sprintf("%s/system.conf", path)
}

// GetConfig 获取配置
func GetConfig() (*Config, error) {
	if globalConfig != nil {
		return globalConfig, nil
	}
	config := &Config{}
	err := config.ReadSrun(srunConf)
	if err != nil {
		return nil, err
	}
	err = config.ReadSystem(systemConf)
	if err != nil {
		return nil, err
	}
	return config, nil
}
