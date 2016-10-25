package maze

//东南西北 四个方向
const (
	None int = iota
	North
	South
	East
	West
)

//进入四个不同方向的操作
type MapSite interface {
	Enter()
}

//房间由四面墙组成
type Room interface {
	MapSite
	GetSide(dir int) MapSite
	SetSide(dir int, site MapSite)
	GetId() int
}

//墙
type Wall interface {
	MapSite
}

//门，用来连接两个房间Room
type Door interface {
	MapSite
	OtherSideFrom(r Room) Room
}

//迷宫
type Maze interface {
	AddRoom(r Room)
	GetRoom(roomNo int) Room
	Info()
}
