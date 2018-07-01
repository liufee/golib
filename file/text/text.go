package text

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
)

func NewFile(filename string)  *File{
	return &File{
		filename : filename,
	}
}

type File struct {
	filename string
}

func (f *File) Write(str string, fileMode os.FileMode) bool {
	var file *os.File
	var err error
	if ! f.Exists() {
		file = f.Create()
	}else {
		file, err = os.OpenFile(f.filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, fileMode)
	}
	if( err != nil ){
		fmt.Print(err)
		os.Exit(1)
	}
	_, err = io.WriteString(file, str)
	if(err != nil) {
		return false
	} else {
		return true
	}
}

func (f *File) Exists() bool {
	stat,_ := os.Stat(f.filename)
	if stat == nil {
		return false
	} else {
		return true
	}
}

func (f *File) Create() *os.File {
	file, err := os.Create(f.filename)
	if( err != nil ){
		fmt.Print(err)
		os.Exit(1)
	}
	return file
}

func (f *File) Delete() bool{
	var err error = os.Remove(f.filename)
	if( err != nil ){
		return false
	}
	return true
}

func (f *File) Read() string{
	result, err:= ioutil.ReadFile(f.filename)
	if( err != nil ){
		fmt.Print(err)
		os.Exit(1)
	}
	return string(result)
}
