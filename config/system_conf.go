package config

import "gopkg.in/ini.v1"

func (c *Config) ReadSystem(filePath string) error {
	cfg, err := ini.Load(filePath)
	if err != nil {
		return err
	}

	err = cfg.Section("").MapTo(&c.SystemConfig)
	if err != nil {
		return err
	}

	return err
}

type SystemConfig struct {
	OnlineServer  string `ini:"online_server"`
	UserServer    string `ini:"user_server"`
	AuthServer    string `ini:"auth_server"`
	DetailServer  string `ini:"detail_server"`
	LogServer     string `ini:"log_server"`
	DyeServer     string `ini:"dye_server"`
	MyIp          string `ini:"my_ip"`
	RemoteDmIp    string `ini:"remote_dm_ip"`
	ThreadNum     string `ini:"thread_num"`
	RedisPassword string `ini:"redis_password"`
	CacheServer   string `ini:"cache_server"`
}
