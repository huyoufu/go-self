package reflection

import (
	"runtime"
	"strings"
)

func GetFuncName(l int) string {
	//创建一个存放堆栈信息的 切片
	pc := make([]uintptr, 1) // at least 1 entry needed
	// skip参数 如果是0 标识当前函数 1代表上级调用者 2 代表更上级调用者 以此类推
	n := runtime.Callers(l, pc)

	//获取堆栈列表
	frames := runtime.CallersFrames(pc[:n])
	//我们只需要一个 无需遍历 取出第一个即可
	frame, _ := frames.Next()

	packageAndFunc := frame.Function
	index := strings.LastIndex(packageAndFunc, ".")

	funcName := packageAndFunc[index+1:]
	return funcName
}
