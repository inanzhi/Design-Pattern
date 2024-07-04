package iterator_pattern

// ConcreteAggregate 具体聚合 具体存储元素的切片
type ConcreteAggregate struct {
	Items []interface{}
}

// 维护一个索引可以让迭代器对象独立于集合的内部结构，并提供一种统一的遍历方式。
// ConcreteIterator 具体迭代器，实现 IIterator 接口   当前迭代的值索引   以及存储元素的切片
type ConcreteIterator struct {
	//合计
	aggregate *ConcreteAggregate
	//跟踪当前迭代的位置索引  索引数值和Items的长度相同，只不过不会取该索引值
	currentIndex int
}

func (a *ConcreteAggregate) CreateIterator() IIterator {
	return &ConcreteIterator{aggregate: a, currentIndex: 0}
}

func (i *ConcreteIterator) HasNext() bool {
	return i.currentIndex < len(i.aggregate.Items)
}

func (i *ConcreteIterator) Next() interface{} {
	if i.HasNext() {

		item := i.aggregate.Items[i.currentIndex]
		//达到了长度，于是HasNext中3<3不成立
		i.currentIndex++
		return item
	}
	return nil
}
