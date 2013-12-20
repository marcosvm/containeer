// Package container provides methods for
// creating and printing containers
package container

import (
	"fmt"
	"github.com/ncw/swift"
	"log"
)

type Container struct {
	Connection *swift.Connection
}

// ContainerName returns a string to represent a container name
// that will be padded with n zeroes
func (c *Container) ContainerName(p *string, n int) string {
	return fmt.Sprintf("%s%05d", *p, n)
}

// PrintContainers prints a container list with f as a
// filter passed for the CloudFiles API
func (c *Container) PrintContainers(f *string) {

	log.Printf("Listing containers using filter: %s", *f)

	opts := swift.ContainersOpts{
		Marker: *f,
	}

	containers, _ := c.Connection.ContainerNames(&opts) // HL
	fmt.Println(containers)
	return
}

// CreateContainer creates a container with the given name
func (c *Container) CreateContainer(name *string) {
	if err := c.Connection.ContainerCreate(*name, nil); err != nil { // HL
		log.Printf("%s failed", *name)
	} else {
		log.Println(*name)
	}
}
