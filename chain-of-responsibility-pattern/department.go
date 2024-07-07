package chain_of_responsibility_pattern

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
