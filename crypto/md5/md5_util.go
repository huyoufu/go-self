package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Encode(source string) (dest string) {

	md5 := md5.New()
	md5.Write([]byte(source))
	dest = hex.EncodeToString(md5.Sum(nil))
	return
}
