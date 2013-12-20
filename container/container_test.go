package container

import (
	"testing"
)

func TestContainerName(t *testing.T) {

  cc := Container{
  }

	prefix := "prefix_"
	num := 1
	expected := "prefix_00001"

	if actual := cc.ContainerName(&prefix, num); actual != expected { // HL
		t.Errorf("got %s but %s was expected", actual, expected) // HL
	}
}
