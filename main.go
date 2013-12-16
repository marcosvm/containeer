package main

import (
	"flag"
	"github.com/marcosvm/containeer/container"
	"github.com/ncw/swift"
	"log"
	"os"
	"sync"
)

var co swift.Connection

func main() {

	// START3 OMIT
	list := flag.Bool("list", false, "prints a list of existent containers and exits")

	listFilter := flag.String("list_filter", "", "a filter to list the containers")

	prefix := flag.String("prefix", "development_", "prefix for the containers names") // HL

	concurrency := flag.Int("concurrency", 50, "how many concurrent requests")

	num := flag.Int("num", 10000, "number of containers to be create")

	dry := flag.Bool("dry", false, "dry run, won't create any container")

	flag.Parse()
	// STOP3 OMIT

	// START4 OMIT
	userName := os.Getenv("SWIFT_API_USER")

	apiKey := os.Getenv("SWIFT_API_KEY") // HL

	authUrl := os.Getenv("SWIFT_AUTH_URL")
	// STOP4 OMIT

	if userName == "" || apiKey == "" || authUrl == "" {
		log.Fatal("SWIFT_API_USER, SWIFT_API_KEY and SWIFT_AUTH_URL environment variables need to be set")
	}

	if *dry {
		return
	}

	log.Println("Starting")
	log.Println("Creating connection to CloudFiles")

	co = swift.Connection{
		UserName: userName,
		ApiKey:   apiKey,
		AuthUrl:  authUrl,
	}

	if err := co.Authenticate(); err != nil {
		log.Fatal(err)
	}

	if *list {
		container.PrintContainers(&co, *listFilter)
		os.Exit(0)
	}

	log.Printf("Creating containers from %s to %s", container.ContainerName(*prefix, 1), container.ContainerName(*prefix, *num))
	log.Printf("Using %d concurrent requests", *concurrency)

	// START1 OMIT
	var throttle = make(chan int, *concurrency) // HL
	var wg sync.WaitGroup

	for i := 1; i <= *num; i++ {
		// send message to channel. buffered channels will block if it reaches maxConcurrency
		throttle <- 1 // HL
		wg.Add(1)
		go handle(container.ContainerName(*prefix, i), &wg, throttle) // HL
	}
	// STOP1 OMIT
}

// START2 OMIT
func handle(c string, wg *sync.WaitGroup, throttle chan int) {
	defer wg.Done()
	container.CreateContainer(&co, c)
	<-throttle // HL
}

// STOP2 OMIT
