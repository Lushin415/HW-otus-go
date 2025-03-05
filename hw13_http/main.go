package main

import (
	"flag"
	"fmt"
	"github.com/hw-otus-go/Lushin415/hw13_http/client"
	"github.com/hw-otus-go/Lushin415/hw13_http/server"
	"time"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.String("port", "8080", "Port number")
	flag.Parse()

	go server.Server(*ip, *port)
	fmt.Printf("Сервер запущен %s:%s\n", *ip, *port)
	// Даем серверу время на инициализацию
	time.Sleep(time.Second)

	client.Klient()

	fmt.Println("Нажмите Ctrl+C для выхода")
	select {} // Бесконечное ожидание
}
