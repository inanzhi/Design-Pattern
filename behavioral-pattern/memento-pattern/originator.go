package memento_pattern

// IOriginator 用户接口
type IOriginator interface {
	CreateMemento() *IMemento
	RestoreMemento() *IMemento
}

// Originator 用户
type Originator struct {
	OriginatorState string
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{mementoState: o.OriginatorState}
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.OriginatorState = m.GetSavedState()
}

func (o *Originator) SetState(state string) {
	o.OriginatorState = state
}

func (o *Originator) GetState() string {
	return o.OriginatorState
}
