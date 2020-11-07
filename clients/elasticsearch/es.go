package elasticsearch

import "github.com/olivere/elastic/v7"

var Client esClientInterface = &esClient{}

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient{
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckIntervalTime(10*time.Second)
	}
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (cln *esClient) setClient(c *elastic.Client) {
	c.client = client
}

func (cln *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	res, err := cln.client.Index().
		Index(index).
		BodyJson(doc),
		Do(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}
