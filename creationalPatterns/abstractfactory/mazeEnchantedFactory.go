package abstractfactory

import (
	"fmt"
	"patterns/creationalPatterns/maze"
)

type EnchantedRoom struct {
	*maze.NormalRoom
}

func (e *EnchantedRoom) CastSpell() {
	fmt.Printf("enchantedRoom[%d] Cast Spell \n", e.GetId())
}

type EnchantedWall struct {
	*maze.NormalWall
}

type EnchantedDoor struct {
	*maze.NormalDoor
}

type EnchantedMaze struct {
	*maze.NormalMaze
}

//实现MazeFactory接口，从而可以创建满足本需求的一系列“魔法迷宫”
type EnchantedMazeFactory struct {
}

func (f *EnchantedMazeFactory) MakeMaze() maze.Maze {
	return &EnchantedMaze{maze.NewNormalMaze("factory-enchanted")}
}
func (f *EnchantedMazeFactory) MakeWall() maze.Wall {
	return &EnchantedWall{maze.NewNormalWall()}
}
func (f *EnchantedMazeFactory) MakeRoom(id int) maze.Room {
	return &EnchantedRoom{maze.NewNormalRoom(id)}
}
func (f *EnchantedMazeFactory) MakeDoor(r1, r2 maze.Room) maze.Door {
	return &EnchantedDoor{maze.NewNormalDoor(r1, r2)}
}
