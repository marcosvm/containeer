package container

import (
	"fmt"
	"github.com/ncw/swift"
	"log"
)

func ContainerName(p string, n int) string {
	return fmt.Sprintf("%s%05d", p, n)
}

func PrintContainers(co *swift.Connection, f string) {

	log.Printf("Listing containers using filter: %s", f)

	opts := swift.ContainersOpts{
		Marker: f,
	}

	c, _ := co.ContainerNames(&opts)
	fmt.Println(c)
	return
}

func CreateContainer(co *swift.Connection, name string) {
	if err := co.ContainerCreate(name, nil); err != nil {
		log.Printf("%s failed", name)
	} else {
		log.Println(name)
	}
}
