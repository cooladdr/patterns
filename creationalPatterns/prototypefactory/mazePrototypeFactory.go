package prototypefactory

import (
	"patterns/creationalPatterns/maze"
	_ "reflect"
)

//prototype模式：初始化一个这样的对象的时候，要提供所有成员变量
//这里的每个成员变量就是一个组件原型对象，当要创建新的组件的时候，
//直接拷贝相应的成员变量的类型和值，并返回即可
//所有的成员变量扮演着一个原型库的角色
type MazePrototypeFactory struct {
	M maze.Maze
	W maze.Wall
	R maze.Room
	D maze.Door
}

//需要用反射来实现对象的拷贝，待实现
func (f *MazePrototypeFactory) MakeMaze() maze.Maze {
	return maze.NewNormalMaze("prototyp-factory")
}
func (f *MazePrototypeFactory) MakeWall() maze.Wall {
	return maze.NewNormalWall()
}
func (f *MazePrototypeFactory) MakeRoom(id int) maze.Room {
	return maze.NewNormalRoom(id)
}
func (f *MazePrototypeFactory) MakeDoor(r1, r2 maze.Room) maze.Door {
	return maze.NewNormalDoor(r1, r2)
}
