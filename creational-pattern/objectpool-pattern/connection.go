package objectpool_pattern

type Connection struct {
	Id string
}

func (c *Connection) GetID() string {
	return c.Id
}
