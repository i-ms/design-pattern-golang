package prototype_file

import "fmt"

func App() {
	file1 := &File{name: "File number 1"}
	file2 := &File{name: "File number 2"}
	file3 := &File{name: "File number 3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder number 1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder number 2",
	}

	fmt.Println("\nPrinting hierarchy for Folder 2")
	folder2.print(" ")

	cloneFolder := folder2.clone()
	fmt.Println("\n Printing hierarchy for clone Folder")
	cloneFolder.print(" ")

}
