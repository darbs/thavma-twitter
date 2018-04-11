package fetch

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/darbs/thavma-twitter/internal/entity"
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

/*
Get tweets done by search term

dream of premium api access
 */
func Get(symbol string) {
	fmt.Printf("Fetching Tweets for %s\n", symbol)

	search, _, err := twClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: symbol, // todo validate symbol
		// todo restrict by last id
	})

	if err != nil {
		log.Fatal(err)
	}

	i := 0
	l := len(search.Statuses)
	for i < l {
		//fmt.Printf("Id: %v U: %v R: %v T: %v\n", search.Statuses[i].ID, search.Statuses[i].User.ScreenName, search.Statuses[i].RetweetCount, search.Statuses[i].Text)
		//fmt.Printf("SEARCH TWEETS:\n%+v\n", search.Statuses[i].ID, search.Statuses[i].Text, search.Statuses[i].RetweetedStatus)
		t, _ := search.Statuses[i].CreatedAtTime()
		tweet := entity.Tweet{
			EntityId: search.Statuses[i].ID,
			Symbol:   strings.Trim(symbol, "$"),
			Date:     t.UTC(),
			Weight:   search.Statuses[i].RetweetCount,
			Creator:  search.Statuses[i].User.ScreenName,
			Content:  search.Statuses[i].Text,
		}

		tweet.Save()
		i++
	}
	fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)
}
