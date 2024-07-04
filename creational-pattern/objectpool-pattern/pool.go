package objectpool_pattern

import (
	"fmt"
	"sync"
)

type Pool struct {
	idle     []IConnection
	active   []IConnection
	capacity int
	muLock   *sync.Mutex
}

// InitPool Initialize the pool
func InitPool(connections []IConnection) (*Pool, error) {
	if len(connections) == 0 {
		return nil, fmt.Errorf("cannot craete a pool of 0 length")
	}
	active := make([]IConnection, 0)
	pool := &Pool{
		//空闲的
		idle: connections,
		//激活的-工作中的
		active: active,
		//容量
		capacity: len(connections),
		muLock:   new(sync.Mutex),
	}
	return pool, nil
}

// Loan 贷款 借
func (p *Pool) Loan() (IConnection, error) {
	p.muLock.Lock()
	defer p.muLock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("no pool object free. Please request after sometime")
	}
	//拿到活跃的第一个连接a
	obj := p.idle[0]
	//活跃的列表-1
	p.idle = p.idle[1:]
	//将a加入激活的连接
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.GetID())
	return obj, nil
}

// Recycle 回收连接
func (p *Pool) Recycle(target IConnection) error {

	p.muLock.Lock()
	defer p.muLock.Unlock()
	//start 从激活切片中移除要被回收的对象
	err := p.Remove(target)
	if err != nil {
		return err
	}
	//end

	//目标对象添加到空闲对象中去
	p.idle = append(p.idle, target)
	fmt.Printf("Recycle Pool Object with ID: %s\n", target.GetID())
	return nil
}

func (p *Pool) Remove(target IConnection) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.GetID() == target.GetID() {
			//激活的连接a回收掉
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			//激活切片长度减1
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("targe pool object doesn't belong to the pool")
}
