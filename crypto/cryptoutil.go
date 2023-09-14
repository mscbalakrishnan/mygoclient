package mycrypto

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
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
		errStr := errors.New("error in creating New Cipher obj")
		fmt.Println(errStr.Error(), err)
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

func EncodeToString(msg string) string {
	encodedStr := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("Encoded String :", encodedStr)
	return encodedStr
}

func DecodeToString(msg string) string {
	fmt.Println("Input Message :", msg)
	decodedStr, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		errString := fmt.Errorf("error in decoding string %v", err.Error())
		fmt.Println(errString)
	}

	fmt.Println("Decoded String: ", string(decodedStr))
	return string(decodedStr)

}
