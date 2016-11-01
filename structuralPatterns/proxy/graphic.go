//proxy模式：
package proxy

import (
	"bufio"
	"fmt"
	"io"
)

type Point struct {
	X, Y int
}

type Event struct {
}

//图形对象统一接口
type Graphic interface {
	Draw(p *Point)
	HandleMouse(e *Event)

	//读取图像尺寸
	GetExtent() *Point

	Load(in io.Reader)
	Save(out io.Writer)
}

//实现Graphic接口，用来显示图像文件
type Image struct {
}

func (i *Image) GetExtent() *Point {
	return &Point{}
}

func (i *Image) HandleMouse(e *Event) {

}

//代理对象，也要实现Graphic接口
type ImageProxy struct {
	img      *Image
	extent   *Point
	fileName string
}

func (i *ImageProxy) Draw(p *Point) {
	fmt.Println("这是一个图片", i.fileName)
}

func (i *ImageProxy) GetImage() *Image {
	if i.img == nil {
		i.img = &Image{}
	}
	return i.img
}

func (i *ImageProxy) GetExtent() *Point {
	if i.extent == nil {
		i.extent = i.GetImage().GetExtent()
	}
	return i.extent
}

func (i *ImageProxy) HandleMouse(e *Event) {
	i.GetImage().HandleMouse(e)
}

func (i *ImageProxy) Save(out io.Writer) {
	out.Write([]byte(i.fileName))
}

func (i *ImageProxy) Load(in io.Reader) {
	buffer := bufio.NewReader(in)
	i.fileName, _ = buffer.ReadString('\n')
}

//一个图文并茂的文档，为了使文本内容更流畅和快速的加载，可以使用代理来暂时代表一个图片
//只有当图片真正被浏览或点击时，才让代理去加载并显示图片
type TextDocument struct {
}

func (t *TextDocument) Insert(Graphic) {

}

func Usage() {
	doc := &TextDocument{}
	doc.Insert(&ImageProxy{fileName: "anImageFileName"})
}
