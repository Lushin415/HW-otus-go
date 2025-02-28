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

	fmt.Printf("Starting server at %s:%s\n", *ip, *port)

	// Запускаем сервер в горутине
	go server.Server(*ip, *port)

	// Даем серверу время на инициализацию
	time.Sleep(time.Second)

	// Запускаем клиент
	client.Klient()

	// Предотвращаем завершение программы сразу после выполнения клиента
	fmt.Println("Press Ctrl+C to exit")
	select {} // Бесконечное ожидание
}
