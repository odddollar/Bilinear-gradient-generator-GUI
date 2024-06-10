package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Custom widget that serves as a spacer
type Spacer struct {
	widget.BaseWidget
	minSize fyne.Size
}

// Creates a new Spacer widget
func NewSpacer(minSize fyne.Size) *Spacer {
	s := &Spacer{minSize: minSize}
	s.ExtendBaseWidget(s)
	return s
}

// Returns a new renderer for the spacer
func (s *Spacer) CreateRenderer() fyne.WidgetRenderer {
	return &spacerRenderer{spacer: s}
}

// Renderer for the Spacer widget
type spacerRenderer struct {
	spacer *Spacer
}

// Returns the minimum size of the spacer
func (r *spacerRenderer) MinSize() fyne.Size {
	return r.spacer.minSize
}

// Does nothing as the spacer doesn't have any child widgets
func (r *spacerRenderer) Layout(size fyne.Size) {}

// Refreshes the spacer
func (r *spacerRenderer) Refresh() {}

// Returns the child widgets of the spacer
func (r *spacerRenderer) Objects() []fyne.CanvasObject {
	return nil
}

// Does nothing as the spacer doesn't hold any resources
func (r *spacerRenderer) Destroy() {}
