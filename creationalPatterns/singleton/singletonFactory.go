package singleton

import (
	"patterns/creationalPatterns/abstractfactory"
	"patterns/creationalPatterns/maze"
)

//singleton模式：控制通过NewMazeSingletonFactory创建的始终只有一个实例

//类型不导出包外，主要防止包外手动创建该类型的对象
type mazeSingletonFactory struct {
	name string
}

//实现abstractfactory.MazeFactory接口，
//以使mazeSingletonFactory对象可通过接口来赋值和访问
func (m *mazeSingletonFactory) MakeMaze() maze.Maze {
	return maze.NewNormalMaze("factory-singleton")
}
func (m *mazeSingletonFactory) MakeWall() maze.Wall {
	return maze.NewNormalWall()
}
func (m *mazeSingletonFactory) MakeRoom(id int) maze.Room {
	return maze.NewNormalRoom(id)
}
func (m *mazeSingletonFactory) MakeDoor(r1, r2 maze.Room) maze.Door {
	return maze.NewNormalDoor(r1, r2)
}

// 唯一的实例，必须用接口来承接，且不导出包外
var singleton abstractfactory.MazeFactory

func NewMazeSingletonFactory(name string) abstractfactory.MazeFactory {
	if singleton == nil {
		f := new(mazeSingletonFactory)
		f.name = name
		singleton = f
	}
	return singleton
}
