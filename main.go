package miningpoolhub_api

import (
	"github.com/ging0044/request"
	"strings"
	"log"
)

var url = "https://miningpoolhub.com/index.php"

type Coin struct{
	name string
	code string
	subDomain string
}

func Public(coin Coin) interface{} {
	split := strings.SplitAfter(url, "://")
	split[0] += coin.subDomain
	url = strings.Join(split, "")

	params := map[string]interface{}{
		"page": "api",
		"action": "public",
	}

	data, err := request.Get(url, params)
	if err != nil {
		log.Fatalf("Failed to GET multipoolhub public: %v\nBody: %#v", err, data)
	}

	return data
}
