//composite模式：定义一个统一的接口，如本例的Equipment，
//然后再定义一个集合类型（该类型也实现Equipment），如本例的CompositeEquipment，
//这样就可以将整个实现了接口Equipment的所有组件，组件成一颗树
package composite

import (
	"fmt"
)

//组合模式常用于类型一台电脑由不同的组件，组合而成的设备
//Equipment定义每个组件都具有的共同操作和属性
type Equipment interface {
	Name() string
	//Power()
	NetPrice() float32
	//DiscountPrice()

	//Add(e Equipment)
	//Remove(e Equipment)
	//AllEquipments() []Equipment
}

//单独的一个软盘设备，要实现Equipment接口
type FloppyDisk struct {
	name string
}

func (f *FloppyDisk) Name() string {
	return f.name
}

func (f *FloppyDisk) NetPrice() float32 {
	return 1.0
}

//单独的一张卡，要实现Equipment接口
type Card struct {
	name string
}

func (c *Card) Name() string {
	return c.name
}

func (c *Card) NetPrice() float32 {
	return 2.0
}

//由多个单独设备组合而成的 ‘组合设备’，要实现Equipment接口
type CompositeEquipment struct {
	name  string
	eList []Equipment
}

func (c *CompositeEquipment) Name() string {
	return c.name
}

func (c *CompositeEquipment) Add(e Equipment) {
	c.eList = append(c.eList, e)
}

func (c *CompositeEquipment) AllEquipments() []Equipment {
	return c.eList
}

func (c *CompositeEquipment) NetPrice() (total float32) {
	for _, e := range c.eList {
		total += e.NetPrice()
	}
	return
}

//底盘
type Chassis struct {
	CompositeEquipment
}

//小格子
type Cabinet struct {
	CompositeEquipment
}

//总线
type Bus struct {
	CompositeEquipment
}

//用各组件组合成一台电脑
func Usage() {
	cabinet := &Cabinet{CompositeEquipment{name: "PC cabinet"}}
	chassis := &Chassis{CompositeEquipment{name: "PC Chassis"}}

	cabinet.Add(chassis)

	bus := &Bus{CompositeEquipment{name: "MCA bus"}}
	bus.Add(&Card{name: "16M token ring"})

	chassis.Add(bus)
	chassis.Add(&FloppyDisk{name: "3.5in Floppy"})

	fmt.Printf("%+#v\n", cabinet)
}
