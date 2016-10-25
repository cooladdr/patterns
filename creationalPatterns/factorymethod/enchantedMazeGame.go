package factorymethod

import (
	"patterns/creationalPatterns/abstractfactory"
	"patterns/creationalPatterns/maze"
)

//有魔法效果的迷宫游戏
type EnchantedMazeGame struct {
}

func (m *EnchantedMazeGame) MakeMaze() maze.Maze {
	return &abstractfactory.EnchantedMaze{maze.NewNormalMaze("factory-method-enchanted ")}
}

func (m *EnchantedMazeGame) MakeWall() maze.Wall {
	return &abstractfactory.EnchantedWall{maze.NewNormalWall()}
}

func (m *EnchantedMazeGame) MakeRoom(id int) maze.Room {
	return &abstractfactory.EnchantedRoom{maze.NewNormalRoom(id)}
}

func (m *EnchantedMazeGame) MakeDoor(r1, r2 maze.Room) maze.Door {
	return &abstractfactory.EnchantedDoor{maze.NewNormalDoor(r1, r2)}
}

func (m *EnchantedMazeGame) CreateMaze() maze.Maze {
	mz := m.MakeMaze()
	r1, r2, r3 := m.MakeRoom(1), m.MakeRoom(2), m.MakeRoom(3)
	d1 := m.MakeDoor(r1, r2)
	d2 := m.MakeDoor(r2, r3)

	mz.AddRoom(r1)
	mz.AddRoom(r2)
	mz.AddRoom(r3)

	r1.SetSide(maze.North, maze.NewNormalWall())
	r1.SetSide(maze.East, maze.NewNormalWall())
	r1.SetSide(maze.South, maze.NewNormalWall())
	r1.SetSide(maze.West, d1)

	r2.SetSide(maze.North, maze.NewNormalWall())
	r2.SetSide(maze.East, d1)
	r2.SetSide(maze.South, maze.NewNormalWall())
	r2.SetSide(maze.West, d2)

	r3.SetSide(maze.North, maze.NewNormalWall())
	r3.SetSide(maze.East, d2)
	r3.SetSide(maze.South, maze.NewNormalWall())
	r3.SetSide(maze.West, maze.NewNormalWall())
	return mz
}
