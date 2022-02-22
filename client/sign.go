package client

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"sort"
	"strings"
	"time"
)

const ISO_FORMAT = "2006-01-02T15:04:05.000Z"

func GetISOTime() string {
	return time.Now().UTC().Format(ISO_FORMAT)
}

//使用已经过sha1加密的SecretKey进行签名
func Sign(timestamp, method, path string, para map[string]interface{}, SecretKey string) string {
	input := timestamp + method + path + buildSortParam(para)
	log.Debug().Msg(input)

	hash := hmac.New(sha256.New, []byte(SecretKey)) //1:HMAC256加密
	hash.Write([]byte(input))
	sum := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(sum) //2: base64编码
}

//使用未经过sha1加密的SecretKey进行签名
func Sign2(timestamp, method, path string, para map[string]interface{}, OriSecretKey string) string {
	bytes := sha1.Sum([]byte(OriSecretKey))   //1:对网站的原始SecretKey进行sha1加密
	SecretKey := hex.EncodeToString(bytes[:]) //2:转成字符串

	return Sign(timestamp, method, path, para, SecretKey)
}

func buildSortParam(para map[string]interface{}) string {
	if para == nil || len(para) == 0 {
		return ""
	}

	var str []string
	for k, v := range para {
		s, ok := v.(string)
		if ok { //本身是string，不需要再json编码了
			str = append(str, k+"="+s)
		} else {
			data, _ := json.Marshal(v)
			str = append(str, k+"="+string(data))
		}
	}

	sort.Strings(str)
	return strings.Join(str, "&")
}
