package main

import (
	"fmt"
)

func generateAlphabet() (string, map[uint8]int) {
	alphabet := make(map[uint8]int)
	var alphabetStr = "abcdefghijklmnopkrstuvwxyz"
	for i := 0; i < len(alphabetStr); i++ {
		alphabet[alphabetStr[i]] = i
	}
	return alphabetStr, alphabet
}

func module(x int, y int) int {
	if x < 0 {
		return x + y
	}
	return x
}
func encrypt(message string, key string, alphabetStr string, alphabet map[uint8]int) string {
	var encryptedMessage = ""
	var nextLetter, currentLetter, keyPos uint8
	var KeyIterator int = 0
	for i := 0; i < len(message); i++ {
		currentLetter = message[i]
		keyPos = key[KeyIterator]
		nextLetter = alphabetStr[(alphabet[currentLetter]+alphabet[keyPos])%len(alphabetStr)]
		encryptedMessage = fmt.Sprint(encryptedMessage, string(nextLetter))
		KeyIterator = (KeyIterator + 1) % len(key)
	}
	return encryptedMessage
}

func decrypt(message string, key string, alphabetStr string, alphabet map[uint8]int) string {
	var decryptedMessage = ""
	var nextLetter, currentLetter, keyPos uint8
	var KeyIterator int = 0
	for i := 0; i < len(message); i++ {
		currentLetter = message[i]
		keyPos = key[KeyIterator]
		var pos = module((alphabet[currentLetter] - alphabet[keyPos]), len(alphabetStr))
		nextLetter = alphabetStr[pos]
		decryptedMessage = fmt.Sprint(decryptedMessage, string(nextLetter))
		KeyIterator = (KeyIterator + 1) % len(key)
	}
	return decryptedMessage
}

func main() {
	alphabetStr, alphabet := generateAlphabet()

	var selector = 2
	var key, message string

	switch selector {
	case 1:
		fmt.Println("Introduzca el mensaje que desee encriptar: ")
		fmt.Scan(&message)
		fmt.Println("Introduzca la clave de cifrado: ")
		fmt.Scan(&key)
		fmt.Println("El mensaje encriptado es " + encrypt(message, key, alphabetStr, alphabet))
	case 2:
		fmt.Println("Introduzca el mensaje que desee deencriptar: ")
		fmt.Scan(&message)
		fmt.Println("Introduzca la clave de cifrado: ")
		fmt.Scan(&key)
		fmt.Println("El mensaje deencriptado es " + decrypt(message, key, alphabetStr, alphabet))
	}
}
