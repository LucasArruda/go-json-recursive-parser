package jsonparser

import (
	"testing" 
)

func Test_RequestUrlWithoutJSON(t *testing.T) {
	uri := "http://letsrevolutionizetesting.com/challenge"
	RequestUrlWithoutJSON(uri)
}