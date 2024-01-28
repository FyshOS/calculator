package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	loadTheme(a)
	w := a.NewWindow("calculator")

	g := newGUI()
	w.SetContent(g.makeUI())

	g.setupActions()
	w.ShowAndRun()
}

// here you can add some button / callbacks code using widget IDs
func (g *gui) setupActions() {
	g.out.SetText("")
	g.bc.OnTapped = func() {
		g.out.SetText("")
	}
	clearNext := false
	g.eval.OnTapped = func() {
		g.evaluate()
		clearNext = true
	}

	append := func(s string) {
		text := g.out.Text
		if clearNext {
			text = ""
			clearNext = false
		}

		g.out.SetText(text + s)
	}
	digits := []*widget.Button{g.b0, g.b1, g.b2, g.b3, g.b4, g.b5, g.b6, g.b7, g.b8, g.b9}
	for _, b := range digits {
		digit := b
		digit.OnTapped = func() { append(digit.Text) }
	}

	g.ba.OnTapped = func() { append("+") }
	g.bs.OnTapped = func() { append("-") }
	g.bm.OnTapped = func() { append("*") }
	g.bd.OnTapped = func() { append("/") }

	g.bbo.OnTapped = func() { append("(") }
	g.bbc.OnTapped = func() { append(")") }
	g.bp.OnTapped = func() { append(".") }
}

func (g *gui) evaluate() {
	expression, err := govaluate.NewEvaluableExpression(g.out.Text)
	if err != nil {
		g.out.SetText("error")
		return
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		g.out.SetText("error")
		return
	}

	value, ok := result.(float64)
	if !ok {
		g.out.SetText("error")
		return
	}

	g.out.SetText(strconv.FormatFloat(value, 'f', -1, 64))
}
