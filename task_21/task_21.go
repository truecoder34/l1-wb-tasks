package main

// import "fmt"

// /*
// 	Клиент. Поставить телефон на зарядку.
// 	Воткнуть зарядник в разетку с тремя дырками (Британия - TypeG)
// */
// type client struct{}

// func (c *client) insertTypeGChargerInSocket(soc socketGType) {
// 	soc.insertInTypeGSocket()
// }

// /*===================================*/

// /*
// 	Цель - разетка Брианского типа G в которую можно воткнуть рзарядку с 3мя штеккарми
// */
// type socketGType interface {
// 	insertInTypeGSocket()
// }

// /*===================================*/

// /*
// 	Британский штекер для зарядки . Type G
// */
// type typeG struct{}

// func (tg *typeG) insertInTypeGSocket() {
// 	fmt.Println("Insert type G charger into Type G socket")
// }

// /*===================================*/

// /*
// 	Европейский штекер на для зарядки. Type C.
// */
// type typeC struct{}

// func (tc *typeC) insertInTypeCSocket() {
// 	fmt.Println("Insert type C charger into Type C socket")
// }

// /*===================================*/

// /*
// 	Type C Adapter to Type G
// */
// type typeCAdapter struct {
// 	typeCSocket *typeC
// }

// func (tca *typeCAdapter) insertInTypeGSocket() {
// 	tca.insertInTypeGSocket()
// }

// /*===================================*/

// func main() {
// 	client := &client{} // client init
// 	typeG := &typeG{}

// 	client.insertTypeGChargerInSocket(typeG) //  вставляем в G сокет G зарядку

// 	typeCCharger := &typeC{}              // инициализируем штекер type c
// 	typeCChargerAdapter := &typeCAdapter{ // оборачиваем его адаптером
// 		typeCSocket: typeCCharger,
// 	}

// 	client.insertTypeGChargerInSocket(typeCChargerAdapter)
// }

import "fmt"

type assembly interface {
	insertEngine()
}

type Auto struct{}

func (a *Auto) insertV12(as assembly) {
	as.insertEngine()
}

type EngineV12 struct{}

func (e *EngineV12) insertEngine() {
	fmt.Println("EngineV12")
}

type EngineV8 struct{}

func (e *EngineV8) addEngine() {
	fmt.Println("EngineV8")
}

type EngineV8Adepter struct {
	enginev8 *EngineV8
}

func (e *EngineV8Adepter) insertEngine() {
	e.enginev8.addEngine()
}

func main() {

	auto := &Auto{}
	engineV12 := &EngineV12{}
	auto.insertV12(engineV12)

	engineV8 := &EngineV8{}
	engineV8Adepter := &EngineV8Adepter{
		enginev8: engineV8,
	}

	auto.insertV12(engineV8Adepter)

}
