package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ChainSafe/ChainBridge/vault/secretes"
)

type CoSignerConverter struct {
	Config CoSignerConfig
}

type CoSignerConfig struct {
	ApiPubKey string `comment:"apiPubKey"`
}

type CoSignerCallBack struct {
	Timestamp  string `json:"timestamp"`
	Sig        string `json:"sig"`
	Key        string `json:"key"`
	BizContent string `json:"bizContent"`
	RsaType    string `json:"rsaType"`
	AesType    string `json:"aesType"`
}

func (c *CoSignerConverter) RequestConvert(d CoSignerCallBack, bizPrivateKey *rsa.PrivateKey) (string, error) {
	responseStringMap := map[string]string{
		"key":        d.Key,
		"timestamp":  d.Timestamp,
		"bizContent": d.BizContent,
	}
	// Verify sign
	verifyRet := secretes.VerifySignWithRSA(serializeParams(responseStringMap), d.Sig, c.Config.ApiPubKey)
	if !verifyRet {
		return "", errors.New("CoSignerCallBack signature verification failed")
	}
	// Use your RSA private key to decrypt response's aesKey and aesIv

	var plaintext []byte
	if d.RsaType == secretes.ECB_OAEP {
		plaintext, _ = secretes.DecryptWithOAEP(d.Key, bizPrivateKey)
	} else {
		plaintext, _ = secretes.DecryptWithRSA(d.Key, bizPrivateKey)
	}
	resAesKey := plaintext[:32]
	resAesIv := plaintext[32:]
	// Use AES to decrypt bizContent
	ciphertext, _ := base64.StdEncoding.DecodeString(d.BizContent)
	var callBackContent []byte
	if d.AesType == secretes.GCM {
		callBackContent, _ = secretes.NewGCMDecrypter(resAesKey, resAesIv, ciphertext)
	} else {
		callBackContent, _ = secretes.NewCBCDecrypter(resAesKey, resAesIv, ciphertext)
	}
	return string(callBackContent), nil
}

type CoSignerResponse struct {
	Approve bool   `json:"approve"`
	TxKey   string `json:"txKey"`
}

// It has been Deprecated,Please use convertCoSignerResponseWithNewCryptoType
func (c *CoSignerConverter) ResponseConverter(d any, bizPrivateKey *rsa.PrivateKey) (map[string]string, error) {
	// Use AES to encrypt request data
	aesKey := make([]byte, 32)
	rand.Read(aesKey)
	aesIv := make([]byte, 16)
	rand.Read(aesIv)
	// Create params map
	params := map[string]string{
		"timestamp": strconv.FormatInt(time.Now().UnixMilli(), 10),
		"code":      "200",
		"message":   "SUCCESS",
	}
	if d != nil {
		payLoad, _ := json.Marshal(d)
		data := string(payLoad)
		encryptBizContent, err := secretes.EncryContentWithAES(data, aesKey, aesIv)
		if err != nil {
			return nil, err
		}
		params["bizContent"] = encryptBizContent
	}

	// Use Safeheron RSA public key to encrypt request's aesKey and aesIv
	encryptedKeyAndIv, err := secretes.EncryptWithRSA(append(aesKey, aesIv...), c.Config.ApiPubKey)
	if err != nil {
		return nil, err
	}
	params["key"] = encryptedKeyAndIv

	// Sign the request data with your RSA private key
	signature, err := secretes.SignParamsWithRSA(serializeParams(params), bizPrivateKey)
	if err != nil {
		return nil, err
	}
	params["sig"] = signature
	return params, nil
}

func (c *CoSignerConverter) ResponseConverterWithNewCryptoType(d any, bizPrivateKey *rsa.PrivateKey) (map[string]string, error) {
	// Use AES to encrypt request data
	aesKey := make([]byte, 32)
	rand.Read(aesKey)
	aesIv := make([]byte, 16)
	rand.Read(aesIv)
	// Create params map
	params := map[string]string{
		"timestamp": strconv.FormatInt(time.Now().UnixMilli(), 10),
		"code":      "200",
		"message":   "SUCCESS",
	}
	if d != nil {
		payLoad, _ := json.Marshal(d)
		data := string(payLoad)
		encryptBizContent, err := secretes.EncryContentWithAESGCM(data, aesKey, aesIv)
		if err != nil {
			return nil, err
		}
		params["bizContent"] = encryptBizContent
	}

	// Use Safeheron RSA public key to encrypt request's aesKey and aesIv
	encryptedKeyAndIv, err := secretes.EncryptWithOAEP(append(aesKey, aesIv...), c.Config.ApiPubKey)
	if err != nil {
		return nil, err
	}
	params["key"] = encryptedKeyAndIv

	// Sign the request data with your RSA private key
	signature, err := secretes.SignParamsWithRSA(serializeParams(params), bizPrivateKey)
	if err != nil {
		return nil, err
	}
	params["sig"] = signature
	params["rsaType"] = secretes.ECB_OAEP
	params["aesType"] = secretes.GCM
	return params, nil
}

func serializeParams(params map[string]string) string {
	// Sort by key and serialize all request param into apiKey=...&bizContent=... format
	var data []string
	for k, v := range params {
		data = append(data, strings.Join([]string{k, v}, "="))
	}
	sort.Strings(data)
	return strings.Join(data, "&")
}
