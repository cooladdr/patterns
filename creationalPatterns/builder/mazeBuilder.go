package builder

import (
	"patterns/creationalPatterns/maze"
)

// builder模式：类似抽象工厂，通过实现MazeBuilder接口定义不同的builder类型
// 在新类型的方法中定义组件的生成方式，及它们的组合
// 从而构建出满足需求的复杂对象，并提供公共方法返回该复杂对象
// 调用者直接得到最终的需求对象
type MazeBuilder interface {
	//生成组件对象，并将各个组件组装起来，达到最终的需求对象目的
	BuildMaze()
	BuildRoom(id int)
	BuildDoor(fromId, toId int)
	//返回最终的需求对象
	GetMaze() maze.Maze
}
