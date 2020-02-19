package files

import (
	"fmt"
	"testing"
)

var dirPath = "C:\\Users\\Administrator\\Desktop\\code\\algorithm"

func TestGetAllByRecursion(t *testing.T) {
	files := GetAllByRecursion(dirPath)
	for _, f := range files {
		fmt.Println(f)
	}
}

func TestGetAllByStack(t *testing.T) {
	files := GetAllByStack(dirPath)
	for _, f := range files {
		fmt.Println(f)
	}
}
