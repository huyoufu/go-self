package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {

	file, _ := os.Open("example.yaml")
	decoder := yaml.NewDecoder(file)

	config := &Config{}
	err := decoder.Decode(config)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
}
func TestXml(t *testing.T) {

}
func TestPath(t *testing.T) {
	wd, _ := os.Getwd()
	fmt.Print(wd)
}
