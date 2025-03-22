package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

type GenAi struct {
	Context context.Context
	Client  *genai.Client
}

const API_KEY = "AIzaSyCvcMgHI0s95XoS59EyDircaTIoX04kBfI"

func NewGenAi() (*GenAi, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  API_KEY,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return &GenAi{
		Context: ctx,
		Client:  client,
	}, nil
}

func main() {
	genAi, err := NewGenAi()
	if err != nil {
		log.Panicln(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("<-----------    Enter Exit To Quit    ------------->")
	for {
		log.Print("Write Something: ")
		questionText, _ := reader.ReadString('\n')
		if strings.ToLower(questionText) == "exit\n" {
			break
		}

		// Create content input
		input := []*genai.Content{
			{
				Parts: []*genai.Part{{
					Text: questionText,
				}},
			},
		}

		result, err := genAi.Client.Models.GenerateContent(genAi.Context, "gemini-2.0-flash", input, nil)
		if err != nil {
			log.Fatalf("Error sending messages: %v", err)
		}

		log.Println(result.Text())
	}
}
