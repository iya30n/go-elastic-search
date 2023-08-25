package elasticsearch

import (
	"context"
	"elastic-search/pkg/elasticsearch"
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func Search(indexName string, searchQuery map[string]string) (map[string]interface{}, error) {
	esclient, err := elasticsearch.Connect()
	if err != nil {
		return nil, err
	}

	equery := elasticQuery{"query": {"match": searchQuery}}

	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(equery.toJson()),
	}

	res, err := req.Do(context.Background(), esclient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r, nil
}

type elasticQuery map[string]map[string]map[string]string

func (eq elasticQuery) toJson() string {
	b, _ := json.Marshal(eq)

	return string(b)
}
