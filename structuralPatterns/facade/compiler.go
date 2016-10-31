//facade模式：就是为一个子系统中的一组（或全部向外的）接口，提供一个简单一致的外观
//这个外观通常就是这个子系统中经常与外部交互的常用接口，目的是让用户更容易的知道
//这个子系统所能提供的最一般的功能，也使得用户容易使用。
//这样做也使用得子系统之间的偶合更松，且各子系统内部也可以更自由的变化
package facade

import (
	"io"
)

type Token string

type Scanner struct {
	input io.Reader
}

func (s *Scanner) Scan() *Token {
	return nil
}

type CodeGenerator struct {
}

type ProgramNode struct {
}

func (p *ProgramNode) GetSourcePosition(line, index int) {

}

func (p *ProgramNode) Add(node *ProgramNode) {

}

func (p *ProgramNode) Remove(node *ProgramNode) {

}

func (p *ProgramNode) Traverse(generator *CodeGenerator) {

}

type ProgramNodeBuilder struct {
	node *ProgramNode
}

func (p *ProgramNodeBuilder) NewVariable(varName string) *ProgramNode {

}

func (p *ProgramNodeBuilder) NewAssignment(varable, expression *ProgramNode) *ProgramNode {

}

func (p *ProgramNodeBuilder) NewReturnStatement(value *ProgramNode) *ProgramNode {

}

func (p *ProgramNodeBuilder) NewCondition(condition, truePart, falsePart *ProgramNode) *ProgramNode {

}

func (p *ProgramNodeBuilder) GetRootNode() *ProgramNode {
	return nil
}

type Parser struct {
}

func (p *Parser) Parse(scanner *Scanner, builder *ProgramNodeBuilder) {

}

type RISCCodeGenerator struct {
	writer io.Writer
}

//以上是编译系统的各个类和函数，为了让外部更容易使用，更容易知道编译子系统的功能，可以给此系统提供一个facade
//即一个Compiler类，在这个类里提供该系统最一般，最常用的方法，这个只有一个Compile方法
//对复杂的系统，会有更多的其他方法
type Compiler struct {
}

func (c *Compiler) Compile(input io.Reader, output io.Writer) {
	scanner := &Scanner{input}
	builder := &ProgramNodeBuilder{}
	parser := &Parser{}

	parser.Parse(scanner, builder)

	Generator := RISCCodeGenerator{output}

	parseTree := builder.GetRootNode()
	parseTree.Traverse(Generator)
}
