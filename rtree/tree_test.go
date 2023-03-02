package rtree

import (
	"fmt"
	"testing"
	"time"
)

func TestCleanup(t *testing.T) {
	newRoot := NewRoot()
	newRoot.AddNode("/user/:id/home/:key", nil)
	newRoot.AddNode("/user/:id/job/:key", nil)
	newRoot.AddNode("/user/:id/job/xxx", nil)
	newRoot.AddNode("/user/:id/wife/:yyy", nil)
	//newRoot.addNode("/user/:id/wife/:yyy",nil)
	newRoot.AddNode("/:user/:id/wife/:yyy", nil)
	newRoot.AddNode("/user/www/xww/yyy", nil)
	newRoot.AddNode("/user/www/xww/yyy1", nil)
	//newRoot.addNode("/user/www/xww/yyy",NotFoundHandler )

	start := time.Now().UnixNano()
	var test = 1
	for i := 0; i < 100000; i++ {
		newRoot.Search("/user/ceshi/job/yyy")
		//s := map1["username"]
		//fmt.Println(s)
	}
	end := time.Now().UnixNano()

	fmt.Println((end - start) / 1e6)
	fmt.Println(test)
	//newRoot.addNode("/user/ceshi/jobx/:id",NotFoundHandler )
	//fmt.Println(search)

}
