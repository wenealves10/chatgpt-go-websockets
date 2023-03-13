package handlers

import (
	"context"
	"io"
	"log"

	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/websocket"
)

type Talks struct {
	conn    *websocket.Conn
	speaker *openai.Client
}

func NewTalks(apiKey string) *Talks {
	return &Talks{
		speaker: openai.NewClient(apiKey),
	}
}

func (t *Talks) Handle(ws *websocket.Conn) {
	t.conn = ws
	t.Read(t.conn)
}

func (t *Talks) Read(conn *websocket.Conn) {
	defer conn.Close()

	msg := make([]byte, 512)

	for {
		n, err := conn.Read(msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		// write
		t.Write(conn, msg[:n])
	}
}

func (t *Talks) Write(conn *websocket.Conn, content []byte) {
	resp, err := t.speaker.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: string(content),
				},
			},
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	if _, err := conn.Write([]byte(resp.Choices[0].Message.Content)); err != nil {
		log.Println(err)
	}

}
