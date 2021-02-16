package config

const (
	ModelDev     = "dev"
	ModelTest    = "test"
	ModelProduct = "product"
)

type Config struct {
	Mode         string        `yaml:"mode"` //模式 可选为 dev,test,product
	ServerConfig *ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port          int    `yaml:"port"`
	Cors          bool   `yaml:"cors"`
	EnableSession bool   `yaml:"enableSession"`
	ServerName    string `yaml:"serverName"`
}
