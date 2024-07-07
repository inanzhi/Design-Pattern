package prototype_pattern

import "fmt"

type Folder struct {
	Children []IInode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

// 只克隆一层
func (f *Folder) Clone() IInode {
	//克隆的文件夹A_clone
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []IInode
	//每个子文件和子文件夹进行克隆
	for _, i := range f.Children {
		//如果有父节点中有子节点a，a也会调用a.clone，复制a的子节点，
		inodeCopy := i.Clone()
		tempChildren = append(tempChildren, inodeCopy)
	}
	//克隆后添加到文件夹A_clone
	cloneFolder.Children = tempChildren
	return cloneFolder
}
