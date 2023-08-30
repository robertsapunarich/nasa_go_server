package main

import (
	"fmt"

	nasa "github.com/robertsapunarich/nasa"
)

const BaseUrl = "https://api.nasa.gov"
const ApiKey = "KcEfEGqDaIpifxHTLajvMgz4BXJXkFUKNoIxGyU8"

func main() {
	client := nasa.NewApiClient(BaseUrl, ApiKey)

	result, err := client.GetApod(nil)

	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("URL: %s", result.Url)

	println(str)
}
