package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	h                   bool
	nowRootDir          string
	oldRootDir          string
	targetRootDir       string
	errorDirName        string
	afterChangeFilePath string
)

func init() {
	flag.BoolVar(&h, "h", false, "显示帮助信息")
	flag.StringVar(&nowRootDir, "n", "D:\\software\\UPUPW_AP7.2_64\\htdocs\\substation", "正在开发的版本根目录")
	flag.StringVar(&oldRootDir, "o", "D:\\software\\UPUPW_AP7.2_64\\htdocs\\substation-old", "旧的版本根目录")
	flag.StringVar(&targetRootDir, "t", "C:\\Users\\EDZ\\Desktop\\2019_04_08", "问题汇总的根目录")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `
fileCreate version: 0.1
author: zrf
email: 1113821597@qq.com
time: 2019/4/9

Usage: fileCreate.exe -n now_root_dir -o old_root_dir -t target_root_dir
Exit: input 'q'

Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	if nowRootDir == "" || oldRootDir == "" || targetRootDir == "" {
		fmt.Println("参数错误")
		return
	}

	if _, err := os.Stat(nowRootDir); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("-n 参数文件路径不存在")
		} else {
			fmt.Println(err)
		}
		return
	}
	if _, err := os.Stat(oldRootDir); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("-o 参数文件路径不存在")
		} else {
			fmt.Println(err)
		}
		return
	}
	if _, err := os.Stat(targetRootDir); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("-t 参数文件路径不存在")
		} else {
			fmt.Println(err)
		}
		return
	}

	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("输入问题目录名：")
		input.Scan()
		if input.Text() == "q" {
			break
		}
		errorDirName = input.Text()

		// 在targetRootDir目录中创建问题目录
		errorDirPath := filepath.Join(targetRootDir, errorDirName)
		if _, err := os.Stat(errorDirPath); os.IsNotExist(err) {
			if err = os.Mkdir(errorDirPath, 0777); err != nil {
				fmt.Println("问题目录创建失败: ", err)
				break
			}
		}
		for {
			fmt.Println("输入修改后的文件相对路径: ")
			input.Scan()
			afterChangeFilePath = input.Text()
			if afterChangeFilePath == "q" {
				break
			}
			// 递归创建路径
			// 修改后的文件相对路径
			fileRelativePath := filepath.Dir(afterChangeFilePath)
			// 目标new目录
			fileNewDirPath := filepath.Join(errorDirPath, "new", fileRelativePath)
			if _, err := os.Stat(fileNewDirPath); os.IsNotExist(err) {
				if err = os.MkdirAll(fileNewDirPath, 0777); err != nil {
					fmt.Println("文件new路径创建失败: ", err)
					break
				}
			}
			// 目标old目录
			fileOldDirPath := filepath.Join(errorDirPath, "old", fileRelativePath)
			if _, err := os.Stat(fileOldDirPath); os.IsNotExist(err) {
				if err = os.MkdirAll(fileOldDirPath, 0777); err != nil {
					fmt.Println("文件old路径创建失败: ", err)
					break
				}
			}

			// 复制文件
			// 文件名
			filename := filepath.Base(afterChangeFilePath)
			// 先从修改后的版本复制到new
			// 修改后的文件绝对路径
			fileNewPath := filepath.Join(nowRootDir, afterChangeFilePath)
			changeFile, err := os.Open(fileNewPath)
			if err != nil {
				fmt.Println("修改后的文件打开失败")
				break
			}
			defer changeFile.Close()
			// 创建目标new路径文件
			newFile, err := os.Create(filepath.Join(fileNewDirPath, filename))
			if err != nil {
				fmt.Println("创建目标new目录文件失败")
				break
			}
			defer newFile.Close()
			// 拷贝修改后的文件到new路径
			byteWritten, err := io.Copy(newFile, changeFile)
			if err != nil {
				fmt.Println("文件拷贝失败")
				break
			}

			if byteWritten > 0 {
				fmt.Println("new 目录文件复制成功")
			}
			// 再从旧版本复制到old
			// 旧文件的绝对路径
			fileOldPath := filepath.Join(oldRootDir, afterChangeFilePath)

			oldFile, err := os.Open(fileOldPath)
			if err != nil {
				fmt.Println("old文件打开失败")
				break
			}
			defer oldFile.Close()

			oldNewFile, err := os.Create(filepath.Join(fileOldDirPath, filename))
			if err != nil {
				fmt.Println("创建目标old目录文件失败")
				break
			}
			defer newFile.Close()

			byteCount, err := io.Copy(oldNewFile, oldFile)
			if err != nil {
				fmt.Println("old文件拷贝失败")
				break
			}

			if byteCount > 0 {
				fmt.Println("old 目录文件复制成功")
			}
		}
	}
}
