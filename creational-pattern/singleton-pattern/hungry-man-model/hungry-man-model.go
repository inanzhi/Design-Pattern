package hungry_man_model

// Singleton 饿汉模式
type Singleton struct {
}

var instance *Singleton = &Singleton{}

func GetInsOr() *Singleton {
	return instance
}
