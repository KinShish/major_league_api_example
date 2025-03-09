package models

// Структура активных клиентов
type Hub struct {
	// Зарегистрированные киленты
	clients map[*Client]bool

	// входящие сообщения от клиентов
	broadcast chan []byte

	// Регистрация заявок от клиента
	register chan *Client

	// отмена регистрации от клиента
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: // Пришла команда на регистрацию
			h.clients[client] = true
		case client := <-h.unregister: // Пришла команда отмены регистрации
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast: // входящее сообщение
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
