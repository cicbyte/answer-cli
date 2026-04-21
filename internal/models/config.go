package models

type AppConfig struct {
	Version string `yaml:"version"`

	Server struct {
		BaseURL string `yaml:"base_url"`
		Token   string `yaml:"token"`
	} `yaml:"server"`

	AI struct {
		Provider    string  `yaml:"provider"`
		BaseURL     string  `yaml:"base_url"`
		ApiKey      string  `yaml:"api_key"`
		Model       string  `yaml:"model"`
		MaxTokens   int     `yaml:"max_tokens"`
		Temperature float64 `yaml:"temperature"`
		Timeout     int     `yaml:"timeout"`
	} `yaml:"ai"`

	Output struct {
		Format string `yaml:"format"`
	} `yaml:"output"`

	Log struct {
		Level      string `yaml:"level"`
		MaxSize    int    `yaml:"max_size"`
		MaxBackups int    `yaml:"max_backups"`
		MaxAge     int    `yaml:"max_age"`
		Compress   bool   `yaml:"compress"`
	} `yaml:"log"`
}

func (c *AppConfig) GetServerURL() string {
	return c.Server.BaseURL
}

func (c *AppConfig) GetServerToken() string {
	return c.Server.Token
}

func (c *AppConfig) GetAIProvider() string   { return c.AI.Provider }
func (c *AppConfig) GetAIBaseURL() string    { return c.AI.BaseURL }
func (c *AppConfig) GetAIAPIKey() string     { return c.AI.ApiKey }
func (c *AppConfig) GetAIModel() string      { return c.AI.Model }
func (c *AppConfig) GetOutputFormat() string { return c.Output.Format }
