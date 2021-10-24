package composite

import (
	"fmt"
)

type Folder struct {
	Name       string
	Components []Component
}

func (folder *Folder) Search(word string) {
	fmt.Println("Searching " + word + " in folder: " + folder.Name)
	for _, file := range folder.Components {
		file.Search(word)
	}
}
