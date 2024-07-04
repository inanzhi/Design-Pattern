package objectpool_pattern

// IConnection 中存有多个connection
type IConnection interface {
	//This is any id which can be used to compare two different pool objects
	GetID() string
}
