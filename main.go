package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"

	"github.com/marcosvm/containeer/container"
	"github.com/ncw/swift"
)

var (
	cc             container.Container
	list           = flag.Bool("list", false, "prints a list of existent containers and exits")
	listFilter     = flag.String("list_filter", "", "a filter to list the containers")
	prefix         = flag.String("prefix", "development_", "prefix for the containers names")
	concurrency    = flag.Int("concurrency", 50, "how many concurrent requests")
	num            = flag.Int("num", 10000, "number of containers to create")
	dry            = flag.Bool("dry", false, "dry run, won't create any container")
	single         = flag.String("single", "", "create a single container")
	cpuprofile     = flag.String("cpuprofile", "", "write cpu profile to file")
	objects        = flag.String("objects", "", "list objects given a container")
	objects_marker = flag.String("objects_marker", "", "marker to list objects given a container")
	objects_limit  = flag.Int("objects_limit", 10000, "limit of objects objects to list given a container")
	getbytes       = flag.String("getbytes", "", "get bytes from object")
	directory      = flag.String("directory", "", "container to get bytes from")
)

func main() {

	flag.Parse()

	userName := os.Getenv("SWIFT_API_USER")
	apiKey := os.Getenv("SWIFT_API_KEY")
	authUrl := os.Getenv("SWIFT_AUTH_URL")

	if userName == "" || apiKey == "" || authUrl == "" {
		log.Fatal("SWIFT_API_USER, SWIFT_API_KEY and SWIFT_AUTH_URL environment variables need to be set")
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	log.Println("Starting")
	log.Println("Creating connection to CloudFiles")

	co := swift.Connection{
		UserName: userName,
		ApiKey:   apiKey,
		AuthUrl:  authUrl,
	}

	if err := co.Authenticate(); err != nil {
		log.Fatal(err)
	}

	cc = container.Container{
		Connection: &co,
	}

	if *list {
		cc.PrintContainers(listFilter)
		os.Exit(0)
	}

	if *single != "" {
		cc.CreateContainer(single)
		os.Exit(0)
	}

	if *objects != "" {
		cc.ListObjects(*objects, *objects_marker, *objects_limit)
		os.Exit(0)
	}

	if *getbytes != "" {
		if *directory != "" {
			cc.GetBytes(*directory, *getbytes)
			os.Exit(0)
		}
	}

	log.Printf("Creating containers from %s to %s", cc.ContainerName(prefix, 1), cc.ContainerName(prefix, *num))
	log.Printf("Using %d concurrent requests", *concurrency)

	if *dry {
		return
	}

	var throttle = make(chan int, *concurrency) // HL
	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 1; i <= *num; i++ {
		// send message to channel. buffered channels will block if it reaches maxConcurrency
		throttle <- 1 // HL
		wg.Add(1)
		go handle(cc.ContainerName(prefix, i), &wg, throttle) // HL
	}
}

func handle(c string, wg *sync.WaitGroup, throttle chan int) {
	defer wg.Done()
	cc.CreateContainer(&c)
	<-throttle // HL
}
