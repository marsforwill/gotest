package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"sort"
	"time"
)

//ak sk鉴权
type signList []string

func (h signList) Len() int {
	return len(h)
}
func (h signList) Less(i, j int) bool {
	return h[i] <= h[j]
}
func (h signList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h signList) Join(sep string) string {
	res := ""
	for _, v := range h {
		res += v + sep
	}
	return res
}

type AuthInfo struct {
	AK    string `json:"ak,omitempty"`
	Nonce string `json:"nonce,omitempty"`
	TS    int64  `json:"ts,omitempty"`
	Sign  string `json:"sign,omitempty"`
	//Token string `json:"token,omitempty"`
}

func GenSignaksk(ak, sk string) *AuthInfo {
	var auth AuthInfo
	auth.AK = ak
	auth.Nonce = fmt.Sprintf("%d", time.Now().UnixNano()&0xFFFFFF)
	auth.TS = time.Now().Unix()

	var arr = signList{auth.AK, auth.Nonce, fmt.Sprintf("%d", auth.TS)}
	sort.Sort(arr)
	str := arr.Join("")
	fmt.Println("str", str)
	h := hmac.New(sha256.New, []byte(sk))
	h.Write([]byte(str))
	res := fmt.Sprintf("%x", h.Sum(nil))
	auth.Sign = res
	return &auth
}

//ak鉴权
func CreatData(ak, sk string) map[string]string {
	data := make(map[string]string)
	auth := GenSignaksk(ak, sk)
	data["ak"] = auth.AK
	data["nonce"] = auth.Nonce
	data["sign"] = auth.Sign
	data["ts"] = fmt.Sprintf("%v", auth.TS)
	return data
}

func main() {
	fmt.Println(CreatData("l1-d98a2a81-64558d59c9c1", "cd624e44f04789487fc074f5de460bbd"))
}
