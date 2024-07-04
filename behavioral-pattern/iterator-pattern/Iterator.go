package iterator_pattern

// IIterator 接口，定义访问和遍历元素的方法
type IIterator interface {
	HasNext() bool
	Next() interface{}
}
