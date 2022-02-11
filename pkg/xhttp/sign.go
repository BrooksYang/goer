package xhttp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Sign(params map[string]interface{}, secretKey string) (sign string) {
	data := FormatURLParam(params)
	data = fmt.Sprintf("%s&secret_key=%s", data, secretKey)

	h := md5.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)

	return hex.EncodeToString(hash)
}

func VerifySign(params map[string]interface{}, secretKey string, sign string) bool {
	resign := Sign(params, secretKey)

	return sign == resign
}
