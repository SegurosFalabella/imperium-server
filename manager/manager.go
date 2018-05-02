package manager

import (
	"errors"
	"log"

	"github.com/segurosfalabella/imperium-server/dispatcher"

	"github.com/gorilla/websocket"
	"github.com/segurosfalabella/imperium-server/connection"
)

var authToken = "alohomora"
var authTokenResponse = "imperio"

// Manage ...
func Manage(conn connection.WsConn) {
	if err := auth(conn); err != nil {
		log.Fatal(err)
	}
	dispatcher.Dispatch(conn, "hola")
}

func auth(conn connection.WsConn) error {
	_, message, err := conn.ReadMessage()
	if err != nil {
		return errors.New("can't read the message")
	}
	if _, error := validateCredentials(message); error != nil {
		return err
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte(authTokenResponse))
	if err != nil {
		return err
	}
	return nil
}

func validateCredentials(message []byte) (bool, error) {
	if string(message) == authToken {
		return true, nil
	}
	return false, errors.New("Invalid Credentials")
}