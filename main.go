package main

import (
	"flag"
	"fmt"
	"github.com/ncw/swift"
	"log"
	"os"
	"sync"
)

var co swift.Connection

func main() {

	list := flag.Bool("list", false, "prints a list of existent containers and exits")
	listFilter := flag.String("list_filter", "", "a filter to list the containers")
	prefix := flag.String("prefix", "development_", "prefix for the containers names")
	concurrency := flag.Int("concurrency", 50, "how many concurrent requests")
	num := flag.Int("num", 10000, "number of containers to be create")
	dry := flag.Bool("dry", false, "dry run, won't create any container")
	flag.Parse()

	userName := os.Getenv("SWIFT_API_USER")
	apiKey := os.Getenv("SWIFT_API_KEY")
	authUrl := os.Getenv("SWIFT_AUTH_URL")

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

	err := co.Authenticate()

	if err != nil {
		panic(err)
	}

	if *list {
		log.Printf("Listing containers using filter: %s", *listFilter)
		containersOpts := swift.ContainersOpts{
			Marker: *listFilter,
		}

		containers, _ := co.ContainerNames(&containersOpts)
		fmt.Println(containers)
		return
	}

	log.Printf("Creating containers from %s to %s", containerName(*prefix, 1), containerName(*prefix, *num))
	log.Printf("Using %d concurrent requests", *concurrency)

	var throttle = make(chan int, *concurrency)
	var wg sync.WaitGroup

	for i := 1; i <= *num; i++ {
		// send message to channel. buffered channels will block if it reaches maxConcurrency
		throttle <- 1
		wg.Add(1)
		go handle(containerName(*prefix, i), &wg, throttle)
	}
}

func containerName(p string, n int) string {
	return fmt.Sprintf("%s%05d", p, n)
}

func handle(c string, wg *sync.WaitGroup, throttle chan int) {
	defer wg.Done()
	createContainer(c)
	<-throttle
}

func createContainer(name string) {
	err := co.ContainerCreate(name, nil)
	if err != nil {
		log.Printf("%s failed", name)
	} else {
		log.Println(name)
	}
}
