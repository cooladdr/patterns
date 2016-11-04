package command

import (
	"fmt"
)

//命令接口
type Command interface {
	Execute()
}

//文档类型，表示一个文档对象
type Document struct {
	filename string
}

func (d *Document) Open() {
	fmt.Println("Document.Open:", d.filename)
}

func (d *Document) Paste() {
	fmt.Println("Document.Paste to", d.filename)
}

//表示一个应用
type Appliction struct {
	docs []*Document
}

func (a *Appliction) Add(d *Document) {
	a.docs = append(a.docs, d)
}

//定义一个打开命令
type OpenCommand struct {
	app      *Appliction
	response *string
}

func (o *OpenCommand) askUser() string {
	return ""
}

func (o *OpenCommand) Execute() {
	filename := o.askUser()

	if len(name) != 0 {
		doc := &Document{name}
		o.app.Add(doc)
		doc.Open()
	}
}

//定义一个粘贴命令
type PasteCommand struct {
	doc *Document
}

func (p *PasteCommand) Execute() {
	p.doc.Paste()
}

//命令集合
type MacroCommand struct {
	cmds []Command
}

func (m *MacroCommand) Add(c Command) {
	m.cmds = append(m.cmds, c)
}

func (m *MacroCommand) Remove(c Command) {
	l := len(m.cmds)
	for i, ic := range m.cmds {
		if ic == c {
			m.cmds[i] = m.cmds[l-1]
			m.cmds = m.cmds[:l-1]
			return
		}
	}
}

func (m *MacroCommand) Execute() {
	for _, c := range m.cmds {
		c.Execute()
	}
}

func Usage() {

}
