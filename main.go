package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/akhilsharma90/file-encrypt/filecrypt"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	// If not enough args, return help text
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run 'GoCryptor encrypt' to encrypt a file and 'GoCryptor decrypt' to decrypt a file.")
		os.Exit(1)
	}

}

func printHelp() {
	fmt.Println("GoCryptor")
	fmt.Println("Simple file encrypter for your daily needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tGoCryptor encrypt <Insert File Path>")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file with a given password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help section for the script")
	fmt.Println("")
}

func encryptHandle() {

	if len(os.Args) < 3 {
		println("Missing File Path. For more information run GoCryptor help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\nFile is now protected.")

}

func getPassword() []byte {
	fmt.Print("Enter Password: ")
	password, _ := terminal.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	password2, _ := terminal.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords do not match. Please try again.\n")
		return getPassword()
	}
	return password
}

func decryptHandle() {

	if len(os.Args) < 3 {
		println("Missing file path. For more information run GoCryptor help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	fmt.Print("Enter Password: ")
	password, _ := terminal.ReadPassword(0)

	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\nFile successfully decrypted.")

}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}

	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}