package mycrypto

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var key string = "this_must_be_of_32_byte_length!!"

func EncryptMessage(msg string) string {

	fmt.Println("Input message to encrypt :" + msg)
	fmt.Println("Input message to encrypt len : {} ", len(msg))

	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("blocksixe:{}", aes.BlockSize)

	dstmsgByte := make([]byte, len(msg))
	cipher.Encrypt(dstmsgByte, []byte(msg))
	encodedStr := hex.EncodeToString(dstmsgByte)
	fmt.Println("Encrypted message: " + encodedStr)
	return encodedStr
}

func DecryptMessage(encryptedMsg string) string {

	fmt.Println("Input message to decrypt :" + encryptedMsg)

	txt, _ := hex.DecodeString(encryptedMsg)

	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	msgByte := make([]byte, len(txt))
	cipher.Decrypt(msgByte, []byte(txt))

	msg := string(msgByte[:])

	fmt.Println("Decrypted message: " + msg)
	return msg
}

func GetMD5Hash(msg string) string {
	md5sum := md5.Sum([]byte(msg))
	return hex.EncodeToString(md5sum[:])
}
