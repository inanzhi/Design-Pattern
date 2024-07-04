package prototype_pattern

// 表示文件或文件夹
type IInode interface {
	Print(string)
	Clone() IInode
}
