package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// your name is (blank), you are (blank) years old, and your username is (blank)

func main() {
	// Create the stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Daily Programmer #1 - Logging User Input")
	fmt.Println("----------------------------------------")

	// Ask the input
	fmt.Print("Your name is? ")
	name, _ := reader.ReadString('\n')

	// Trim the ending \n
	name = strings.TrimSuffix(name, "\n")

	fmt.Print("Your age is? ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSuffix(age, "\n")

	fmt.Print("Your username is? ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSuffix(username, "\n")

	// Create the logging utility
	file, err := os.Create("dp1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	myInfo := "Hello! my name is " + name + " and I'm " + age + " years old. My username is " + username + ". Nice to meet you!"

	file.WriteString(myInfo)

	stream, err := ioutil.ReadFile("dp1.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	fmt.Println(readString)
}
