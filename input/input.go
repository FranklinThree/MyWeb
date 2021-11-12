package input

import (
	"MyWeb/universal"
	"fmt"
)

type Input struct {
	Count    int
	Elements []string
}

func NewInput(expr string) (input Input) {
	var es []string
	cs := []rune(expr)
	st := 0
	for i, c := range cs {
		if i == len(cs)-1 {
			if c != ' ' {
				es = append(es, string(cs[st:]))
			}
		}
		if c == ' ' {
			if i == 0 {
				universal.ConsolePrint(universal.Warning, "First character is a space. ", "sentence", expr)
			} else if st == i {
				universal.ConsolePrint(universal.Error, "This is a undefined Error", "st == i ==", i, "sentence", expr)
				continue
			}
			if cs[st] != ' ' {
				es = append(es, string(cs[st:i]))
			}
			st = i + 1
		}

	}
	return Input{len(es), es}
}

type InputSwitcher struct {
	ifs []InputFunction
}

func (is *InputSwitcher) InputSwitch(input Input) {
	var target1 []*InputFunction
	for _, inputFunction := range is.ifs {
		if input.Count == inputFunction.Count {
			target1 = append(target1, &inputFunction)
		}
	}
	target2 := make([]*InputFunction, len(target1))
	length := len(target1)
	count := 0
	for step := 0; step < input.Count; step++ {
		for _, inputFunction := range target1[:length] {
			if inputFunction.IsFixed[step] != 1 || input.Elements[step] == inputFunction.Elements[step] {
				target2 = append(target2, inputFunction)
				count++
			}
		}
		if step != input.Count-1 {
			copy(target1, target2[:count])
			length = count
			count = 0
		}

	}
	fmt.Println(target2[0])
	target2[0].function()
}
func (is *InputSwitcher) ElementMatch(input Input) {

}
