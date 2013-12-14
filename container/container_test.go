package container

import "testing"

func testContainerName(t *testing.T) {
	expected := "prefix_0002"
	if actual := ContainerName(expected, 1); actual != expected {
		t.Errorf("got %s but %s was expected", actual, expected)
	}
}
