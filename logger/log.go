package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type FuncInfoHook struct {
	mu sync.Mutex
}

func (hook *FuncInfoHook) Fire(entry *logrus.Entry) error {
	hook.mu.Lock()
	defer hook.mu.Unlock()
	info := GenerateFuncInfo()
	entry.Data["func"] = info
	return nil
}

func (hook *FuncInfoHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

var Log = logrus.New()

func init() {

	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	Log.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	Log.Formatter = &logrus.TextFormatter{}

	Log.AddHook(&FuncInfoHook{})
	//Log.SetReportCaller(true)

}

func GenerateFuncInfo() string {
	//pc, file, line, ok := runtime.Caller(1)
	//.Print(pc,file,line,ok)
	//创建一个存放堆栈信息的 切片
	pc := make([]uintptr, 10) // at least 1 entry needed
	// skip参数 如果是0 标识当前函数 1代表上级调用者 2 代表更上级调用者 以此类推
	n := runtime.Callers(3, pc)

	//获取堆栈列表
	frames := runtime.CallersFrames(pc[:n])
	//我们只需要一个 无需遍历 取出第一个即可
	var f runtime.Frame
	more := true
	contain := false
	for f, more = frames.Next(); more; f, more = frames.Next() {
		//fmt.Println(f.Function)
		file := f.File
		if strings.Contains(file, "sirupsen/logrus") {
			contain = true
			continue
		}
		if contain {
			break
		}
	}

	funcName := f.Function
	file := f.File
	index := strings.LastIndex(file, "/")
	file = file[index+1:]

	funcLine := strconv.Itoa(f.Line)
	//time.Now().Format("2006-01-02 15:04:05")+" "
	printFuncStr := funcName + "@" + file + ":" + funcLine
	return printFuncStr
}
