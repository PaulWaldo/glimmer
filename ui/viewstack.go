package ui

import (
	"errors"

	"fyne.io/fyne/v2"
)

// ViewStack represents a stack data structure
type ViewStack struct {
	elements []*fyne.Container
	window   fyne.Window
}

// NewViewStack creates a new stack
func NewViewStack(window fyne.Window) *ViewStack {
	return &ViewStack{
		elements: []*fyne.Container{},
		window:   window,
	}
}

// First removes all elements from the stack except the first and displays it
// It does nothing if the stack is empty
func (v *ViewStack) First() {
	if !v.IsEmpty() {
		v.window.SetContent(v.elements[0])
	}
}

// Pop removes and returns the top element of the stack
// It returns an error if the stack is empty
func (v *ViewStack) Push(view *fyne.Container) {
	v.elements = append(v.elements, view)
	v.window.SetContent(view)
}

// Pop removes the top element of the stack, displays it in the window, and returns the element
// It returns an error if the stack is empty
func (v *ViewStack) Pop() (*fyne.Container, error) {
	if v.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	// Get the last element
	index := len(v.elements) - 1
	// Remove the last element
	v.elements = v.elements[:index]

	// Show the new top element if there is one
	var element *fyne.Container
	if index > 0 {
		element := v.elements[index-1]
		v.window.SetContent(element)
	}
	return element, nil
}

// IsEmpty checks if the stack is empty
func (v *ViewStack) IsEmpty() bool {
	return len(v.elements) == 0
}
