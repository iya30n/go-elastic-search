package elasticsearch

import (
	"context"
	elastic_contract "elastic-search/internal/contracts/elasticsearch"
	"elastic-search/pkg/elasticsearch"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func Index(indexName string, document elastic_contract.Model) (*esapi.Response, error) {
	esclient, err := elasticsearch.Connect()
	if err != nil {
		return nil, err
	}

	req := esapi.IndexRequest{
		Index:      indexName,                               // Index name
		Body:       strings.NewReader(document.ToJson()), // Document body
		DocumentID: document.GetId(),                             // Document ID
		Refresh:    "true",                                  // Refresh
	}

	res, err := req.Do(context.Background(), esclient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}