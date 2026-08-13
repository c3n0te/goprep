package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func encryptCaesar(key int, msg string) string {
	alphabet := make(map[rune]int)
	reverse := make(map[int]rune)
	for i := range 26 {
		letter := rune('a' + i)
		alphabet[letter] = i + 1
		reverse[i+1] = letter
	}

	var res strings.Builder
	for _, ch := range strings.ToLower(msg) {
		chRune := rune(ch)
		x := alphabet[chRune]
		encrypted := (x + key) % 26
		chEnc := reverse[encrypted]
		res.WriteRune(chEnc)
	}

	return res.String()
}

func decryptCaesar(key int, msg string) string {
	alphabet := make(map[rune]int)
	reverse := make(map[int]rune)
	for i := range 26 {
		letter := rune('a' + i)
		alphabet[letter] = i + 1
		reverse[i+1] = letter
	}

	var res strings.Builder
	for _, ch := range strings.ToLower(msg) {
		chRune := rune(ch)
		x := alphabet[chRune]
		unencrypted := (x - key) % 26
		chUnenc := reverse[unencrypted]
		res.WriteRune(chUnenc)
	}

	return res.String()
}

func main() {
	fmt.Printf("=== Caesar Cipher Cryptography ===\n[1] Encrypt a message\n[2] Decrypt a message\nChoose an option: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		opt, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("Failed to convert input into integer: %v", err)
		}

		switch opt {
		case 1:
			fmt.Printf("Enter your shift key (e.g., 3): ")
			if scanner.Scan() {
				keyStr := scanner.Text()
				key, err := strconv.Atoi(keyStr)
				if err != nil {
					log.Fatalf("Failed to convert shift key into integer: %v", err)
				}

				fmt.Printf("Enter your message: ")
				if scanner.Scan() {
					msg := scanner.Text()
					res := encryptCaesar(key, msg)
					fmt.Printf("Your encrypted message: %v\n", res)
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatalf("Failed reading key string: %v", err)
			}

			return
		case 2:
			fmt.Printf("Enter your shift key (e.g., 3): ")
			if scanner.Scan() {
				keyStr := scanner.Text()
				key, err := strconv.Atoi(keyStr)
				if err != nil {
					log.Fatalf("Failed to convert shift key into integer: %v", err)
				}

				fmt.Printf("Enter your message: ")
				if scanner.Scan() {
					msg := scanner.Text()
					res := decryptCaesar(key, msg)
					fmt.Printf("Your unencrypted message: %v\n", res)
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatalf("Failed reading key string: %v", err)
			}
			return
		default:
			fmt.Println("Invalid option, please enter '1' or '2'")
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to retrieve user input: %v", err)
	}
}
