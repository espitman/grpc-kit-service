package elk

import (
	"context"
	"grpc-kit-service/provider/config"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func init() {
	var cl, err = elastic.NewSimpleClient(elastic.SetURL(config.GetString("elastic.url")))
	if err != nil {
		panic(err)
	}
	client = cl
}

func Connect() *elastic.Client {
	return client
}

func CreateIndex(name string) {
	_, _ = client.CreateIndex(name).Do(context.Background())
}

func AddData(index string, docType string, body interface{}) *elastic.IndexResponse {
	result, err := client.Index().
		Index(index).
		Type(docType).
		BodyJson(body).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return result
}

func UpsertData(index string, docType string, id string, body interface{}) *elastic.IndexResponse {
	result, err := client.Index().
		Index(index).
		Type(docType).
		Id(id).
		BodyJson(body).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return result
}

func Find(index string, termQuery *elastic.TermQuery) *elastic.SearchResult {
	searchResult, err := client.Search().
		Index(index).     // search in index "twitter"
		Query(termQuery). // specify the query
		From(0).Size(1).  // take documents 0-9
		Pretty(true).     // pretty print request and response JSON
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return searchResult
}

func FindOne(index string, termQuery *elastic.TermQuery) *elastic.SearchHit {
	searchResult, err := client.Search().
		Index(index).     // search in index "twitter"
		Query(termQuery). // specify the query
		From(0).Size(1).  // take documents 0-9
		Pretty(true).     // pretty print request and response JSON
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	if searchResult.Hits.TotalHits.Value > 0 {
		for _, hit := range searchResult.Hits.Hits {
			return hit
		}
	}
	return nil
}
