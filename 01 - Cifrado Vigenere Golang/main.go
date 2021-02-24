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

	var selector = 3
	var key, message string
	for selector > 0 {
		fmt.Print(`
		############################################
		# Descifrador de Vigenere                  #
		############################################
		
		>>> Seleccione una opción para continuar: 
		>>>		1. Encriptar
		>>>		2. Desencriptar
		>>>		3. Salir
		
		Su seleccion:	`)
		fmt.Scan((&selector))
		switch selector {
		case 1:
			fmt.Print("\t\t>>> Introduzca el mensaje que desee encriptar: ")
			fmt.Scan(&message)
			fmt.Print("\t\t>>> Introduzca la clave de cifrado: ")
			fmt.Scan(&key)
			fmt.Print("\t\t----> El mensaje encriptado es " + encrypt(message, key, alphabetStr, alphabet) + " <----")
			fmt.Print("\n\n\t\t>>> ¿Desea continuar?\n\t\t>>> 1. Si\n\t\t>>> 2. No\n\t\tSu seleccion: ")
			fmt.Scan((&selector))
			if selector == 2 {
				selector = -1
			}
		case 2:
			fmt.Print("\t\t>>> Introduzca el mensaje que desee deencriptar: ")
			fmt.Scan(&message)
			fmt.Print("\t\t>>> Introduzca la clave de cifrado: ")
			fmt.Scan(&key)
			fmt.Print("\t\t----> El mensaje deencriptado es " + decrypt(message, key, alphabetStr, alphabet) + " <----")
			fmt.Print("\n\n\t\t>>> ¿Desea continuar?\n\t\t>>> 1. Si\n\t\t>>> 2. No\n\t\tSu seleccion: ")
			fmt.Scan((&selector))
			if selector == 2 {
				selector = -1
			}
		case 3:
			fmt.Print("\t\t>>> Gracias por utilizar este programa. \n\n")
			selector = -1
		default:
			fmt.Print("No ha seleccionado una opción válida")
		}
	}
}
