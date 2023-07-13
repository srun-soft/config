package config

import (
	"flag"
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

// GetConfig 获取配置
func GetConfig(path string) (*Config, error) {
	mode := flag.String("mode", "prod", "")
	flag.Parse()
	if *mode != "prod" {
		srunConf = fmt.Sprintf("%s/srun.conf", path)
		systemConf = fmt.Sprintf("%s/system.conf", path)
	}
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
