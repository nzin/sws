package sws

type Hr struct {
	CoreWidget
}

func (self *Hr) Repaint() {
	self.CoreWidget.Repaint()
	self.SetDrawColorHex(0xff888888)
	self.DrawLine(0, 0, self.Width(), 0)
}

func NewHr(w int32) *Hr {
	corewidget := NewCoreWidget(w, 2)
	widget := &Hr{CoreWidget: *corewidget}
	return widget
}
