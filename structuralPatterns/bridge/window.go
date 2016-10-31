//bridge模式：是将接口部分和实现部分进行分开，让他们各自可以自由衍变
//如本例的Window接口，和WindowImp接口
//这样，Window接口就可以自由的实现，从而定义出不同种类的窗口（如本例中的普通窗口和icon窗口等）
//同时，WindowImp接口可以通过不同的实现，来兼容不同操作系统上的窗口
package bridge

import (
	_ "fmt"
)

type View struct {
}

func (v *View) DrawOn() {
}

//窗口接口的定义，具体的实现由WindowImp定义
//即Window中的方法实现的时候，使用WindowImp里定义的方法
//不同平台的窗口系统，可以通过实现这个接口达到多态
type Window interface {
	DrawContents()
	Open()
	Close()
	Iconify()
	Deiconify()

	SetOrigin(at Point)
	SetExtent(extent Point)
	Raise()
	Lower()
	DrawLine(p1, p2 Point)
	DrawRect(p1, p2 Point)
	DrawPolygon(ps []Point, n int)
	DrawText(txt string, p Point)

	//提供一个方法，用来返回实现的接口，即达到桥接的目的
	//getWindowImp可以通过工厂方法来实现...
	getWindowImp(factory WindowSystemFactory) WindowImp
	getView() *View
}

//------------------------------------窗口类型1-----------------------------------------------
//普通的应用窗口
type ApplicationWindow struct{}

func (a *ApplicationWindow) getView() *View {
	return &View{}
}

func (a *ApplicationWindow) getWindowImp(factory WindowSystemFactory) WindowImp {
	return factory.MakeWindowImp()
}

func (a *ApplicationWindow) DrawContents() {
	a.getView().DrawOn()
}

//调用WindowImp中的方法来定义Window的方法
func (a *ApplicationWindow) DrawRect(p1, p2 Point) {
	a.getWindowImp(new(XWindowImpFactory)).DeviceRect(p1.X, p1.Y, p2.X, p2.Y)
}

//-------------------------------------窗口类型2----------------------------------------------
//用来显示icon的窗口
type IconWindow struct {
	name string
}

func (i *IconWindow) getView() *View {
	return &View{}
}

func (i *IconWindow) getWindowImp(factory WindowSystemFactory) WindowImp {
	return factory.MakeWindowImp()
}

func (i *IconWindow) DrawContents() {
	i.getWindowImp(new(XWindowImpFactory)).DeviceBitmap(i.name, Coord(0), Coord(0))
}

//调用WindowImp中的方法来定义Window的方法
func (i *IconWindow) DrawRect(p1, p2 Point) {
	i.getWindowImp(new(XWindowImpFactory)).DeviceRect(p1.X, p1.Y, p2.X, p2.Y)
}
