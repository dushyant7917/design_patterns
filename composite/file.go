package composite

import (
	"fmt"
)

type File struct {
	Name string
}

func (file *File) Search(word string) {
	fmt.Println("Searching " + word + " in file: " + file.Name)
}
