//Adapter模式的使用场景
//有两个接口，这里假设是Shape和TextView
//我们在开发时发现Shape和TextView彼此并不兼容，但他们又必须要在一起协同工作
//然而TextView的接口又不符合当前的要求
//这样我们就需要通过一个适配器adapter来对TextView进行改造，从而达到和Shape一样的接口要求

//Adapter模式，分两种实现方式
//1) 类适配器：可通过类的继承来实现，即adapter公有继承target的所有方法，同步私有继承adaptee的实现
// 		但由于golang中并没有严格的继承机制，因此这种方式在golang中，本人尚未找到好的实现方式
//2）对象适配器：公有实现所有target的方法，同时将一个adaptee对象私有的方式组合到adapter中，
//		我们就是用此方法来实现即首先在adapter（即TextShap）中插入一个私有adaptee对象指针，
//		从而达到隐藏被适配者的目的
//
//
package adapter

type Point struct {
	X, Y int
}

type Coord struct {
}

type Manipulator struct {
}

//target 要适配的目标，就是要使被适配的类具有这个接口的功能
type Shape interface {
	BoundingBox(bottomLeft, topRight Point)
	createmanipulator() *Manipulator
}

//adaptee 被适配的对象，通过一个adapter来转换使这个类具有target类所具有的功能
type TextView struct {
}

//读取原点坐标
func (t *TextView) GetOrigin(x, y *Coord) {
}
func (t *TextView) GetExtent(width, height *Coord) {
}
func (t *TextView) IsEmpty() bool {
	return false
}

//adapter要私有继承被适配对象adaptee（这里是TextView），就是将TextView的细节隐藏起来，
//同时要实现目标接口target的所有方法，将TextView进行转换使之能匹配Shap接口（即target接口）
//从而达到TextView可以通过TextShap能当作Shape使用
type TextShape struct {
	txView *TextView
}

//然后公有的实现target的所有方法，如：
func (t *TextShape) BoundingBox(bottomLeft, topRight Point) {
}
func (t *TextShape) createmanipulator() *Manipulator {
	return nil
}

func (t *TextShape) IsEmpty() bool {
	return t.txView.IsEmpty()
}
