package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateSecretKey(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func main() {
	secretKey := generateSecretKey(32) // Gera uma chave de 32 bytes
	fmt.Println("Sua chave secreta:", secretKey)
}
