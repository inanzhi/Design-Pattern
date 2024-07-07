package memento_pattern

// 管理员
type Caretaker struct {
	MementoList []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.MementoList = append(c.MementoList, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	return c.MementoList[index]
}
