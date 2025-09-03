package main

import "fmt"

// type Engine interface {
// 	Start() string
// }

// type GasolineEngine struct{}
// func (e GasolineEngine) Start() string {
// 	return "GasolineEngine Started"
// }

// type ElectronicEngine struct{}
// func (e ElectronicEngine) Start() string {
// 	return "ElectronicEngineEngine Started"
// }

// type HybridEngine struct{}
// func (e HybridEngine) Start() string {
// 	return "HybridEngineEngineEngine Started"
// }


// type Car struct {
// 	engine Engine // インターフェースで依存関係を定義
// }

// func NewCar(engine Engine) *Car { // 外部から依存関係を注入
// 	return &Car{engine: engine}
// }

// func (c *Car) Start() string {
// 	return c.engine.Start()
// }

// func main() {
// 	ge := NewCar(GasolineEngine{})
// 	el := NewCar(ElectronicEngine{})
// 	hy := NewCar(HybridEngine{})

// 	fmt.Println(ge.Start())
// 	fmt.Println(el.Start())
// 	fmt.Println(hy.Start())
// }

type Engine interface {
	Start() string
}

type GasolineEngine struct {}
func (e GasolineEngine) Start() string {
	return "GasolineEngine Started"
}


type Car struct {
	engine Engine
}

func NewCar(engine Engine) *Car {
	return &Car{engine: engine}
}

func (c *Car) Start() string {
	return c.engine.Start()
}

func main() {
	g := NewCar(GasolineEngine{})
	fmt.Println(g.Start())
}
