package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"
	"time"
)

const (
	defaultHeader string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	defaultSecret string = "www.jk1123.com"
)

/*  iss string//：Issuer，发行者 jwt签发者
sub string//：Subject，主题 jwt所面向的用户
aud string//：Audience，观众 接收jwt的一方
exp string//：Expiration time，过期时间 jwt的过期时间，这个过期时间必须要大于签发时间
nbf string//：Not before 定义在什么时间之前，该jwt都是不可用的.
iat int64//：Issued at，发行时间  jwt的签发时间
jti string//：JWT ID jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击*/

func NewPayload() map[string]interface{} {
	now_ := time.Now()
	now := now_.UnixNano() / 1e6
	exp := now_.Add(30*time.Minute).UnixNano() / 1e6
	payload := map[string]interface{}{
		"iss": "www.jk1123.com",
		"sub": "vip",
		"aud": "server",
		"exp": exp,
		"nbf": now,
		"iat": now,
	}
	return payload
}
func Sign(data map[string]interface{}) string {
	payload := NewPayload()
	for k, v := range data {
		payload[k] = v
	}
	bytes, _ := json.Marshal(payload)
	s := base64.StdEncoding.EncodeToString(bytes)
	s = strings.Replace(s, "=", "", -1)
	//fmt.Println(s)
	return defaultHeader + "." + s + "." + getHmacCode(defaultHeader+"."+s)
}
func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte(defaultSecret))
	io.WriteString(h, s)
	sum := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
