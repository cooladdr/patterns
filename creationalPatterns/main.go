// maze project main.go
package main

import (
	"fmt"
	"patterns/creationalPatterns/abstractfactory"
	"patterns/creationalPatterns/builder"
	"patterns/creationalPatterns/factorymethod"
	"patterns/creationalPatterns/maze"
	"patterns/creationalPatterns/mazegame"
	"patterns/creationalPatterns/prototypefactory"
	"patterns/creationalPatterns/singleton"
	"reflect"
)

func main() {
	game := new(mazegame.MazeGame)

	//abstract factory 模式
	var factory abstractfactory.MazeFactory
	var aMaze maze.Maze

	//创建“普通迷宫”
	factory = new(abstractfactory.NormalMazeFactory)
	aMaze = game.CreateMazeByFactory(factory)
	fmt.Println(reflect.TypeOf(aMaze))
	aMaze.Info()

	//创建“魔法迷宫”--其结构和组成与“普通迷宫”一样，只是门需要口语才能打开，且房间里会有魔法钥匙或口语
	factory = new(abstractfactory.EnchantedMazeFactory)
	aMaze = game.CreateMazeByFactory(factory)
	fmt.Println(reflect.TypeOf(aMaze))
	aMaze.Info()

	room := aMaze.GetRoom(1)
	//通过接口断言将Room接口转成*EnchantedRoom
	enchantedRoom := room.(*abstractfactory.EnchantedRoom)
	enchantedRoom.CastSpell()
	fmt.Println("————————————————————————————————————————————————————————————————————————")

	//builder 模式
	//类似抽象工厂，也是通过不同的builder来创建不同的对象
	aMaze = game.CreateMazeByBuilder(new(builder.NormalMazeBuilder))
	fmt.Println(reflect.TypeOf(aMaze))
	aMaze.Info()

	fmt.Println("————————————————————————————————————————————————————————————————————————")

	//factory method模式
	var iGame mazegame.IMazeGame
	iGame = new(factorymethod.NormalMazeGame)
	aMaze = iGame.CreateMaze()
	fmt.Println(reflect.TypeOf(aMaze))
	aMaze.Info()

	iGame = new(factorymethod.EnchantedMazeGame)
	aMaze = iGame.CreateMaze()
	fmt.Println(reflect.TypeOf(aMaze))
	aMaze.Info()

	fmt.Println("————————————————————————————————————————————————————————————————————————")

	//prototype模式
	factory = &prototypefactory.MazePrototypeFactory{
		maze.NewNormalMaze("prototyp-factory"),
		maze.NewNormalWall(),
		abstractfactory.EnchantedRoom{maze.NewNormalRoom(0)},
		abstractfactory.EnchantedDoor{maze.NewNormalDoor(maze.NewNormalRoom(1), maze.NewNormalRoom(2))}}
	aMaze = game.CreateMazeByPrototype(factory)
	fmt.Println(reflect.TypeOf(aMaze))
	aMaze.Info()

	fmt.Println("————————————————————————————————————————————————————————————————————————")

	//singleton模式
	factory = singleton.NewMazeSingletonFactory("first singleton2")
	factory2 := singleton.NewMazeSingletonFactory("first singleton3")
	fmt.Printf("%t, %t, %t\n", factory == factory2,
		reflect.TypeOf(factory) == reflect.TypeOf(factory2),
		reflect.ValueOf(factory) == reflect.ValueOf(factory2))

	//f := singleton.mazeSingletonFactory{}
	//f.MakeWall()

	aMaze = game.CreateMazeByFactory(factory2)
	fmt.Println(reflect.TypeOf(aMaze))
	aMaze.Info()

	fmt.Println("————————————————————————————————————————————————————————————————————————")

}
