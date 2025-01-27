package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

// Карта для хранения подключенных клиентов
var (
	clients   = make(map[net.Conn]struct{})
	clientsMu sync.Mutex
)

// Обработка подключения клиента
func handleConnection(conn net.Conn) {
	defer conn.Close() // Закрываем соединение при выходе из функции
	clientsMu.Lock()
	clients[conn] = struct{}{} // Добавляем клиента в карту подключенных клиентов
	clientsMu.Unlock()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("Received: %s\n", msg) // Выводим полученное сообщение на сервер

		// Отправка сообщения всем клиентам
		clientsMu.Lock()
		for client := range clients {
			if client != conn {
				fmt.Fprintf(client, "%s\n", msg) // Отправляем сообщение всем клиентам кроме отправителя
			}
		}
		clientsMu.Unlock()
	}

	clientsMu.Lock()
	delete(clients, conn) // Удаляем клиента из карты подключенных клиентов при отключении
	clientsMu.Unlock()
}

func main() {
	ln, err := net.Listen("tcp", ":8080") // Открываем порт 8080 для прослушивания
	if err != nil {
		fmt.Println("Error setting up server:", err)
		os.Exit(1)
	}
	defer ln.Close() // Закрываем порт при выходе из программы

	fmt.Println("Server is running on port 8080")

	for {
		conn, err := ln.Accept() // Ожидаем подключения клиентов
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // Обрабатываем подключение в новой горутине
	}
}
