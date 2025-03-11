package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

func readWs(conn *websocket.Conn, ch chan string) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("ERROR: Could not read message:", err)
			return
		}
		ch <- string(msg)
	}
}

func shutdownHandler(shutdown chan os.Signal, conn *websocket.Conn) {
	<-shutdown
	log.Println("Received shutdown signal. Shutting down gracefully...")
	conn.Close()
	os.Exit(0)
}

func main() {
	url := flag.String("url", "wss://websocket-echo.com", "WebSocket URL")
	noColor := flag.Bool("no-color", false, "Disable colored output")
	flag.Parse()

	inPrompt := ">> "
	outPrompt := "<<"

	if !(*noColor) {
		inPrompt = color.BlueString(inPrompt)
		outPrompt = color.GreenString(outPrompt)
	}

	conn, _, err := websocket.DefaultDialer.Dial(*url, nil)
	if err != nil {
		log.Fatal("ERROR: Could not connect to WebSocket:", err)
	}
	defer conn.Close()

	// Handling shutdown signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go shutdownHandler(shutdown, conn)

	messageCh := make(chan string)
	go readWs(conn, messageCh)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(inPrompt)
		if scanner.Scan() {
			text := scanner.Text()
			err := conn.WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				log.Println("ERROR: Could not send message:", err)
				return
			}
		}
		msg := <-messageCh
		fmt.Println(outPrompt, msg)
	}
}
