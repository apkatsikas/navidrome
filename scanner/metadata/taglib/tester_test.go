package taglib

import (
	"fmt"
	"testing"
)

const testFile = "../../../tests/fixtures/test.shn"

func runTest() {
	e := &Extractor{}
	result, err := e.Parse(
		testFile,
	)
	if err != nil {
		fmt.Println(err)
	}
	file := result[testFile]
	fmt.Println(file["title"])
}

func TestHelloName(t *testing.T) {
	runTest()
}
