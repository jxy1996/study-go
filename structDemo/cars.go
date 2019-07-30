package structDemo

type Car struct {
	Model        string
	BuildYear    int
	Manufacturer string
}

type Cars []*Car

type Any interface{}

//高阶函数
//直接操作所有汽车
func (cs Cars) Process(f func(car *Car)) {
	for _, v := range cs {
		f(v)
	}
}

//获取特定子集
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(car *Car) {
		if f(car) {
			cars = append(cars, car)
		}
	})
	return cars
}

func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	i := 0
	cs.Process(func(car *Car) {
		result[i] = f(car)
		i++
	})
	return result
}
