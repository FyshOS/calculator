// auto-generated
// Code generated by GUI builder.

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	out  *widget.Label
	bc   *widget.Button
	bbo  *widget.Button
	bbc  *widget.Button
	bd   *widget.Button
	b7   *widget.Button
	b8   *widget.Button
	b9   *widget.Button
	bm   *widget.Button
	b4   *widget.Button
	b5   *widget.Button
	b6   *widget.Button
	bs   *widget.Button
	b1   *widget.Button
	b2   *widget.Button
	b3   *widget.Button
	ba   *widget.Button
	b0   *widget.Button
	bp   *widget.Button
	eval *widget.Button
}

func newGUI() *gui {
	return &gui{}
}

func (g *gui) makeUI() fyne.CanvasObject {
	g.out = widget.NewLabel("0")
	g.bbo = widget.NewButton("(", func() {})
	g.bd = widget.NewButton("/", func() {})
	g.b8 = widget.NewButton("8", func() {})
	g.b9 = widget.NewButton("9", func() {})
	g.b4 = widget.NewButton("4", func() {})
	g.b2 = widget.NewButton("2", func() {})
	g.bp = widget.NewButton(".", func() {})
	g.eval = &widget.Button{Text: "=", Importance: 1, OnTapped: func() {}}
	g.ba = widget.NewButton("+", func() {})
	g.b5 = widget.NewButton("5", func() {})
	g.b3 = widget.NewButton("3", func() {})
	g.b0 = widget.NewButton("0", func() {})
	g.bc = &widget.Button{Text: "C", Importance: 4, OnTapped: func() {}}
	g.bbc = widget.NewButton(")", func() {})
	g.b7 = widget.NewButton("7", func() {})
	g.bm = widget.NewButton("*", func() {})
	g.b6 = widget.NewButton("6", func() {})
	g.bs = widget.NewButton("-", func() {})
	g.b1 = widget.NewButton("1", func() {})

	return container.NewGridWithColumns(1,
		g.out,
		container.NewGridWithColumns(4,
			g.bc,
			g.bbo,
			g.bbc,
			g.bd),
		container.NewGridWithColumns(4,
			g.b7,
			g.b8,
			g.b9,
			g.bm),
		container.NewGridWithColumns(4,
			g.b4,
			g.b5,
			g.b6,
			g.bs),
		container.NewGridWithColumns(4,
			g.b1,
			g.b2,
			g.b3,
			g.ba),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				g.b0,
				g.bp),
			g.eval))
}