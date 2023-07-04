package utils

import (
	"fmt"
	"math/big"
)

var b58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encoding(src string) {
	// 1.ascii码对应的值

	srcByte := []byte(src)
	fmt.Println(srcByte)

	// 转成十进制
	i := big.NewInt(0).SetBytes(srcByte)
	fmt.Println(i)

	var modSlice []byte
	//for i.Cmp(big.NewInt(0)) != 0 {
	for i.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0)
		i58 := big.NewInt(58)
		// 取余
		i.DivMod(i, i58, mod)
		// 将余数添加到数组中
		modSlice = append(modSlice, b58[mod.Int64()])
	}
}
