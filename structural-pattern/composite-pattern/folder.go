package composite_pattern

import "fmt"

// 复合结构
type Folder struct {
	components []component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c component) {
	f.components = append(f.components, c)
}
