package container_test

import (
	"github.com/marcosvm/containeer/container"
	"testing"
)

// START1 OMIT
func TestContainerName(t *testing.T) {
	prefix := "prefix_"
	num := 1
	expected := "prefix_00001"

	if actual := container.ContainerName(prefix, num); actual != expected { // HL
		t.Errorf("got %s but %s was expected", actual, expected) // HL
	}
}

// STOP1 OMIT
