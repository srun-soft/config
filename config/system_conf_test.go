package config

import (
	"os"
	"testing"
)

func TestConfig_ReadSystem(t *testing.T) {
	// 创建临时文件
	tempFile, err := os.CreateTemp("", "system.conf")
	if err != nil {
		t.Fatalf("Failed to create tempFile: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tempFile.Name()) // 测试完成后删除临时文件

	// 写入测试数据到临时文件
	testData := `
		online_server = 127.0.0.1
		my_ip = 127.0.0.1
		redis_password = srun_3000@redis
	`
	if _, err := tempFile.WriteString(testData); err != nil {
		t.Fatalf("Failed to write test data to temporary file: %v", err)
	}

	// 调用 ReadSrun 函数
	config := &Config{}
	err = config.ReadSystem(tempFile.Name())
	if err != nil {
		t.Fatalf("ReadSrun returned an unexpected error: %v", err)
	}

	// 在这里进行断言，验证读取到的配置是否符合预期
	expectedOnlineServer := "127.0.0.1"
	if config.OnlineServer != expectedOnlineServer {
		t.Errorf("ReadSystem: expected OnlineServer = %s, got %s", expectedOnlineServer, config.OnlineServer)
	}
}
