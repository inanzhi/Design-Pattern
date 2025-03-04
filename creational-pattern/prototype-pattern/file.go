package prototype_pattern

import "fmt"

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() IInode {
	return &File{Name: f.Name + "_clone"}
}
