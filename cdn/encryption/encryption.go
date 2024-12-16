package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
    "os"
)

func pad(data []byte) []byte {
	blockSize := aes.BlockSize
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func unpad(data []byte) ([]byte, error) {
	padding := data[len(data)-1]
	if int(padding) > len(data) {
		return nil, fmt.Errorf("invalid padding size")
	}
	return data[:len(data)-int(padding)], nil
}

func EncryptFile(inputFile, outputFile string, key []byte) error {
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	plaintext = pad(plaintext)

	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return fmt.Errorf("failed to generate IV: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create AES cipher: %v", err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	copy(ciphertext[:aes.BlockSize], iv)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	err = os.WriteFile(outputFile, ciphertext, 0644)
	if err != nil {
		return fmt.Errorf("failed to write encrypted file: %v", err)
	}

	return nil
}

func DecryptFile(inputFile, outputFile string, key []byte) error {
	ciphertext, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read encrypted file: %v", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create AES cipher: %v", err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = unpad(plaintext)
	if err != nil {
		return fmt.Errorf("failed to unpad decrypted data: %v", err)
	}

	err = os.WriteFile(outputFile, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("failed to write decrypted file: %v", err)
	}

	return nil
}

