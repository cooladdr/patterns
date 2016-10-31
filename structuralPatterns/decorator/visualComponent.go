//decorator模式，其实就是对类进行了重新包裹封装（wrapper）
package decorator

import (
	"fmt"
)

//可视化组件的接口定义
type VisualComponent interface {
	Draw()
	Resize()
}

//------------------------------------------------------
//文本组件
type TextView struct {
	content string
}

//实现可视化接口
func (t *TextView) Draw() {
	fmt.Println(t.content)
}

func (t *TextView) Resize() {
	fmt.Println(t.content, "---尺寸变化")
}

//------------------------------------------------------
//普通装饰
type Decorator struct {
	vc VisualComponent
}

func (d *Decorator) Draw() {
	d.vc.Draw()
	fmt.Println("这是一个普通的组件")
}

func (d *Decorator) Resize() {
	fmt.Println("尺寸变化")
}

//------------------------------------------------------
//组件有边框
type BorderDecorator struct {
	vc    VisualComponent
	width int
}

func (b *BorderDecorator) Draw() {
	b.vc.Draw()
	fmt.Println("给组件添加‘边框’，边框宽度为：", b.width)
}

func (b *BorderDecorator) Resize() {
	fmt.Println("尺寸变化")
}

//------------------------------------------------------
//组件有滚动条
type ScrollDecorator struct {
	vc VisualComponent
}

func (s *ScrollDecorator) Draw() {
	s.vc.Draw()
	fmt.Println("给组件添加‘滚动条’")
}

func (b *ScrollDecorator) Resize() {
	fmt.Println("尺寸变化")
}

//------------------------------------------------------
//组件有阴影
type DropShadowDecorator struct {
	vc VisualComponent
}

func (d *DropShadowDecorator) Draw() {
	d.vc.Draw()
	fmt.Println("给组件添加‘阴影’")
}

func (d *DropShadowDecorator) Resize() {
	fmt.Println("尺寸变化")
}

func Usage() {
	ex := &DropShadowDecorator{&ScrollDecorator{&BorderDecorator{&Decorator{&TextView{"I am Jeremy"}}, 23}}}
	ex.Draw()

}
