package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type stack []float64

func (s *stack) Push(str float64) {
	*s = append(*s, str)
}

func (s *stack) Pop() float64 {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

type calc struct {
	equation string
	input    *widget.Entry
	output   *widget.Entry
	window   fyne.Window
}

func (c *calc) loadGUI(app fyne.App, s stack) {
	var selectedOutput []string
	var selectedInput []string
	c.window = app.NewWindow("RPN64")
	c.window.SetFixedSize(true)

	c.input = widget.NewEntry()
	c.output = widget.NewEntry()
	c.output.Disable()
	c.output.SetText("Output:")

	c.window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		c.input,
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			c.output,
			widget.NewButton(".", func() {
				selectedInput = []string{c.input.Text, "."}
				c.input.SetText(strings.Join(selectedInput, ""))
			})),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			widget.NewButton("1", func() {
				selectedInput = []string{c.input.Text, "1"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("2", func() {
				selectedInput = []string{c.input.Text, "2"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("3", func() {
				selectedInput = []string{c.input.Text, "3"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("+", func() {
				selectedInput = []string{c.input.Text, "+"}
				c.input.SetText(strings.Join(selectedInput, ""))
			})),

		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			widget.NewButton("4", func() {
				selectedInput = []string{c.input.Text, "4"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("5", func() {
				selectedInput = []string{c.input.Text, "5"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("6", func() {
				selectedInput = []string{c.input.Text, "6"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("-", func() {
				selectedInput = []string{c.input.Text, "-"}
				c.input.SetText(strings.Join(selectedInput, ""))
			})),

		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			widget.NewButton("7", func() {
				selectedInput = []string{c.input.Text, "7"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("8", func() {
				selectedInput = []string{c.input.Text, "8"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("9", func() {
				selectedInput = []string{c.input.Text, "9"}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("/", func() {
				selectedInput = []string{c.input.Text, "/"}
				c.input.SetText(strings.Join(selectedInput, ""))
			})),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			widget.NewButton("*", func() {
				selectedInput = []string{c.input.Text, "."}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton(" ", func() {
				selectedInput = []string{c.input.Text, " "}
				c.input.SetText(strings.Join(selectedInput, ""))
			}),
			widget.NewButton("C", func() {
				c.input.SetText("")
			}),
			widget.NewButton("=", func() {
				c.equation = c.input.Text
				selectedOutput = []string{"Output:", c.calculate(c.equation, s)}
				c.output.SetText(strings.Join(selectedOutput, " "))
			}))))

	c.window.Resize(fyne.NewSize(300, 400))
	c.window.ShowAndRun()
}

func (c *calc) calculate(equation string, s stack) string {
	e := strings.Fields(c.equation)
	for iterations, num := range e {
		if iterations == len(e) {
			break
		}
		if IsLetter(num) {
			switch num {
			case "+":
				p1 := s.Pop()
				p2 := s.Pop()
				s.Push(float64(p1) + float64(p2))
			case "-":
				p1 := s.Pop()
				p2 := s.Pop()
				s.Push(float64(p2) - float64(p1))
			case "*":
				p1 := s.Pop()
				p2 := s.Pop()
				s.Push(float64(p1) * float64(p2))
			case "/":
				p1 := s.Pop()
				p2 := s.Pop()
				s.Push(float64(p2) / float64(p1))
			default:
				c.output.SetText("Output: Error!")
			}
		} else {
			n2, _ := strconv.ParseFloat(num, 64)
			s.Push(n2)
		}
		log.Println("stack is", s)
	}
	result := s.Pop()
	resultStr := FloatToString(result)
	return resultStr

}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func IsLetter(v string) bool {
	var flag = false
	if _, err := strconv.Atoi(v); err == nil {
		return false
	} else {
		for _, c := range v {
			if !unicode.IsDigit(c) && c != '.' {
				flag = true
			}
		}
		return flag
	}
}

func newCalculator() *calc {
	c := &calc{}
	return c
}

func main() {
	var s stack
	c := newCalculator()
	c.loadGUI(app.New(), s)
}
