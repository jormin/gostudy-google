package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseUserList(t *testing.T) {
	b,_ := ioutil.ReadFile("city.txt")
	ParseUserList(string(b))
}
