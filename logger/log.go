package logger

import (
	"log"
	"os"
	"runtime"
	"strings"

	"strconv"
)

const (
	Level_fetal = iota //fetal级别
	Level_error        //error级别
	Level_info         //info级别
	Level_debug        //debug级别
)

var log_level = Level_info

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func ChangeLevel(level int) {
	log_level = level
}

func Debug(msg string) {
	if log_level >= Level_debug {
		log.Print(msg)
	}
}
func DebugF(partten string, msg ...interface{}) {
	if log_level >= Level_debug {
		log.Printf(partten, msg)
	}
}
func Info(msg string) {
	if log_level >= Level_info {
		log.Print(msg)
	}
}
func InfoF(partten string, msg ...interface{}) {
	if log_level >= Level_info {
		log.Printf(partten, msg)
	}
}
func Error(msg string) {
	if log_level >= Level_error {
		log.Print(msg)
	}
}
func ErrorF(partten string, msg ...interface{}) {
	if log_level >= Level_error {
		log.Printf(partten, msg)
	}
}
func Fetal(msg string) {
	if log_level >= Level_fetal {
		log.Print(msg)
	}
	os.Exit(1)
}
func FetalF(partten string, msg ...interface{}) {
	if log_level >= Level_fetal {
		log.Printf(partten, msg)
	}
	os.Exit(1)
}
func generateFuncInfo() string {
	//pc, file, line, ok := runtime.Caller(1)
	//.Print(pc,file,line,ok)
	//创建一个存放堆栈信息的 切片
	pc := make([]uintptr, 1) // at least 1 entry needed
	// skip参数 如果是0 标识当前函数 1代表上级调用者 2 代表更上级调用者 以此类推
	n := runtime.Callers(2, pc)
	//获取堆栈列表
	frames := runtime.CallersFrames(pc[:n])
	//我们只需要一个 无需遍历 取出第一个即可
	frame, _ := frames.Next()

	packageAndFunc := frame.Function
	index := strings.LastIndex(packageAndFunc, ".")
	pkgName := packageAndFunc[:index]
	funcName := packageAndFunc[index+1:]

	file := frame.File

	funcFile := strings.Split(file, pkgName)[1]

	funcLine := strconv.Itoa(frame.Line)

	printFuncStr := pkgName + funcFile + ":" + funcLine + "-" + funcName
	return printFuncStr
}
