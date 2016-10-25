package builder

import (
	_ "fmt"
	"patterns/creationalPatterns/maze"
)

type NormalMazeBuilder struct {
	maze *maze.NormalMaze
}

func (b *NormalMazeBuilder) CommonWall(r1, r2 maze.Room) int {
	return maze.South
}

func (b *NormalMazeBuilder) BuildMaze() {
	b.maze = maze.NewNormalMaze("builder-normal")
}

func (b *NormalMazeBuilder) BuildRoom(id int) {
	if b.maze.GetRoom(id) == nil {
		room := maze.NewNormalRoom(id)

		room.SetSide(maze.North, maze.NewNormalWall())
		room.SetSide(maze.South, maze.NewNormalWall())
		room.SetSide(maze.East, maze.NewNormalWall())
		room.SetSide(maze.West, maze.NewNormalWall())

		b.maze.AddRoom(room)
	}
}

func (b *NormalMazeBuilder) BuildDoor(fromId, toId int) {
	r1, r2 := b.maze.GetRoom(fromId), b.maze.GetRoom(toId)
	door := maze.NewNormalDoor(r1, r2)

	r1.SetSide(b.CommonWall(r1, r2), door)
	r2.SetSide(b.CommonWall(r2, r1), door)

}

func (b *NormalMazeBuilder) GetMaze() maze.Maze {
	return b.maze
}
