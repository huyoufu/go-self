package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	/*sum2561 := sha256.Sum256([]byte("哈哈哈"))
	fmt.Println(string(sum2561[:]))
	s := base64.StdEncoding.EncodeToString([]byte(`{"alg": "HS256", "typ": "JWT"}"`))
	fmt.Println(s)*/

	start := time.Now().UnixNano()
	for i := 0; i < 10000; i++ {
		Sign(map[string]interface{}{"name": "xiaohu"})
		//fmt.Println(sign)
	}
	end := time.Now().UnixNano()

	fmt.Println((end - start) / 1e6)

}
