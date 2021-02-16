package logger

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

type LogX struct {
	Level Level
	Out   io.Writer
	buff  *bytes.Buffer
}

func New() *LogX {
	x := &LogX{}
	x.buff = bytes.NewBuffer(make([]byte, 8192))
	pool := sync.Pool{}
	pool.New = func() interface{} {
		return nil
	}
	pool.Get()

	return x
}

func (l *LogX) Log(args ...interface{}) error {
	s := fmt.Sprint(args...)

	var b *bytes.Buffer = &bytes.Buffer{}

	b.WriteString(s)
	b.WriteString("\n")
	l.Out.Write(b.Bytes())
	//if (l.buff.Len()+len(s))>8192 {
	//
	//	b.WriteString(s)
	//	b.WriteString("\n")
	//	l.Out.Write(b.Bytes())
	//
	//	l.buff.Reset()
	//
	//}else{
	//	b.WriteString(s)
	//	b.WriteString("\n")
	//	l.Out.Write(b.Bytes())
	//}

	return nil
}
