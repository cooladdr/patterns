package abstractfactory

import (
	_ "fmt"
	"patterns/creationalPatterns/maze"
)

//实现MazeFactory接口，从而可以创建满足本需求的一系列“普通迷宫”
type NormalMazeFactory struct {
}

func (f *NormalMazeFactory) MakeMaze() maze.Maze {
	return maze.NewNormalMaze("factory-normal")
}
func (f *NormalMazeFactory) MakeWall() maze.Wall {
	return maze.NewNormalWall()
}
func (f *NormalMazeFactory) MakeRoom(id int) maze.Room {
	return maze.NewNormalRoom(id)
}
func (f *NormalMazeFactory) MakeDoor(r1, r2 maze.Room) maze.Door {
	return maze.NewNormalDoor(r1, r2)
}
