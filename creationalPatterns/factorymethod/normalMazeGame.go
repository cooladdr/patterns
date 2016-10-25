package factorymethod

import (
	"patterns/creationalPatterns/maze"
)

//墙已经被炸，且房间里有个炸弹的迷宫游戏
type NormalMazeGame struct {
}

func (m *NormalMazeGame) MakeMaze() maze.Maze {
	return maze.NewNormalMaze("factory-method-normal ")
}

func (m *NormalMazeGame) MakeWall() maze.Wall {
	return maze.NewNormalWall()
}

func (m *NormalMazeGame) MakeRoom(id int) maze.Room {
	return maze.NewNormalRoom(id)
}

func (m *NormalMazeGame) MakeDoor(r1, r2 maze.Room) maze.Door {
	return maze.NewNormalDoor(r1, r2)
}

func (m *NormalMazeGame) CreateMaze() maze.Maze {
	mz := m.MakeMaze()
	r1, r2 := m.MakeRoom(1), m.MakeRoom(2)
	d := m.MakeDoor(r1, r2)

	mz.AddRoom(r1)
	mz.AddRoom(r2)

	r1.SetSide(maze.North, maze.NewNormalWall())
	r1.SetSide(maze.East, d)
	r1.SetSide(maze.South, maze.NewNormalWall())
	r1.SetSide(maze.West, maze.NewNormalWall())

	r2.SetSide(maze.North, maze.NewNormalWall())
	r2.SetSide(maze.East, maze.NewNormalWall())
	r2.SetSide(maze.South, maze.NewNormalWall())
	r2.SetSide(maze.West, d)
	return mz
}
