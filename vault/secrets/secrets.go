package secrets

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func encryptPEMBlock(data []byte, password string) (*pem.Block, error) {
	key := sha256.Sum256([]byte(password))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	encryptedData := gcm.Seal(nonce, nonce, data, nil)
	return &pem.Block{Type: "ENCRYPTED DATA", Bytes: encryptedData}, nil
}

func decryptPEMBlock(block *pem.Block, password string) ([]byte, error) {
	key := sha256.Sum256([]byte(password))
	blockCipher, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(block.Bytes) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := block.Bytes[:nonceSize], block.Bytes[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func LoadEncrypedPrivateKeyFromPath(path string, password string) (*rsa.PrivateKey, error) {
	pemData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		panic("failed to decode PEM block")
	}
	context, err := decryptPEMBlock(block, password)
	if err != nil {
		panic(err)
	}
	pemBlock, _ := pem.Decode(context)
	if pemBlock == nil {
		return nil, fmt.Errorf("Could not read private key from[%s]. Please make sure the file in pem format, with headers and footers.(e.g. '-----BEGIN PRIVATE KEY-----' and '-----END PRIVATE KEY-----')", path)
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	return privateKey.(*rsa.PrivateKey), err
}

func GenerateEncrypedPEMFromPath(originFilepath string, targetFilePath, password string) error {
	data, err := os.ReadFile(originFilepath)
	if err != nil {
		panic(err)
	}
	encryptedBlock, err := encryptPEMBlock(data, password)
	if err != nil {
		panic(err)
	}
	pemData := pem.EncodeToMemory(encryptedBlock)
	if err := os.WriteFile(targetFilePath, pemData, 0644); err != nil {
		return err
	}
	return nil
}

func DecryptPEMFromPath(originFilepath string, targetFilePath, password string) error {
	pemData, err := os.ReadFile(originFilepath)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		panic("failed to decode PEM block")
	}
	decryptedData, err := decryptPEMBlock(block, password)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(targetFilePath, decryptedData, 0644); err != nil {
		return err
	}

	return nil
}
