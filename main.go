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

		fmt.Println("Please enter the file name you wish to encrypt:")
		filename := e.ScanUserInput("")

		fmt.Println("Set a password:")
		passphrase := e.ScanUserInput("")

		fmt.Println("Please enter the data you wish to encrypt:")
		dataName := bufio.NewReader(os.Stdin)
		data, _ := dataName.ReadString('\n')

		e.EncryptFile(filename, []byte(data), passphrase)
		fmt.Println("File encrypted!")
	} else if option == "d" {

		fmt.Println("Please enter the file name you wish to decrypt:")
		filename := e.ScanUserInput("")

		fmt.Println("enter the password:")
		passphrase := e.ScanUserInput("")

		fmt.Println("Decryption complete!")
		fmt.Println(string(e.DecryptFile(filename, passphrase)))
	}
}
