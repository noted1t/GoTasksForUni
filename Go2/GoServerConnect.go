package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // Подключаемся к серверу на порту 8080
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close() // Закрываем соединение при выходе из программы

	// Чтение сообщений от сервера в отдельной горутине
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println("Message from server:", scanner.Text()) // Выводим полученные сообщения на экран
		}
	}()

	// Отправка сообщений на сервер
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		if scanner.Scan() {
			msg := scanner.Text()
			fmt.Fprintf(conn, "%s\n", msg) // Отправляем введенное сообщение на сервер
		}
	}
}
