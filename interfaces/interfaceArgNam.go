package main

import (
    "fmt"
)

type File struct {
    name string
    size int
}

type Dependency struct {
    language string
    size     int
}

func main() {
    file1 := File{name: "test", size: 123}
    source := "source_dir"
    destination := "destination_dir"
    test(file1, source, destination)

    dep := Dependency{language: "lib32", size: 1233}
    test(dep, source, destination)
}

func test(m Mover, source string, destination string) {
    result := m.Move(source, destination)
    fmt.Printf("Move result: %d\n", result)
    m.Print()
}

type Mover interface {
    Move(source, destination string) int
    Print()
}

func (file File) Move(source, destination string) int {
    fmt.Printf("Moving file: %s from %s to %s\n", file.name, source, destination)
    return 0
}

func (file File) Print() {
    fmt.Println("File has been moved")
}

func (dep Dependency) Move(source, destination string) int {
    fmt.Printf("Moving dependency: %s from %s to %s\n", dep.language, source, destination)
    return 0
}

func (dep Dependency) Print() {
    fmt.Println("Dependency has been moved")
}
