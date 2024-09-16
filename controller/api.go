package controller

import (
	"crypto/sha1"
	"encoding/hex"
)

type ApiController struct{}

func (ac *ApiController) TransformSHA1(stringConvert string) string {
	encrypt := sha1.New()
	encrypt.Write([]byte(stringConvert))
	return hex.EncodeToString(encrypt.Sum(nil))
}
