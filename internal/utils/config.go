package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/cicbyte/answer-cli/internal/models"
	"go.yaml.in/yaml/v3"
)

var ConfigInstance = Config{}

type Config struct {
	HomeDir      string
	AppSeriesDir string
	AppDir       string
	ConfigDir    string
	ConfigPath   string
	LogDir       string
	LogPath      string
}

func (c *Config) GetHomeDir() string {
	if c.HomeDir != "" {
		return c.HomeDir
	}
	usr, err := user.Current()
	if err != nil {
		if dir, e := os.UserHomeDir(); e == nil {
			c.HomeDir = dir
			return c.HomeDir
		}
		panic(fmt.Sprintf("cannot get home dir: %v", err))
	}
	c.HomeDir = usr.HomeDir
	return c.HomeDir
}

func (c *Config) GetAppSeriesDir() string {
	if c.AppSeriesDir != "" {
		return c.AppSeriesDir
	}
	c.AppSeriesDir = filepath.Join(c.GetHomeDir(), ".cicbyte")
	return c.AppSeriesDir
}

func (c *Config) GetAppDir() string {
	if c.AppDir != "" {
		return c.AppDir
	}
	c.AppDir = filepath.Join(c.GetAppSeriesDir(), "answer-cli")
	return c.AppDir
}

func (c *Config) GetConfigDir() string {
	if c.ConfigDir != "" {
		return c.ConfigDir
	}
	c.ConfigDir = filepath.Join(c.GetAppDir(), "config")
	return c.ConfigDir
}

func (c *Config) GetConfigPath() string {
	if c.ConfigPath != "" {
		return c.ConfigPath
	}
	c.ConfigPath = filepath.Join(c.GetConfigDir(), "config.yaml")
	return c.ConfigPath
}

func (c *Config) GetLogDir() string {
	if c.LogDir == "" {
		c.LogDir = filepath.Join(c.GetAppDir(), "logs")
	}
	return c.LogDir
}

func (c *Config) GetLogPath() string {
	if c.LogPath == "" {
		now := time.Now().Format("20060102")
		c.LogPath = filepath.Join(c.GetLogDir(), fmt.Sprintf("answer-cli_%s.log", now))
	}
	return c.LogPath
}

func (c *Config) LoadConfig() *models.AppConfig {
	configPath := c.GetConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := GetDefaultConfig()
		if data, err := yaml.Marshal(defaultConfig); err == nil {
			_ = os.WriteFile(configPath, data, 0644)
		}
		return defaultConfig
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return GetDefaultConfig()
	}

	var config models.AppConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return GetDefaultConfig()
	}

	return &config
}

func (c *Config) SaveConfig(config *models.AppConfig) {
	configPath := c.GetConfigPath()
	data, err := yaml.Marshal(config)
	if err != nil {
		return
	}
	os.WriteFile(configPath, data, 0644)
}

func GetDefaultConfig() *models.AppConfig {
	config := &models.AppConfig{}

	config.Version = "0.1.0"

	config.Server.BaseURL = ""
	config.Server.Token = ""

	config.AI.Provider = "ollama"
	config.AI.BaseURL = "http://localhost:11434/v1"
	config.AI.Model = "gemma4:e4b"
	config.AI.MaxTokens = 2048
	config.AI.Temperature = 0.8
	config.AI.Timeout = 60

	config.Output.Format = "table"

	config.Log.Level = "info"
	config.Log.MaxSize = 10
	config.Log.MaxBackups = 30
	config.Log.MaxAge = 30
	config.Log.Compress = true

	return config
}
