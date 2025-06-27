package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)


func main() {
	// get user input (name)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your name: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Uh oh - error reading name:", err)
			continue
		}

		name = strings.TrimSpace(name)

		fmt.Println("Nice you meet you,", name, "don't mind me, just mowing the lawn...")
		break
	}

	time.Sleep(2 * time.Second)  

	// make cool ascii lag effect with lawnmower farmer 
	text := asciiArt["farmer"]
	typewriterSpeed := 2 * time.Millisecond

	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(typewriterSpeed)
	}
	fmt.Println()
	
	time.Sleep(2 * time.Second)  

	fmt.Println("Oh! While you're here, I just planted some new vegetables on my farm - do you want to see?")
	
	prompt := promptui.Select{
		Label: "Explore my farm? ",
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Yes" {
		fmt.Println(asciiArt["welcome-border"])
	} else {
		fmt.Println("If no, say... at the very least, you should start an internet farm. I've got www.corn.com seeds for you. Plant them on a nice plot of virtual land for me. I am on a journey to cultivate the INTERWEBS!")
		time.Sleep(2 * time.Second)  

	}
}
