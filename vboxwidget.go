package sws

//	"github.com/veandco/go-sdl2/sdl"

//
// a simple vertical box container
//
type VBoxWidget struct {
	CoreWidget
}

func (self *VBoxWidget) Resize(width, height int32) {
	if width < 20 {
		width = 20
	}
	if height < 20 {
		height = 20
	}
	self.CoreWidget.Resize(width, height)

	for _, child := range self.children {
		child.Resize(width, child.Height())
	}
	self.PostUpdate()
}

//
// Rebox assumes that children changes from size (from height in fact), and recompute position
func (self *VBoxWidget) Rebox() {
	var width, height int32
	for _, child := range self.children {
		child.Move(0, height)
		if width < child.Width() {
			width = child.Width()
		}
		height += child.Height()
	}
	self.height = height
	self.width = width
	self.CoreWidget.Resize(width, height)
	self.PostUpdate()
}

func (self *VBoxWidget) AddChildTop(child Widget) {
	// in case of it already exist
	self.CoreWidget.RemoveChild(child)
	self.children = append([]Widget{child}, self.children...)
	self.Rebox()
}

func (self *VBoxWidget) AddChild(child Widget) {
	self.CoreWidget.AddChild(child)
	self.Rebox()
}

func (self *VBoxWidget) RemoveChild(child Widget) {
	self.CoreWidget.RemoveChild(child)
	var width, height int32
	for _, child := range self.children {
		child.Move(0, height)
		if width < child.Width() {
			width = child.Width()
		}
		height += child.Height()
	}
	self.height = height
	self.width = width
	self.PostUpdate()
}

func (self *VBoxWidget) RemoveAllChildren() {
	for _, child := range self.children {
		self.CoreWidget.RemoveChild(child)
	}
	self.width = 0
	self.height = 0
	self.PostUpdate()
}

func NewVBoxWidget(w, h int32) *VBoxWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &VBoxWidget{CoreWidget: *corewidget}
	return widget
}
