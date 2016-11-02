package chainofresponsibility

import (
	"fmt"
)

type Topic int

const (
	NO_HELP_TOPIC Topic = iota
	PRINT_TOPIC
	PAPER_ORIENTATION_TOPIC
	APPLICTATON_TOPIC
)

//统一的处理接口
type Handler interface {
	//判断是否需要处理，由具体的实现类型定义其行为
	HasHandler() bool
	//请求处理，由具体的实现类型定义其行为
	Handle()
}

//定义一个帮助处理类型（即一个帮助提供者），实现Handler接口
type HelpHandler struct {
	//指向下一个帮助，以形成一个帮助链
	successor Handler
	//定义帮助的主题编号
	topic Topic
}

//决定请求是否由该对象来处理
func (h *HelpHandler) HasHandler() bool {
	return h.topic != NO_HELP_TOPIC
}

//如果帮助链上有后继者，直接转递给下一个帮助者
func (h *HelpHandler) Handle() {
	if h.successor != nil {
		h.successor.Handle()
	} else {
		fmt.Println("没找到帮助者")
	}
}

//所有的窗口组件都是继承于Widget，同时Widget也继承了HelpHandler，因为所有的窗口都可能需要提供
//一个帮助按钮
type Widget struct {
	HelpHandler
	parent *Widget
}

//定义一个按钮组件，重新实现Handler接口
type Button struct {
	Widget
}

func (b *Button) HasHandler() bool {
	return b.Widget.HasHandler()
}

func (b *Button) Handle() {
	//先判断自身是否需要处理，否则传递给父类，由父类处理或提交给帮助链上的下一下帮助者
	if b.HasHandler() {
		fmt.Println("Button【", b.topic, "】帮助")
	} else {
		b.Widget.Handle()
	}
}

//定义一个对话框，也需要重新实现Handler接口
type Dialog struct {
	Widget
}

func (d *Dialog) HasHandler() bool {
	return d.Widget.HasHandler()
}

func (d *Dialog) Handle() {
	//先判断自身是否需要处理，否则传递给父类，由父类处理或提交给帮助链上的下一下帮助者
	if d.HasHandler() {
		fmt.Println("Dialog【", d.topic, "】帮助")
	} else {
		d.Widget.Handle()
	}
}

//定义一个应用类型，并重新实现Handler接口
type Application struct {
	HelpHandler
}

func (a *Application) HasHandler() bool {
	return a.HelpHandler.HasHandler()
}

func (a *Application) Handle() {
	if a.HasHandler() {
		fmt.Println("Application【", a.topic, "】帮助")
	} else {
		a.HelpHandler.Handle()
	}
}

func Usage() {
	app := new(Application)
	app.topic = NO_HELP_TOPIC //PRINT_TOPIC

	dialog := new(Dialog)
	dialog.successor = app
	dialog.topic = NO_HELP_TOPIC //PAPER_ORIENTATION_TOPIC

	btn := new(Button)
	btn.successor = dialog
	btn.topic = NO_HELP_TOPIC // APPLICTATON_TOPIC

	btn.Handle()

}
