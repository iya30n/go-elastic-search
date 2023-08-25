package elasticsearch

import (
	"context"
	elastic_contract "elastic-search/internal/contracts/elasticsearch"
	"elastic-search/pkg/elasticsearch"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func Delete(indexName string, document elastic_contract.Model) (*esapi.Response, error) {
	esclient, err := elasticsearch.Connect()
	if err != nil {
		return nil, err
	}

	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: document.GetId(),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), esclient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}
