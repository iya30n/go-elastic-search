package elasticsearch

import (
	"crypto/tls"
	"net/http"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

var doOnce sync.Once
var singletonConnection *elasticsearch.Client
var connectionError error

func Connect() (*elasticsearch.Client, error) {
	doOnce.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200",
				"http://localhost:9300",
			},
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			// NOTE: we should env config for username, password
			Username: "iya30n",
			Password: "11111111",
		}

		singletonConnection, connectionError = elasticsearch.NewClient(cfg)
	})

	return singletonConnection, connectionError
}
