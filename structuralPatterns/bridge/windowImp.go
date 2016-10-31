package bridge

import (
	_ "fmt"
)

type Point struct {
	X, Y int
}

type Coord int

//窗口接口的实现接口
//主要是用来实现 兼容不同系统平台的窗口系统的具体实现
type WindowImp interface {
	ImpTop()
	ImpBottom()
	ImpSetExtent(p Point)
	ImpSetOrigin(p Point)

	DeviceRect(int, int, int, int)
	DeviceText(string, Coord, Coord)
	DeviceBitmap(string, Coord, Coord)
}

//在X Window System系统上的窗口系统
type XWindowImp struct {
	//该类型原则上 要实现所有这些接口，为简单起见，这里直接插入WindowImp这个接口
	WindowImp
}

//在IBM的Presentation Manager系统上的窗口系统
type PMWindowImp struct {
	//该类型原则上 要实现所有这些接口，为简单起见，这里直接插入WindowImp这个接口
	WindowImp
}

type WindowSystemFactory interface {
	MakeWindowImp() WindowImp
}

type XWindowImpFactory struct {
}

func (x *XWindowImpFactory) MakeWindowImp() WindowImp {
	return new(XWindowImp)
}

type PMWindowImpFactory struct {
}

func (p *PMWindowImpFactory) MakeWindowImp() WindowImp {
	return new(PMWindowImp)
}
