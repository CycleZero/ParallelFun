package test

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

func Test_ES(t *testing.T) {

	url := "http://es.server.poyuan233.cn:8088"

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{url},
		Username:  "elastic",
		Password:  "poyuan666",
		//CertificateFingerprint: "",
	})
	if err != nil {
		log.Info(err)
		t.Fatal(err)
	}

	pingResp, err := client.Ping(client.Ping.WithPretty(), client.Ping.WithHuman())
	if err != nil {
		panic(err)
	}
	fmt.Println(pingResp)

}
