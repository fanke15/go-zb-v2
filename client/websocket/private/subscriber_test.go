package private

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
	"testing"
	"zb_sdk_go/client"
)

func TestGetISOTime(t *testing.T) {
	fmt.Println(client.GetISOTime())
}
func TestSing(t *testing.T) {
	ts := client.GetISOTime()
	ts = "2021-03-31T06:34:44.506Z"
	fmt.Println(ts)
	fmt.Println(client.Sign2(ts, "GET", "login", nil, SecretKey))
}

func TestSign2(t *testing.T) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = 1
	params["convertUnit"] = "usdt"
	ts := "2021-04-19T06:21:13.089Z"
	fmt.Println(ts)
	fmt.Println(client.Sign2(ts, "GET", "/Server/api/v2/Fund/getAccount", params, SecretKey))
}
func TestSign2_1(t *testing.T) {
	params := make(map[string]interface{})
	params["orderIds"] = []int64{6790106477685121024, 6790106476573630469}
	params["symbol"] = "ETH_USDT"
	ts := "2021-04-20T03:16:11.957Z"
	fmt.Println(ts)
	fmt.Println(client.Sign2(ts, "POST", "/Server/api/v2/trade/batchCancelOrder", params, SecretKey))
}

func TestStringInterface(t *testing.T) {
	para := make(map[string]interface{})
	para["a1"] = "a1"
	para["a2"] = "a2"
	para["aa2"] = "aa2"
	para["a0"] = "a0"
	para["A0"] = "A0"
	para["b1"] = 11
	para["c1"] = 12.3
	para["d1"] = big.NewInt(1)
	subMap := make(map[string]string)
	subMap["aa1"] = "aa1"
	para["e1"] = subMap

	var str []string
	for k, v := range para {
		str = append(str, k+"="+fmt.Sprint(v))
	}
	fmt.Println(str)
	sort.Strings(str)
	fmt.Println(str)
	fmt.Println(strings.Join(str, "&"))
}

func TestString(t *testing.T) {
	fmt.Println(fmt.Sprint(1))
	fmt.Println(fmt.Sprint("str1"))
	fmt.Println(fmt.Sprint(1.1))
	fmt.Println(fmt.Sprint([]int{1, 2}))

	str := make(map[string]interface{})
	str["str1"] = "str1"
	str["int2"] = 2
	fmt.Println(fmt.Sprint(str))
	fmt.Println(fmt.Sprintf("%v", str))
}
