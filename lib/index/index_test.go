package index

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var data = `
remotes:
  - abc
  - xxx
items:
  -
    name: github.com/foo/abc
    repo: abc/bar
    ref: asdasdas
    installed: true
  -
    name: github.com/foo2/abc2
    repo: abc2/barx
    ref: asdasdas
`

func TestIndex(t *testing.T) {
	idx := Index{}
	err := yaml.Unmarshal([]byte(data), &idx)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	assert.IsType(t, new(Index), &idx)
	// assert.Equal(t, idx.Remotes[0].Name, "abc")
}

func TestLoadIdex(t *testing.T) {
	idx := LoadIndex("testdata/index.yaml")

	assert.IsType(t, new(Index), &idx)
}
