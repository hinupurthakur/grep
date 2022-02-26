package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_grep_Fail(t *testing.T) {
	fileName := filepath.Join("grep_testdata", "nonExistantFile.txt")
	regex := "hello"
	_, err := grepFunctionality(fileName, regex)
	var expPathError *os.PathError
	if !errors.As(err, &expPathError) {
		t.Fail()
	}
}

func Test_grep_Success(t *testing.T) {
	fileName := filepath.Join("grep_testdata", "temp.txt")
	regex := "hello"
	result, err := grepFunctionality(fileName, regex)
	if err != nil {
		t.Fail()
	}
	expected := []string{"This file is to check the existence of the word hello.", "Say hello test case."}
	if !reflect.DeepEqual(result, expected) {
		fmt.Printf("expected: %v didn't match result: %v", expected, result)
		t.Fail()
	}
}
