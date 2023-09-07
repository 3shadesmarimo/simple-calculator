package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strings"
)

// run the (go mod tidy) command in the terminal to download the required dependencies

type calculatorEntry struct {
	widget.Entry
}

func (c *calculatorEntry) TypedKey(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyBackspace:
		c.Entry.TypedKey(key)

	case fyne.KeyReturn, fyne.KeyEnter:
		result, err := evaluateExpression(c.Text)
		if err != nil {
			return
		}

		resultStr := fmt.Sprintf("%v", result)
		c.SetText(resultStr)
	default:
		c.Entry.TypedKey(key)
	}
}

func newCalculator() *calculatorEntry {
	entry := &calculatorEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")
	//set the frame width and height bigger
	myWindow.Resize(fyne.NewSize(300, 300))

	display := newCalculator()
	display.Alignment = fyne.TextAlignTrailing
	display.SetText("0")
	display.ReadOnly = true

	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), display)
	//make sure there are 4 columns within the calculator app

	var currentVal string

	buttonLeftBracket := widget.NewButton("(", func() {
		currentVal += "("
		display.SetText(currentVal)
	})
	buttonRightBracket := widget.NewButton(")", func() {
		currentVal += ")"
		display.SetText(currentVal)
	})
	buttonPercent := widget.NewButton("%", func() {
		currentVal += "%"
		display.SetText(currentVal)
	})
	buttonCE := widget.NewButton("CE", func() {
		if len(currentVal) > 0 {
			currentVal = currentVal[:len(currentVal)-1]

			if len(currentVal) == 0 {
				display.SetText("0")
			} else {
				display.SetText(currentVal)
			}
		}
	})
	buttons1stRow := fyne.NewContainerWithLayout(layout.NewGridLayout(4), buttonLeftBracket, buttonRightBracket,
		buttonPercent, buttonCE)
	content.AddObject(buttons1stRow)

	button7 := widget.NewButton("7", func() {
		currentVal += "7"
		display.SetText(currentVal)
	})
	button8 := widget.NewButton("8", func() {
		currentVal += "8"
		display.SetText(currentVal)
	})
	button9 := widget.NewButton("9", func() {
		currentVal += "9"
		display.SetText(currentVal)
	})
	buttonDivide := widget.NewButton("÷", func() {
		currentVal += "÷"
		display.SetText(currentVal)
	})
	button2ndRow := fyne.NewContainerWithLayout(layout.NewGridLayout(4), button7, button8, button9, buttonDivide)
	content.Add(button2ndRow)

	button4 := widget.NewButton("4", func() {
		currentVal += "4"
		display.SetText(currentVal)
	})
	button5 := widget.NewButton("5", func() {
		currentVal += "5"
		display.SetText(currentVal)
	})
	button6 := widget.NewButton("6", func() {
		currentVal += "6"
		display.SetText(currentVal)
	})
	buttonMultiply := widget.NewButton("x", func() {
		currentVal += "x"
		display.SetText(currentVal)
	})
	button3rdRow := fyne.NewContainerWithLayout(layout.NewGridLayout(4), button4, button5, button6, buttonMultiply)
	content.Add(button3rdRow)

	button1 := widget.NewButton("1", func() {
		currentVal += "1"
		display.SetText(currentVal)
	})
	button2 := widget.NewButton("2", func() {
		currentVal += "2"
		display.SetText(currentVal)
	})
	button3 := widget.NewButton("3", func() {
		currentVal += "3"
		display.SetText(currentVal)
	})
	buttonMinus := widget.NewButton("-", func() {
		currentVal += "-"
		display.SetText(currentVal)
	})
	button4thRow := fyne.NewContainerWithLayout(layout.NewGridLayout(4), button1, button2, button3, buttonMinus)
	content.Add(button4thRow)

	button0 := widget.NewButton("0", func() {
		currentVal += "0"
		display.SetText(currentVal)
	})
	buttonDot := widget.NewButton(".", func() {
		currentVal += "."
		display.SetText(currentVal)
	})
	buttonEqual := widget.NewButton("=", func() {
		result, err := evaluateExpression(currentVal)
		if err != nil {
			return
		}

		resultStr := fmt.Sprintf("%v", result)
		display.SetText(resultStr)
		currentVal = resultStr
	})
	buttonPlus := widget.NewButton("+", func() {
		currentVal += "+"
		display.SetText(currentVal)
	})
	button5thRow := fyne.NewContainerWithLayout(layout.NewGridLayout(4), button0, buttonDot, buttonEqual, buttonPlus)
	content.Add(button5thRow)

	myWindow.SetContent(content)
	myWindow.Canvas().Focus(display)
	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		display.TypedKey(key)
	})
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}

func evaluateExpression(input string) (float64, error) {
	//Replacing the multiplication and	division symbols with their counterparts
	input = strings.ReplaceAll(input, "x", "*")
	input = strings.ReplaceAll(input, "÷", "/")

	expr, err := govaluate.NewEvaluableExpression(input)
	if err != nil {
		return 0, err
	}

	//Evaluate the parsed expression
	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	return result.(float64), nil
}
