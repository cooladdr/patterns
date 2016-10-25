package maze

import (
	"fmt"
)

type NormalRoom struct {
	roomId int
	sides  map[int]MapSite
}

func NewNormalRoom(roomId int) *NormalRoom {
	return &NormalRoom{roomId, make(map[int]MapSite)}
}

func (r *NormalRoom) Enter() {
	fmt.Printf("NormalRoom[%d] Enter\n", r.roomId)
}

func (r *NormalRoom) SetSide(dir int, s MapSite) {
	r.sides[dir] = s
}

func (r *NormalRoom) GetSide(dir int) MapSite {
	return r.sides[dir]
}

func (r *NormalRoom) GetId() int {
	return r.roomId
}

type NormalWall struct {
}

func NewNormalWall() *NormalWall {
	return &NormalWall{}
}

func (w *NormalWall) Enter() {
	fmt.Printf("NormalWall Enter\n")
}

type NormalDoor struct {
	r1, r2 Room
	isOpen bool
}

func NewNormalDoor(r1, r2 Room) *NormalDoor {
	return &NormalDoor{r1: r1, r2: r2}
}

func (d *NormalDoor) Enter() {
	fmt.Printf("NormalDoor Enter\n")
}

func (d *NormalDoor) OtherSideFrom(r Room) Room {
	nr := r.(*NormalRoom)
	if nr.GetId() == d.r1.(*NormalRoom).GetId() {
		return d.r2
	}
	if nr.GetId() == d.r2.(*NormalRoom).GetId() {
		return d.r1
	}
	return nil
}

type NormalMaze struct {
	rooms map[int]Room
	name  string
}

func NewNormalMaze(name string) *NormalMaze {
	return &NormalMaze{make(map[int]Room), name}
}

func (m *NormalMaze) AddRoom(r Room) {
	m.rooms[r.GetId()] = r
}

func (m *NormalMaze) GetRoom(id int) Room {
	return m.rooms[id]
}

func (m *NormalMaze) Info() {
	fmt.Printf("%s maze:\n", m.name)
	for id, room := range m.rooms {
		fmt.Printf("room[%d]: ", id)
		room.Enter()
	}
	fmt.Println()
}
