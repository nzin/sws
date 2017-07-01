package sws

type SWS_Hr struct {
	SWS_CoreWidget
}

func (self *SWS_Hr) Repaint() {
	self.SWS_CoreWidget.Repaint()
	self.SetDrawColorHex(0xff888888)
	self.DrawLine(0,0,self.Width(),0)
}

func CreateHr(w int32) *SWS_Hr {
	corewidget := CreateCoreWidget(w, 2)
	widget := &SWS_Hr{SWS_CoreWidget: *corewidget,
	}
	return widget
}
