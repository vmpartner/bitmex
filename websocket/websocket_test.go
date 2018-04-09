package websocket

import (
	"testing"
	"time"
	"fmt"
	"strings"
)

func TestConnectMaster(t *testing.T) {
	conn := Connect("www.bitmex.com")
	if conn == nil {
		t.Error("No connect to ws")
	}
}

func TestConnectDev(t *testing.T) {
	conn := Connect("testnet.bitmex.com")
	if conn == nil {
		t.Error("No connect to testnet ws")
	}
}

func TestConnectFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Connect("")
}

func TestWorkerReadMessages(t *testing.T) {
	chReaderMessage := make(chan []byte)
	conn := Connect("testnet.bitmex.com")
	go ReadFromWSToChannel(conn, chReaderMessage)
	message := <-chReaderMessage
	if message == nil {
		t.Error("Empty message")
	}
	close(chReaderMessage)
}

func TestWorkerWriteMessages(t *testing.T) {

	conn := Connect("testnet.bitmex.com")

	// Read
	chReadFromWS := make(chan []byte, 10)
	go ReadFromWSToChannel(conn, chReadFromWS)

	// Write
	chWriteToWS := make(chan interface{}, 10)
	go WriteFromChannelToWS(conn, chWriteToWS)

	// Send ping
	chWriteToWS <- []byte(`ping`)

	// Read first response message
	message := <-chReadFromWS
	if !strings.Contains(string(message), "Welcome to the BitMEX") {
		fmt.Println(string(message))
		t.Error("No welcome message")
	}

	// Read second response message
	message = <-chReadFromWS
	if !strings.Contains(string(message), "pong") {
		fmt.Println(string(message))
		t.Error("No pong message")
	}

	time.Sleep(1 * time.Second)
}
