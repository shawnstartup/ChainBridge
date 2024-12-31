package secretes

import (
	"encoding/pem"
	"fmt"
	"os"
	"testing"
)

func TestEncryptPem(t *testing.T) {
	data, err := os.ReadFile("api_private.pem")
	if err != nil {
		panic(err)
	}
	password := "your-password"
	encryptedBlock, err := encryptPEMBlock(data, password)
	if err != nil {
		panic(err)
	}
	pemData := pem.EncodeToMemory(encryptedBlock)
	if err := os.WriteFile("biz_private_encrypted.pem", pemData, 0644); err != nil {
		panic(err)
	}
	fmt.Println("PEM file encrypted successfully.")
}

func TestDecryptPem(t *testing.T) {
	pemData, err := os.ReadFile("biz_private_encrypted.pem")
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		panic("failed to decode PEM block")
	}
	password := "your-password"
	decryptedData, err := decryptPEMBlock(block, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("PEM file decrypted successfully.")
	fmt.Println(string(decryptedData))
}
