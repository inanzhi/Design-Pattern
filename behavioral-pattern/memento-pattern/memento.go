package memento_pattern

type IMemento interface {
	GetSavedState()
}

// Memento 备忘录  存储的是用户的状态、信息
type Memento struct {
	mementoState string
}

func (m *Memento) GetSavedState() string {
	return m.mementoState
}
