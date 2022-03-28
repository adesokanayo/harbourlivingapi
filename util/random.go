package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	cryptoRand "crypto/rand"
	"encoding/hex"
	"io"
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

func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(cryptoRand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
