package mockServer

import (
	"math/rand"
)

// algorithm from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

var charBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

const (
    letterIdxBits = 6                    // 6 bits to represent a char index (62 possible chars)
    letterIdxMask = 1<<letterIdxBits - 1 // bitmask - only get the last 6 bits
    letterIdxMax  = 63 / letterIdxBits   // number of random chars that can be generated with 63 bits
)

func GenerateCode(n int) string {
    b := make([]byte, n)
    // A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
    for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = rand.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(charBytes) {
            b[i] = charBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}