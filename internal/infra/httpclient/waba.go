package httpclient

import (
	"fmt"
	"os"
)

func GetWabaHttpClient() *HttpClient {
	accessToken := os.Getenv("META_ACCESS_TOKEN")
	baseUrl := os.Getenv("META_BASE_URL")
	apiVersion := os.Getenv("META_API_VERSION")

	client, err := NewHttpClient(
		WithBaseUrl(fmt.Sprintf("%s/%s", baseUrl, apiVersion)),
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)),
		WithHeader("Content-Type", "application/json"),
	)
	if err != nil {
		panic(err.Error())
	}
	return client
}
