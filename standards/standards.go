package standards

import (
	"crypto/aes"
	"fmt"
	"strings"
)

func StandardLibrary() {
	greetings := "Hello, There Friends"
	fmt.Println(greetings)

	keyBlockSize := aes.BlockSize
	fmt.Printf("AES Block Size -> %v\n", keyBlockSize)

	if keyBlockSize == 16 {
		fmt.Println("AES Block size is 16")
	} else {
		fmt.Println("No block size define")
	}

	fmt.Printf("Block Size -> %T, Block Size Values -> %v\n", keyBlockSize, keyBlockSize)

	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Yo"))

	if strings.Contains(greetings, "Hello") {
		fmt.Println("Yup welcome")
	} else {
		fmt.Println("Nope message is not correct")
	}

	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "th"))

	first := "Hello world"
	second := "Hello world"

	fmt.Println(strings.Compare(first, second))

	fmt.Println(strings.Split(greetings, " "))

}
