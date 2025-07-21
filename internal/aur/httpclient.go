package aur

import (
	"net/http"
	"time"

	"github.com/joaogiacometti/aurgo/internal/config"
)

var client *http.Client

func getClient() *http.Client {
	if client == nil {
		client = &http.Client{
			Timeout: time.Duration(config.HTTPTimeout) * time.Second,
		}
	}

	return client
}
