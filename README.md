# remade-for-loop-in-go

remade for loops because I was bored, 'for' keyword is not used, it is recursive.

Example:
```go
func main() {
	forLoop := NewForLoop(NewForLoopParams{
		// the initial value of the for loop
		value: 0,

    		// the condition that will determine if the loop runs a block or exits
		condition: func(f *ForLoop) bool {
			return f.value < 10
		},

    		// the iterator ran each time the loop is ran
		iterator: func(f *ForLoop) {
			f.value++
		},

		// the code block that will be executed in the for loop
		block: func(f *ForLoop) {
			fmt.Println(f.value)
		},

    		// a code block that will run after the loop is complete
		done: func(f *ForLoop) {
			fmt.Println("done")
		},
	})

	forLoop.Run()

  // the code above is the same as:
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  fmt.Println("done")
}
```
