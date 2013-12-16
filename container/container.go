package container

import (
	"fmt"
	"github.com/ncw/swift"
	"log"
)

// Create a string represent a container name
// n will be padded with zeroes
func ContainerName(p string, n int) string {
	return fmt.Sprintf("%s%05d", p, n)
}

// START1 OMIT
func PrintContainers(co *swift.Connection, f string) {

	log.Printf("Listing containers using filter: %s", f)

	opts := swift.ContainersOpts{
		Marker: f,
	}

	c, _ := co.ContainerNames(&opts) // HL
	fmt.Println(c)
	return
}

// STOP1 OMIT

// START2 OMIT
func CreateContainer(co *swift.Connection, name string) {
	if err := co.ContainerCreate(name, nil); err != nil { // HL
		log.Printf("%s failed", name)
	} else {
		log.Println(name)
	}
}

// STOP2 OMIT
