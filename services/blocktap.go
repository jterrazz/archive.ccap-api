package services

import (
	"context"
	"fmt"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

const BASE_URL = "https://api.blocktap.io/server"

var src = oauth2.StaticTokenSource(
	&oauth2.Token{
		AccessToken: "4ca7c933b5d50a6f0dcb5b9070e85b4373592ad589d118f2290137d300fffb52",
		TokenType: "Bearer",
	},
)
var httpClient = oauth2.NewClient(context.Background(), src)
var client = graphql.NewClient(BASE_URL, httpClient)





type BlocktapAsset struct {
	Name string `graphql:"assetName"`
}

func GetAssets() []*BlocktapAsset {
	var query struct {
		Assets []*BlocktapAsset `graphql:"assets(sort: [{ marketCapRank: ASC }]"`
	}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Println(err)
	}
	return query.Assets
}
