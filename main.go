package thavma_twitter

import (
	"fmt"
	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"

	//"golang.org/x/oauth2"
	"github.com/dghubble/oauth1"
	"github.com/darbs/thavma-twitter/internal/fetch"

	"os"
	//"os/signal"
	//"syscall"
	"flag"
	"log"
)

func main() {



	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient)
	//demux := twitter.NewSwitchDemux()
	//demux.Tweet = func(tweet *twitter.Tweet) {
	//	fmt.Println(tweet.Text)
	//}
	//demux.DM = func(dm *twitter.DirectMessage) {
	//	fmt.Println(dm.SenderID)
	//}
	//demux.Event = func(event *twitter.Event) {
	//	fmt.Printf("%#v\n", event)
	//}
	//go demux.HandleChan(Stream.Messages)

	fmt.Println("Starting Stream...")

	//filterParams := &twitter.StreamFilterParams{
	//	Track:         []string{"cat"},
	//	StallWarnings: twitter.Bool(true),
	//}
	//stream, err := client.Streams.Filter(filterParams)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//filterParams := &twitter.StreamFilterParams{
	//	Track:         []string{"AMD"},
	//	StallWarnings: twitter.Bool(true),
	//}
	//stream, err := client.Streams.Filter(filterParams)

	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "$atrs",
	})

	//sampleParams := &twitter.StreamSampleParams{
	//	StallWarnings: twitter.Bool(true),
	//}
	//stream, err := client.Streams.Sample(sampleParams)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	len := len(search.Statuses)
	for i < len {
		fmt.Printf("Id: %v U: %v R: %v T: %v\n", search.Statuses[i].ID, search.Statuses[i].User.Name, search.Statuses[i].Retweeted, search.Statuses[i].Text)
		//fmt.Printf("SEARCH TWEETS:\n%+v\n", search.Statuses[i].ID, search.Statuses[i].Text, search.Statuses[i].RetweetedStatus)

		i++
	}
	fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	//ch := make(chan os.Signal)
	//signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	//log.Println(<-ch)

	//stream.Stop()
	fmt.Printf("hello, world\n")
}
