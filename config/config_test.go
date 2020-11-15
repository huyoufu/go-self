package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	bytes, _ := ioutil.ReadFile("example.yaml")
	config := &Config{}
	yaml.Unmarshal(bytes, config)
	fmt.Println(config)
}
func TestPath(t *testing.T) {
	wd, _ := os.Getwd()
	fmt.Print(wd)
}
