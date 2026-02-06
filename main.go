package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-email-plugin] Starting Email plugin...")
	plugin := sdk.Init("hc-email-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("EmailStatus", "Email plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getEmailStatus",
		sdk.ComplexObjectField("Get email plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-email-plugin] Plugin ready")
	plugin.Serve()
}
