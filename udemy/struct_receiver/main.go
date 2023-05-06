package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Title    string
	Estimate int
}

func funcDefer() {
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}

func trimExtension(files ...File) (res []File) {
	for _, file := range files {
		name := strings.Split(file.Name, ".")[0]
		res = append(res, File{Name: name})
	}
	return
}

func fileChecker(file File) (name string, er error) {
	name, er = "", nil
	f, err := os.Open(file.Name)
	defer f.Close()

	if err != nil {
		er = errors.New("file not found")
		return
	}

	name = f.Name()
	return
}

type File struct {
	Name string
}

func (f File) String() string {
	return "filename:" + f.Name
}

func main() {
	files := []File{
		{Name: "index.html"},
		{Name: "main.go"},
		{Name: "app.js"},
	}
	fmt.Println(trimExtension(files...))
	fmt.Println(fileChecker(File{Name: "aaaa.go"}))

}
