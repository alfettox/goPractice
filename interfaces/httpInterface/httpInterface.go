package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type File interface {
	io.Reader
	io.Writer
	io.Seeker
	Readdir(path string) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

type CustomFile struct {
	data     []byte
	position int
}

func (cf *CustomFile) Read(p []byte) (n int, err error) {
	n = copy(p, cf.data[cf.position:])
	cf.position += n
	if cf.position >= len(cf.data) {
		err = io.EOF
	}
	return n, err
}

func (cf *CustomFile) Write(p []byte) (n int, err error) {
	cf.data = append(cf.data, p...)
	return len(p), nil
}

func (cf *CustomFile) Seek(offset int64, whence int) (int64, error) {
	var newPosition int

	switch whence {
	case io.SeekStart:
		newPosition = int(offset)
	case io.SeekCurrent:
		newPosition = cf.position + int(offset)
	case io.SeekEnd:
		newPosition = len(cf.data) + int(offset)
	default:
		return int64(cf.position), fmt.Errorf("Invalid whence")
	}

	if newPosition < 0 {
		newPosition = 0
	}
	if newPosition > len(cf.data) {
		newPosition = len(cf.data)
	}

	cf.position = newPosition
	return int64(cf.position), nil
}

func (cf *CustomFile) Readdir(path string) ([]os.FileInfo, error) {
	return nil, nil
}


func (cf *CustomFile) Stat() (os.FileInfo, error) {
	return &customFileInfo{cf}, nil
}

type customFileInfo struct {
	cf *CustomFile
}

func (fi *customFileInfo) Name() string {
	return "CustomFile"
}

func (fi *customFileInfo) Size() int64 {
	return int64(len(fi.cf.data))
}

func (fi *customFileInfo) Mode() os.FileMode {
	return os.FileMode(0644)
}

func (fi *customFileInfo) ModTime() time.Time {
	return time.Now()
}

func (fi *customFileInfo) IsDir() bool {
	return false
}

func (fi *customFileInfo) Sys() interface{} {
	return nil
}

func main() {
	customFile := &CustomFile{}

	data := make([]byte, 5)
	_, _ = customFile.Read(data)
	_, _ = customFile.Write([]byte("Hello"))
	_, _ = customFile.Seek(0, io.SeekStart)
	fileInfo, _ := customFile.Stat()

	fmt.Println("File operations completed successfully")

	fmt.Printf("Name: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Mode: %s\n", fileInfo.Mode())
}
