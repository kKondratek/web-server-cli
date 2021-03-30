package main

import (
	"testing"
)

func TestValidatePath(t *testing.T) {
	var validPathWithFile = "sample/index.html"
	var validPath = "sample/"
	var invalidPathWithFile = "invalid/index.html"
	var invalidPath = "invalid/"

	var resultPath string
	var err error

	resultPath, err = validatePath(validPathWithFile)
	if resultPath != validPath || err != nil {
		t.Log("validPathWithFile was given, should return validPath and no error")
		t.Fail()
	}

	err = nil
	resultPath, err = validatePath(validPath)
	if resultPath != validPath || err != nil {
		t.Log("validPath was given, should return validPath and no error")
		t.Fail()
	}

	err = nil
	_, err = validatePath(invalidPathWithFile)
	if err == nil {
		t.Log("invalidPathWithFile was given, should return error")
		t.Fail()
	}

	err = nil
	_, err = validatePath(invalidPath)
	if err == nil {
		t.Log("invalidPath was given, should return error")
		t.Fail()
	}
}
