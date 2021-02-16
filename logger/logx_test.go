package logger

import (
	"os"
	"testing"
)

func init_() *LogX {
	logX := New()

	logFile, _ := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)

	logX.Out = logFile
	return logX
}
func TestLog(t *testing.T) {

	logX := init_()

	for i := 0; i < 1000000; i++ {
		//Log.Info("测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度");
		logX.Log("测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度测试一下速度")
	}

}
