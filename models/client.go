package models

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Врамя ожидания сообщения от клиента
	writeWait = 10 * time.Second

	// максимальное время ответа клиента(пинг понг), если до этого таймера не приходили сообщения происходит разрыв
	pongWait = 60 * time.Second

	// отправка сообщения для поддержания соединения с клиентом
	pingPeriod = (pongWait * 9) / 10

	// Максимальный размер сообщения
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Структура где хранятся все данные клиента подключенного к чату
type Client struct {
	hub *Hub

	// Соединение через вебсокеты
	conn *websocket.Conn

	// канал исходящих данных
	send chan []byte
}

// readPump чтение сообщения от клиента
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for { // Бесконечный цикл, сообщения могут прийти в любой момент, постоянно должны слушать
		// чтение сообщения
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// Передача сообщения на сервер
		c.hub.broadcast <- message
	}
}

// writePump отправка сообщений клиенту
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for { // Бесконечный цикл, если получаем сообщение от клиента записываем его в хаб.
		//  После чего считываем его и отдаем всем остальным клентам
		select {
		case message, ok := <-c.send: // Проверяем есть ли сообщение в памяти у клиента
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// если потерено соединение выдаем сообщение о закрытии соединения
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Отправляем все сообщения из очереди, если такие образовались
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C: // Постоянная отправка сигнала ПИНГ до клиента. Проверка клиента
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// запуск сервера, который будет обрабатывать запросы клиентов
// выполняется когда клиент первый раз подключается к серверу
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client // регистрация клиента

	// Асинхронный запуск методов
	go client.writePump()
	go client.readPump()
}
