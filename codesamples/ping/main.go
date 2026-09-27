package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AhaSend/ahasend-go/api"
)

func main() {
	client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))

	response, _, err := client.UtilityAPI.Ping(context.Background())
	if err != nil {
		log.Fatalf("Error pinging API: %v", err)
	}
	fmt.Println(response.Message)
}
