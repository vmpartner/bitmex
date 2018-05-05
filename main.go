package main

import (
	"fmt"
	"github.com/vmpartner/bitmex/bitmex"
	"github.com/vmpartner/bitmex/config"
	"github.com/vmpartner/bitmex/rest"
	"github.com/vmpartner/bitmex/tools"
	"github.com/vmpartner/bitmex/websocket"
	"strings"
)

// Usage example
func main() {

	// Load config
	cfg := config.LoadConfig("config.json")
	ctx := rest.MakeContext(cfg.Key, cfg.Secret, cfg.Host, cfg.Timeout)

	// Get wallet
	w, response, err := rest.GetWallet(ctx)
	tools.CheckErr(err)
	fmt.Printf("Status: %v, wallet amount: %v\n", response.StatusCode, w.Amount)

	// Connect to WS
	conn := websocket.Connect(cfg.Host)
	defer conn.Close()

	// Listen read WS
	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(conn, chReadFromWS)

	// Listen write WS
	chWriteToWS := make(chan interface{}, 100)
	go websocket.WriteFromChannelToWS(conn, chWriteToWS)

	// Authorize
	chWriteToWS <- websocket.GetAuthMessage(cfg.Key, cfg.Secret)

	// Read first response message
	message := <-chReadFromWS
	if !strings.Contains(string(message), "Welcome to the BitMEX") {
		fmt.Println(string(message))
		panic("No welcome message")
	}

	// Read auth response success
	message = <-chReadFromWS
	res, err := bitmex.DecodeMessage(message)
	tools.CheckErr(err)
	if res.Success != true || res.Request.(map[string]interface{})["op"] != "authKey" {
		panic("No auth response success")
	}

	// Listen websocket before subscribe
	go func() {
		for {
			message := <-chReadFromWS
			res, err := bitmex.DecodeMessage(message)
			tools.CheckErr(err)

			// Your logic here
			fmt.Printf("%+v\n", res)
		}
	}()

	// Subscribe
	messageWS := websocket.Message{Op: "subscribe"}
	messageWS.AddArgument("orderBookL2:XBTUSD")
	messageWS.AddArgument("order")
	messageWS.AddArgument("position")
	chWriteToWS <- messageWS

	// Loop forever
	select {}
}
