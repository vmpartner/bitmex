# Bitmex API
Packages for work with bitmex rest and websocket API on golang.  
Target of this packages make easy access to bitmex API including testnet platform.  

Packages covered tests.  

In rest package implemented basic must have methods, you can easy add needed method by extending rest package. Autocomplete working based on swagger bitmex package. Please contribute if you will add new methods.


## Usage
Please see full example in main.go

####  REST
```
// Load config
cfg := config.LoadConfig("config.json")
ctx := rest.MakeContext(cfg.Key, cfg.Secret, cfg.Host)

// Get wallet
w, response, err := rest.GetWallet(ctx)
tools.CheckErr(err)
fmt.Printf("Status: %v, wallet amount: %v\n", response.StatusCode, w.Amount)
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

        // Your logic here
        fmt.Printf("%+v\n", res)
    }
}()

```

## Example
Example of usage look in main.go

## More
Those who will donated more $50 i will send my working private code bot based on neural analyze.  
I spent a lot of time implementing this packages and will be glad of any support. Thank you!
```
eth: 0x3e9b92625c49Bfd41CCa371D1e4A1f0d4c25B6fC
btc: 35XDoFSA8QeM26EnCyhQPTMBZm4S1DvncE
```
vmpartner[a]gmail.com


