package md5

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	dest := Encode("123123")
	fmt.Print(dest)

}
