package fetch

import (
	"fmt"
	"log"
	"os"

	//"github.com/darbs/thavma-twitter/internal/entity"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
	twitterConsumerKey    = os.Getenv("TWITTER_CONSUMER_KEY")
	twitterConsumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	twitterAccessToken    = os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessSecret   = os.Getenv("TWITTER_ACCESS_SECRET")
	twClient              *twitter.Client
)

func init() {
	if twitterConsumerKey == "" || twitterConsumerSecret == "" || twitterAccessToken == "" || twitterAccessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(twitterConsumerKey, twitterConsumerSecret)
	token := oauth1.NewToken(twitterAccessToken, twitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twClient = twitter.NewClient(httpClient)
}

func Get(symbol string) {
	fmt.Printf("Fetching Tweets for %s\n", symbol)

	search, _, err := twClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: symbol, // todo validate symbol
	})

	if err != nil {
		log.Fatal(err)
	}

	i := 0
	l := len(search.Statuses)
	for i < l {
		fmt.Printf("Id: %v U: %v R: %v T: %v\n", search.Statuses[i].ID, search.Statuses[i].User.Name, search.Statuses[i].Retweeted, search.Statuses[i].Text)
		//fmt.Printf("SEARCH TWEETS:\n%+v\n", search.Statuses[i].ID, search.Statuses[i].Text, search.Statuses[i].RetweetedStatus)

		i++
	}
	fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)
}
