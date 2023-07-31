package config

import (
	"os"
	"testing"
)

func TestConfig_ReadSrun(t *testing.T) {
	// 创建临时文件
	tempFile, err := os.CreateTemp("", "srun.conf")
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
		username = srun
		hostname = 127.0.0.1
		dbname = srun4k
		password = srun4000@srun.com
		port = 3506
	`
	if _, err := tempFile.WriteString(testData); err != nil {
		t.Fatalf("Failed to write test data to temporary file: %v", err)
	}

	// 调用 ReadSrun 函数
	config := &Config{}
	err = config.ReadSrun(tempFile.Name())
	if err != nil {
		t.Fatalf("ReadSrun returned an unexpected error: %v", err)
	}

	// 在这里进行断言，验证读取到的配置是否符合预期
	expectedUsername := "srun"
	if config.Username != expectedUsername {
		t.Errorf("ReadSrun: expected Username = %s, got %s", expectedUsername, config.Username)
	}

	// 验证其他配置项...

	// 进行其他断言...
}
