package rest

import (
	"context"
	"github.com/vmpartner/bitmex/swagger"
	"net/http"
)

func MakeContext(key string, secret string, host string, timeout int64) context.Context {
	return context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
		Key:     key,
		Secret:  secret,
		Host:    host,
		Timeout: timeout,
	})
}

func GetClient(ctx context.Context) *swagger.APIClient {
	c := ctx.Value(swagger.ContextAPIKey).(swagger.APIKey)
	cfg := &swagger.Configuration{
		BasePath:      "https://" + c.Host + "/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		ExpireTime:    5, //seconds
	}

	return swagger.NewAPIClient(cfg)
}

func NewOrder(ctx context.Context, params map[string]interface{}) (swagger.Order, *http.Response, error) {
	client := GetClient(ctx)
	order, response, err := client.OrderApi.OrderNew(ctx, "XBTUSD", params)

	return order, response, err
}

func AmendOrder(ctx context.Context, params map[string]interface{}) (swagger.Order, *http.Response, error) {
	client := GetClient(ctx)
	order, response, err := client.OrderApi.OrderAmend(ctx, params)

	return order, response, err
}

func GetOrder(ctx context.Context, params map[string]interface{}) ([]swagger.Order, *http.Response, error) {
	client := GetClient(ctx)
	orders, response, err := client.OrderApi.OrderGetOrders(ctx, params)

	return orders, response, err
}

func GetPosition(ctx context.Context, params map[string]interface{}) ([]swagger.Position, *http.Response, error) {
	client := GetClient(ctx)
	positions, response, err := client.PositionApi.PositionGet(ctx, params)

	return positions, response, err
}

func GetTrade(ctx context.Context, params map[string]interface{}) ([]swagger.Trade, *http.Response, error) {
	client := GetClient(ctx)
	positions, response, err := client.TradeApi.TradeGet(params)

	return positions, response, err
}

func CancelOrder(ctx context.Context, params map[string]interface{}) ([]swagger.Order, *http.Response, error) {
	client := GetClient(ctx)
	orders, response, err := client.OrderApi.OrderCancel(ctx, params)

	return orders, response, err
}

func GetWallet(ctx context.Context) (swagger.Wallet, *http.Response, error) {
	params := map[string]interface{}{
		"currency": "",
	}
	client := GetClient(ctx)
	wallet, response, err := client.UserApi.UserGetWallet(ctx, params)

	return wallet, response, err
}
