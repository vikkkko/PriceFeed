package utils

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

func StrToFloat(str string) float64 {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return v
}

func StrToInt64(str string) (int64, error) {
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func FloatToStr(f float64) string {
	v := strconv.FormatFloat(f, 'f', -1, 64)
	return v
}

func Uuid() string {
	return uuid.NewV4().String()
}
func IsImageExt(extName string) bool {
	var supportExtNames = map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".ico": true, ".svg": true, ".bmp": true, ".gif": true,
	}
	return supportExtNames[extName]
}

func GetSwapHash(swapType, sender string, created int64) string {
	return "0xswap" + hex.EncodeToString(
		crypto.Keccak256Hash([]byte(swapType+sender+strconv.FormatInt(created, 10))).Bytes())
}

func ToUpperList(list []string) []string {
	for i, _ := range list {
		list[i] = strings.ToUpper(list[i])
	}
	return list
}
