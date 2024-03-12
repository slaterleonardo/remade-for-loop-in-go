package main

import "fmt"

type condition func(f *ForLoop) bool
type part func(f *ForLoop)

type ForLoop struct {
	value     int
	condition condition
	iterator  part
	block     part
	done      part
}

func (f *ForLoop) Run() {
	if !f.condition(f) {
		f.done(f)
		return
	}

	f.block(f)
	f.iterator(f)
	f.Run()
}

type NewForLoopParams struct {
	value     int
	condition condition
	iterator  part
	block     part
	done      part
}

func NewForLoop(params NewForLoopParams) ForLoop {
	return ForLoop(params)
}

func main() {
	forLoop := NewForLoop(NewForLoopParams{
		value: 0,
		condition: func(f *ForLoop) bool {
			return f.value < 10
		},
		iterator: func(f *ForLoop) {
			f.value++
		},
		block: func(f *ForLoop) {
			fmt.Println(f.value)
		},
		done: func(f *ForLoop) {
			fmt.Println("done!")
		},
	})

	forLoop.Run()
}
