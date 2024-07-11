package flyweight_pattern

// FlyweightFactory 负责创建和管理 Flyweight 对象
type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := f.flyweights[key]; ok {
		return flyweight
	}
	flyweight := &ConcreteFlyweight{intrinsicState: key}
	f.flyweights[key] = flyweight
	return flyweight
}
