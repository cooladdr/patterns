//flyweight模式：对那些通常因为数量太大面难以用对象来表示的概念或实体进行建模
//如文档编辑器，我们不可能为每个字符创建一个对象来代表，但对于像英文文档来说，我们可以为每个字母
//创建一个flyweight，每个flyweight存储一个字符代码，因此，逻辑上，文档中的给定字符每次出现
//都有一个对象与之对应；而物理上，文档中出现的每个相同的字符都共享一个flyweight对象，
//每个字符的位置和字体由用户提供相应的场景信息，flyweight只记录字符和绘出来，如RowGlyph知道它
//的子女应该在哪个位置绘制自己，并保证所有子女绘出来后保证是横向排列的，所以RowGlyph可以在绘制
//请求中向每个子女传递它的位置信息
package flyweight

type Window struct {
}

type BTree struct {
}

//用来存储字符的外部状态，如字符的字体等信息
type GlyphContext struct {
	index int
	fonts *BTree
}

func (g *GlyphContext) Next(step int) {

}

func (g *GlyphContext) Insert(quantity int) {

}

func (g *GlyphContext) GetFont() *Font {

}
func (g *GlyphContext) SetFont(f *Font, span int) {

}

type Font struct {
}

//图形化接口，可以绘制自己
type Glyph interface {
	Draw(w *Window, context *GlyphContext)
	SetFont(f *Font, context *GlyphContext)
	GetFont(context *GlyphContext) *Font
	First(context *GlyphContext)
	Next(context *GlyphContext)
	IsDone(context *GlyphContext)
	Current(context *GlyphContext) *Glyph
}

//定义一个flyweight类型，用于存储一个字符代码
//亦即存储字符的内部状态
type Character struct {
	code byte
}

func (c *Character) Draw(w *Window, context *GlyphContext) {

}
