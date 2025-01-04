package secrets

import (
	"encoding/pem"
	"fmt"
	"os"
	"reflect"
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
	if err := os.WriteFile("api_private_encrypted.pem", pemData, 0644); err != nil {
		panic(err)
	}
	fmt.Println("PEM file encrypted successfully.")
}

func TestDecryptPem(t *testing.T) {
	pemData, err := os.ReadFile("api_private_encrypted.pem")
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		panic("failed to decode PEM block")
	}
	password := "your-password"
	_, err = decryptPEMBlock(block, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("PEM file decrypted successfully.")
}

func TestEncrypedPEMFromPath(t *testing.T) {
	data, err := os.ReadFile("api_private.pem")
	if err != nil {
		panic(err)
	}
	err = GenerateEncrypedPEMFromPath("api_private.pem", "api_private_encrypted.pem", "your-password")
	if err != nil {
		panic(err)
	}
	DecryptPEMFromPath("api_private_encrypted.pem", "api_private_decrypted.pem", "your-password")
	decryptedData, err := os.ReadFile("api_private_decrypted.pem")
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(data, decryptedData) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", data, decryptedData)
	}
}
