package composite_file

import "fmt"

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for a keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}
