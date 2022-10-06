package config

import "testing"

type Log struct {
	LogLevel string `yaml:"logLevel"` // 日志级别，支持debug,info,warn,error,panic
	LogFile  string `yaml:"logFile"`  // 日志文件存放路径,如果为空，则输出到控制台
}

type MyConfig struct {
	Log Log    `yaml:"log"`
	Msg string `yaml:"msg"`
}

func TestLoad(t *testing.T) {
	var testConfig *MyConfig
	err := Load("test/config.yaml", &testConfig)
	if err != nil {
		t.Errorf("Load config failed: %v", err)
	}
	if testConfig.Log.LogLevel == "debug" && testConfig.Msg == "hello,world" {
		t.Logf("Load config succeed: %v", testConfig)
	} else {
		t.Errorf("Load config failed: %v", err)
	}

}
