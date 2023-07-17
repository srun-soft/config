package config

import "gopkg.in/ini.v1"

func (c *Config) ReadSrun(filePath string) error {
	cfg, err := ini.Load(filePath)
	if err != nil {
		return err
	}

	err = cfg.Section("").MapTo(&c.SrunConfig)
	if err != nil {
		return err
	}

	return err
}

type SrunConfig struct {
	Username            string `ini:"username"`
	Hostname            string `ini:"hostname"`
	Dbname              string `ini:"dbname"`
	Password            string `ini:"password"`
	Port                string `ini:"port"`
	DetailUsername      string `ini:"detail_username"`
	DetailHostname      string `ini:"detail_hostname"`
	DetailDbname        string `ini:"detail_dbname"`
	DetailPassword      string `ini:"detail_password"`
	DetailPort          string `ini:"detail_port"`
	LogUsername         string `ini:"log_username"`
	LogHostname         string `ini:"log_hostname"`
	LogDbname           string `ini:"log_dbname"`
	LogPassword         string `ini:"log_password"`
	LogPort             string `ini:"log_port"`
	StatisticsHostname  string `ini:"statistics_hostname"`
	StatisticsUsername  string `ini:"statistics_username"`
	StatisticsDbname    string `ini:"statistics_dbname"`
	StatisticsPassword  string `ini:"statistics_password"`
	StatisticsTcpPort   string `ini:"statistics_tcp_port"`
	StatisticsHttpPort  string `ini:"statistics_http_port"`
	StatisticsMysqlPort string `ini:"statistics_mysql_port"`
	InterfaceIp	    string `ini:"interface_ip"`
}
