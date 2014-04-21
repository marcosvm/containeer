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
// listing object that names are greater than filter
// obs: limit will be the API default and
// endmarker is not used
func (c *Container) PrintContainers(f *string) {

	log.Printf("Listing containers using filter: %s", *f)

	opts := swift.ContainersOpts{
		Marker: *f,
	}

	containers, _ := c.Connection.Containers(&opts) // HL
	for _, c := range containers {
		log.Printf("Name:%s, objects:%d , bytes:%d", c.Name, c.Count, c.Bytes)
	}
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

// ListObjects prints a list of objects given a container name
// and a limit to start
func (c *Container) ListObjects(container, marker string, limit int) {
	log.Printf("Listing objects for container: %s", container)
	log.Printf("listing %d objects", limit)

	if marker != "" {
		log.Printf("starting at: %s", marker)
	}

	// empty for now, let's see how it behaves
	opts := swift.ObjectsOpts{
		Marker: marker,
		Limit:  limit,
	}

	objects, err := c.Connection.ObjectNames(container, &opts)

	if err != nil {
		log.Fatal(err)
	}

	for _, o := range objects {
		log.Printf("%s", o)
	}
}
