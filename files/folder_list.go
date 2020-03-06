package files

import (
	"github.com/FHigher/algorithm/stack"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GetAllByRecursion(dirPath string, level int) []string {
	var files []string
	var levelStr = ""

	// 无需层级前缀时，去掉level参数及此处即可
	if level > 1 {
		levelStr = strings.Repeat("|--", level-1)
	}

	dirFiles, _ := ioutil.ReadDir(dirPath)

	for _, f := range dirFiles {
		fpath := filepath.Join(dirPath, f.Name())
		files = append(files, levelStr+f.Name())
		if f.IsDir() {
			files = append(files, GetAllByRecursion(fpath, level+1)...)
		}
	}

	return files
}

func GetAllByStack(dirPath string) []string {
	var files []string
	var level = 1
	type elem struct {
		path  string
		level int
	}

	fileStack := stack.NewStack(10)
	_ = fileStack.Push(&elem{dirPath, 0})

	for !fileStack.IsEmpty() {
		var (
			dirLock  bool
			fileLock bool
			levelStr string
		)
		fPath, _ := fileStack.Pop()
		fileElem := fPath.(*elem)

		// 当上一个目录层数深,使得level大当前于目录的层级时，此时要从当前目录层级开始记录
		if level > fileElem.level {
			level = fileElem.level
		}

		// 第一层不添加前缀
		if fileElem.level > 0 {
			levelStr = strings.Repeat("|--", fileElem.level-1)
		}

		files = append(files, levelStr+filepath.Base(fileElem.path))

		dirFile, _ := ioutil.ReadDir(fileElem.path)
		for _, f := range dirFile {
			abspath := filepath.Join(fileElem.path, f.Name())
			if f.IsDir() {
				// 锁定，同一层级的目录，level只增加一次
				if !dirLock {
					level++
				}
				dirLock = true
				_ = fileStack.Push(&elem{abspath, level})
			} else {
				// 锁定，同一层级的文件，只追加一次
				if !fileLock {
					levelStr += "|--"
				}
				fileLock = true
				files = append(files, levelStr+f.Name())
			}
		}
	}

	return files
}
