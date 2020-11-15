package config

type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
}
type ServerConfig struct {
	Port          int    `yaml:"port"`
	Cors          bool   `yaml:"cors"`
	EnableSession bool   `yaml:"enableSession"`
	LogLevel      int    `yaml:"logLevel"`
	ServerName    string `yaml:"serverName"`
}
