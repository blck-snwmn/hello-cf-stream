package main

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type key struct {
	ID      string `json:"id"`
	Pem     string `json:"pem"`
	JWK     string `json:"jwk"`
	Created string `json:"created"`
}

type keys struct {
	Result   key      `json:"result"`
	Success  bool     `json:"success"`
	Errors   []string `json:"errors"`
	Messages []string `json:"messages"`
}

type accessRule struct {
	Type   string `json:"type"`
	Action string `json:"action"`
}

type claims struct {
	jwt.RegisteredClaims
	KID         string       `json:"kid"`
	AccessRules []accessRule `json:"accessRules,omitempty"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <jwt token>")
		os.Exit(1)
	}
	videoID := os.Args[1]
	fmt.Printf("video id: %s\n", videoID)

	f, err := os.Open("keys.json")
	if err != nil {
		panic(err)
	}

	var keys keys
	err = json.NewDecoder(f).Decode(&keys)
	if err != nil {
		panic(err)
	}
	// {
	// 	k, err := base64.URLEncoding.DecodeString(keys.Result.JWK)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(k))
	// }

	p, err := base64.URLEncoding.DecodeString(keys.Result.Pem)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode([]byte(p))
	if block == nil {
		panic("failed to parse PEM block containing private key")
	}
	if block.Type != "RSA PRIVATE KEY" {
		panic("invalid PEM type")
	}
	k, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	jwt := jwt.NewWithClaims(jwt.SigningMethodRS256, claims{
		KID: keys.Result.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   videoID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
	})
	jwt.Header["kid"] = keys.Result.ID

	ss, err := jwt.SignedString(k)
	if err != nil {
		panic(err)
	}
	fmt.Printf("jwt: %s\n", ss)
}
