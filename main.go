package main

import (
	"bufio"
	e "encryption/encryption"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the golang encryption/decryption CLI tool.")

	option := e.ChooseEncryptOrDecrypt()

	if option == "e" {

		fmt.Println("Please enter the name of the text file you wish to create and encrypt:")
		filename := e.ScanUserInput("")
		if !e.CheckFileExtension(filename) {
			filename = filename + ".txt"
		}

		fmt.Println("Set a password:")
		passphrase := e.ScanUserInput("")

		fmt.Println("Type the data you wish to encrypt:")
		dataName := bufio.NewReader(os.Stdin)
		data, err := dataName.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading data")
		}

		e.EncryptFile(filename, []byte(data), passphrase)
		fmt.Println("File encrypted.")
		fmt.Println("And don't forget your password!")

	} else if option == "d" {

		fmt.Println("Please enter the text file you wish to decrypt:")
		filename := e.ScanUserInput("")

		if !e.CheckFileExtension(filename) {
			filename = filename + ".txt"
		}

		fmt.Println("enter the password:")
		passphrase := e.ScanUserInput("")

		fmt.Println(string(e.DecryptFile(filename, passphrase)))
	}
}
