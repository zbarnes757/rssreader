package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

// Run performs the search logic
func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()

	// log the error if it failed
	if err != nil {
		log.Fatal(err)
	}

	// create an unbuffered channel to retrieve match results
	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// launch a goroutine to monitor when all the work is done
	go func() {
		// wait for everything to be processed
		waitGroup.Wait()

		close(results)
	}()

	Display(results)
}

// Register is called to register a mathcer for use by the program
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
