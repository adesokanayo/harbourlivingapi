package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabeth = "abcdefjhijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomTime(min, max int64) time.Time {
	year := RandomInt(1988, 2020)
	return time.Date(int(year), 10, 01, 0, 0, 0, 0, time.UTC)
}

//RandomString generates random string of lenght n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabeth)
	for i := 0; i < n; i++ {
		c := alphabeth[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()

}

func RandomName() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(10) + "@gmail.com"
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}
