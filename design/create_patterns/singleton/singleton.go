package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *Singleton

func GetInstance() Singleton {
	return nil
}

func (s *singleton) AddOne() int {
	return 0
}
