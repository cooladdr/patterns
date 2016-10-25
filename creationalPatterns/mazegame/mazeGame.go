package mazegame

import (
	"patterns/creationalPatterns/abstractfactory"
	"patterns/creationalPatterns/builder"
	"patterns/creationalPatterns/maze"
)

//factory method模式：将方法抽象化，并提供一个类似builder模式一样的方法【这里叫CreateMaze】
//来创建和组装各个组件并返回最终的对象；
//通过实现该接口的具体类型的方法定义每个组件的生成，并实现CreateMaze方法将各个组件组合成最终的对象并返回
//抽象工厂模式由调用者来组装最终的对象，这是与抽象工厂的主要区别
//【取builder模式之长，弥补抽象工厂之短；】
type IMazeGame interface {
	//生成各个组件的抽象接口，它会返回最简单的组件对象
	MakeMaze() maze.Maze
	MakeWall() maze.Wall
	MakeRoom(id int) maze.Room
	MakeDoor(r1, r2 maze.Room) maze.Door

	//将各个组件组合成最终需求对象的接口
	//实现该接口的不同类型，可以按自己的需求组装出不同需求的对象出来
	//
	CreateMaze() maze.Maze
}

type MazeGame struct {
}

//+ 接收不同的工厂对象参数，创建具有不同性质属性的组件
// 从而组装成具有相同结构形式的对象maze.Maze，但由于组件的属性不一样
// 最终的maze.Maze也会具有不同的呈现方式
//【由该方法组装并决定各个组件间的组合方式】
func (m *MazeGame) CreateMazeByFactory(factory abstractfactory.MazeFactory) maze.Maze {
	mz := factory.MakeMaze()
	room1 := factory.MakeRoom(1)
	room2 := factory.MakeRoom(2)
	door := factory.MakeDoor(room1, room2)

	mz.AddRoom(room1)
	mz.AddRoom(room2)

	room1.SetSide(maze.North, factory.MakeWall())
	room1.SetSide(maze.East, door)
	room1.SetSide(maze.South, factory.MakeWall())
	room1.SetSide(maze.West, factory.MakeWall())

	room2.SetSide(maze.North, factory.MakeWall())
	room2.SetSide(maze.East, factory.MakeWall())
	room2.SetSide(maze.South, factory.MakeWall())
	room2.SetSide(maze.West, door)

	return mz
}

//builder模式：通过不同的builder，创建不同需要的对象
//【由该方法只决定生成多少个组件，组件间的组合方式由不同的builder决定，这是与抽象工厂的主要区别】
func (m *MazeGame) CreateMazeByBuilder(builder builder.MazeBuilder) maze.Maze {
	builder.BuildMaze()
	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}

func (m *MazeGame) CreateMazeByPrototype(factory abstractfactory.MazeFactory) maze.Maze {
	mz := factory.MakeMaze()
	room1 := factory.MakeRoom(1)
	room2 := factory.MakeRoom(2)
	door := factory.MakeDoor(room1, room2)

	mz.AddRoom(room1)
	mz.AddRoom(room2)

	room1.SetSide(maze.North, factory.MakeWall())
	room1.SetSide(maze.East, door)
	room1.SetSide(maze.South, factory.MakeWall())
	room1.SetSide(maze.West, factory.MakeWall())

	room2.SetSide(maze.North, factory.MakeWall())
	room2.SetSide(maze.East, factory.MakeWall())
	room2.SetSide(maze.South, factory.MakeWall())
	room2.SetSide(maze.West, door)

	return mz
}
