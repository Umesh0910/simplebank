package util

import (
	"math/rand"
	"strings"
)

func RandomInt(min, max int64) int64{
	randomInt:= rand.Int63n(max-min+1) + min
	return randomInt
}

func RandomString(n int) string{
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	
	var sb strings.Builder
	k := len(alphabets)

	for i:=0;i<n;i++{
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner()string{
	return RandomString(5)
}

func RandomBalance() int64{
	return RandomInt(0,2000)
}

func RandomCurrency() string{
	currency := []string{"USD","INR","CAD", "BTC"}
	n := len(currency)
	return currency[rand.Intn(n)]
}