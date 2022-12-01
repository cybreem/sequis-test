package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"
	"net/http"

	"github.com/cybreem/sequis-test/app"
	"github.com/gin-gonic/gin"
	sdkLogging "github.com/h4lim/go-sdk/logging"
	sdkUtils "github.com/h4lim/go-sdk/utils"
)

const (
	ENVIRONMENT_PRODUCTION = "production"

	INITIAL_VECTOR_STAGING    = "xOL9fnf1wsGw9iyn"
	INITIAL_KEY_STAGING       = "puiRpYZApan3L5E3uKCT2ZXcCiqGlg6C"
	INITIAL_VECTOR_PRODUCTION = "xOL9fnf1wsGw9iyn"
	INITIAL_KEY_PRODUCTION    = "puiRpYZApan3L5E3uKCT2ZXcCiqGlg6C"
)

func GeneratePassword() string {

	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, "+
			"Authorization, accept, Origin, Cache-Control, X-Requested-With, Access-ID,"+
			"Host, Connection, Pragma, Cache-Control, sec-ch-ua-mobile, User-Agent, sec-ch-ua, Sec-Fetch-Site, Sec-Fetch-Mode, Sec-Fetch-Dest, Referer")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func AESEncrypt(src string) string {

	vector := []byte(INITIAL_VECTOR_PRODUCTION)
	key := []byte(INITIAL_KEY_PRODUCTION)
	if sdkUtils.GetRunMode() != ENVIRONMENT_PRODUCTION {
		vector = []byte(INITIAL_VECTOR_STAGING)
		key = []byte(INITIAL_KEY_STAGING)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		app.Log.Error(sdkLogging.INTERNAL, "Error occurred %s ", err.Error())
		return ""
	}

	ecb := cipher.NewCBCEncrypter(block, vector)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())

	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	strEncrypted := base64.StdEncoding.EncodeToString(crypted)

	return strEncrypted
}
func AESDecrypt(params string) string {

	vector := []byte(INITIAL_VECTOR_PRODUCTION)
	key := []byte(INITIAL_KEY_PRODUCTION)
	if sdkUtils.GetRunMode() != ENVIRONMENT_PRODUCTION {
		vector = []byte(INITIAL_VECTOR_STAGING)
		key = []byte(INITIAL_KEY_STAGING)
	}

	crypt, _ := base64.StdEncoding.DecodeString(params)

	block, err := aes.NewCipher(key)
	if err != nil {
		app.Log.Error(sdkLogging.INTERNAL, "Error occurred %s ", err.Error())
	}

	if len(crypt) == 0 {
		app.Log.Error(sdkLogging.INTERNAL, "plain content empty")
	}

	ecb := cipher.NewCBCDecrypter(block, vector)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return string(PKCS5Trimming(decrypted))
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {

	padding := blockSize - len(cipherText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, paddingText...)
}

func PKCS5Trimming(encrypt []byte) []byte {

	padding := encrypt[len(encrypt)-1]

	return encrypt[:len(encrypt)-int(padding)]
}

func DecodeBase64(logID string, param string) string {

	byteDecrypt, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		app.Log.Errorf(logID, "Error occurred %s ", err.Error())
		return ""
	}

	return string(byteDecrypt)
}
