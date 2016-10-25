package abstractfactory

import (
	"patterns/creationalPatterns/maze"
)

//abstract factory模式：顾名思义将工厂抽象化，而factory method是将方法抽象化 这是两者的主要区别
//抽象工厂：通过实现工厂接口定义不同的工厂类型
//在新类型的方法中定义每个组件的生成方式
//从而创建出一系列符合不同需要的组件
//由调用者将每个组件构建成最终的需求对象

type MazeFactory interface {
	//生成最简单的组件对象
	MakeMaze() maze.Maze
	MakeWall() maze.Wall
	MakeRoom(rId int) maze.Room
	MakeDoor(r1, r2 maze.Room) maze.Door
}
