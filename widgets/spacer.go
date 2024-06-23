package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Custom widget that serves as spacer
type Spacer struct {
	widget.BaseWidget
	minSize fyne.Size
}

// Creates new Spacer widget
func NewSpacer(minSize fyne.Size) *Spacer {
	s := &Spacer{minSize: minSize}
	s.ExtendBaseWidget(s)
	return s
}

// Returns new renderer for Spacer
func (s *Spacer) CreateRenderer() fyne.WidgetRenderer {
	return &spacerRenderer{spacer: s}
}

// Renderer for Spacer widget
type spacerRenderer struct {
	spacer *Spacer
}

// Returns minimum size of Spacer
func (r *spacerRenderer) MinSize() fyne.Size {
	return r.spacer.minSize
}

// Does nothing as Spacer doesn't have any child widgets
func (r *spacerRenderer) Layout(size fyne.Size) {}

// Refreshes Spacer
func (r *spacerRenderer) Refresh() {}

// Returns child widgets of Spacer
func (r *spacerRenderer) Objects() []fyne.CanvasObject {
	return nil
}

// Does nothing as Spacer doesn't hold any resources
func (r *spacerRenderer) Destroy() {}
