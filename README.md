# Bitmex API
Packages for work with bitmex rest and websocket API on golang.  
Target of this packages make easy access to bitmex API including testnet platform.


## Usage
Please see full example in main.go

####  REST
```
// Load config
cfg := config.LoadConfig("config.json")
ctx := rest.MakeContext(cfg.Key, cfg.Secret, cfg.Host)

// Get wallet
wallet, response, err := rest.GetWallet(ctx)
tools.CheckErr(err)
fmt.Printf("Status: %v, wallet amount: %v\n", response.StatusCode, wallet.Amount)

// Place order
params := map[string]interface{}{
    "side":     "Buy",
    "symbol":   "XBTUSD",
    "ordType":  "Limit",
    "orderQty": 1,
    "price":    9000,
    "clOrdID":  "MyUniqID_123",
    "execInst": "ParticipateDoNotInitiate",
}
order, response, err := rest.NewOrder(ctx, params)
tools.CheckErr(err)
fmt.Printf("Order: %+v, Response: %+v\n", order, response)
```

#### Websocket
```
// Load config
cfg := config.LoadConfig("config.json")

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

// Listen
go func() {
    for {
        message := <-chReadFromWS
        res, err := bitmex.DecodeMessage(message)
        tools.CheckErr(err)

        // Business logic
        switch res.Table {
        case "orderBookL2":
            if res.Action == "partial" {
                // Update table
            } else {
                // Update row
            }
        case "order":
            if res.Action == "partial" {
                // Update table
            } else {
                // Update row
            }
        case "position":
            if res.Action == "partial" {
                // Update table
            } else {
                // Update row
            }
        }
    }
}()

```

## Example
Example of usage look in main.go

## More
I will be glad of any support. Thank you!
```
eth: 0x3e9b92625c49Bfd41CCa371D1e4A1f0d4c25B6fC
btc: 35XDoFSA8QeM26EnCyhQPTMBZm4S1DvncE
```
