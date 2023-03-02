# go-self
Simple easy lightweight frame(简单 容易 轻量的 web 框架)
## 1.install
```go
go get -u github.com/huyoufu/go-self
```
## 2.start
```go
package main

import (
	"github.com/huyoufu/go-self"
)

func main() {
	app := self.NewServer()
	app.Port(80)
	self.Any("/", func(ctx self.Context) {
		ctx.WriteString("你好")
	})
	app.Start()
}

```
