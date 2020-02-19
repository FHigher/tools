package files

import (
	"github.com/FHigher/algorithm/stack"
	"io/ioutil"
	"path/filepath"
)

func GetAllByRecursion(dirPath string) []string {
	var files []string
	dirFiles, _ := ioutil.ReadDir(dirPath)

	for _, f := range dirFiles {
		fpath := filepath.Join(dirPath, f.Name())
		files = append(files, fpath)
		if f.IsDir() {
			files = append(files, GetAllByRecursion(fpath)...)
		}
	}

	return files
}

func GetAllByStack(dirPath string) []string {
	var files []string

	fileStack := stack.NewStack(10)
	_ = fileStack.Push(dirPath)

	for !fileStack.IsEmpty() {
		fPath, _ := fileStack.Pop()
		files = append(files, fPath.(string))
		dirFile, _ := ioutil.ReadDir(fPath.(string))

		for _, f := range dirFile {
			abspath := filepath.Join(fPath.(string), f.Name())
			if f.IsDir() {
				_ = fileStack.Push(abspath)
			} else {
				files = append(files, abspath)
			}
		}
	}

	return files
}
