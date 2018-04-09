package websocket

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"time"
	"github.com/gorilla/websocket"
	"github.com/vmpartner/bitmex/tools"
)

type Message struct {
	Op   string        `json:"op,omitempty"`
	Args []interface{} `json:"args,omitempty"`
}

func (m *Message) AddArgument(argument string) {
	m.Args = append(m.Args, argument)
}

func Connect(host string) *websocket.Conn {
	u := url.URL{Scheme: "wss", Host: host, Path: "/realtime"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	tools.CheckErr(err)

	return conn
}

func ReadFromWSToChannel(c *websocket.Conn, chRead chan<- []byte) {
	for {
		_, message, err := c.ReadMessage()
		//fmt.Println("Read", string(message))
		tools.CheckErr(err)
		chRead <- message
	}
}

func WriteFromChannelToWS(c *websocket.Conn, chWrite <-chan interface{}) {
	for {
		message := <-chWrite
		if reflect.TypeOf(message).String() == "websocket.Message" {
			var err error
			message, err = json.Marshal(message)
			tools.CheckErr(err)
		}
		err := c.WriteMessage(websocket.TextMessage, message.([]byte))
		tools.CheckErr(err)
	}
}

func GetAuthMessage(key string, secret string) Message {
	nonce := time.Now().Unix() + 412
	req := fmt.Sprintf("GET/realtime%d", nonce)
	sig := hmac.New(sha256.New, []byte(secret))
	sig.Write([]byte(req))
	signature := hex.EncodeToString(sig.Sum(nil))
	var msgKey []interface{}
	msgKey = append(msgKey, key)
	msgKey = append(msgKey, nonce)
	msgKey = append(msgKey, signature)

	return Message{"authKey", msgKey}
}
