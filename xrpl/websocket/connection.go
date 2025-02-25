package websocket

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	ErrNotConnected = errors.New("connection is not connected")
)

// Connection is a wrapper around a websocket connection.
// It provides a method to read messages from the connection.
type Connection struct {
	conn *websocket.Conn
	url  string

	mu sync.Mutex
}

// NewConnection creates a new Connection.
func NewConnection(url string) *Connection {
	return &Connection{
		url: url,
	}
}

// Connect opens a websocket connection to the server.
func (c *Connection) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

// Disconnect closes the websocket connection and sets the connection to nil.
// It returns an error if the connection is not connected.
func (c *Connection) Disconnect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.IsConnected() {
		return ErrNotConnected
	}

	if err := c.conn.Close(); err != nil {
		return err
	}
	c.conn = nil
	return nil
}

// IsConnected returns true if the connection is connected.
func (c *Connection) IsConnected() bool {
	return c.conn != nil
}

// ReadMessage reads a message from the connection.
// It returns the message and an error if the message is not read.
// This method is blocking, it will block until a message is red.
func (c *Connection) ReadMessage() ([]byte, error) {
	if !c.IsConnected() {
		return nil, ErrNotConnected
	}
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	return message, nil
}

// WriteMessage writes a message to the connection.
// It returns an error if the message is not written.
func (c *Connection) WriteMessage(message []byte) error {
	if !c.IsConnected() {
		return ErrNotConnected
	}
	return c.conn.WriteMessage(websocket.TextMessage, message)
}
