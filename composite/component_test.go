package composite

import (
	"testing"
)

func TestCompositePatter(t *testing.T) {
	/*
			Folder 2
			/       \
		File 3		Folder 1
					/	  \
				File 1	 File 2
	*/
	file1 := &File{Name: "File 1"}
	file2 := &File{Name: "File 2"}
	folder1 := &Folder{
		Name:       "Folder 1",
		Components: []Component{file1, file2},
	}
	file3 := &File{Name: "File 3"}
	folder2 := &Folder{
		Name:       "Folder 2",
		Components: []Component{file3, folder1},
	}
	folder2.Search("foo")
}
