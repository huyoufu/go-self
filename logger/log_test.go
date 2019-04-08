package logger

import (
	"testing"
)

func TestLog(t *testing.T) {

	Debug("debug你好")
	Info("info你好")
	Error("error你好")

}
