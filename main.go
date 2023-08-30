package main

import (
	nasa "github.com/robertsapunarich/nasa"
)

const BaseUrl = "https://api.nasa.gov"
const ApiKey = "KcEfEGqDaIpifxHTLajvMgz4BXJXkFUKNoIxGyU8"

func main() {
	client := nasa.NewApiClient(BaseUrl, ApiKey)
	var result interface{}
	var err error

	result, err = nasa.GetApod(client, nil, nil)
}
