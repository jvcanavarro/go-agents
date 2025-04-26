package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jvcanavarro/go-agents/brain"
)

func main() {
	agent, err := brain.NewAgent(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	agent.Debug = true

	fmt.Println("Assistente: Seja bem vindo meu Mestre, tudo bem?")
	
	ctx := context.Background()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Eu: ")
		input, _ := reader.ReadString('\n')

		resp, err := agent.SendMessage(ctx, input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Assistente: %s\n", resp.Text())
	}
}