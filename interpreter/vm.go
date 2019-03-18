package interpreter

import (
	"fmt"
	"bufio"
	"os"
)

// Execute runs a brainfuck program.
func Execute(input string) {
	code := tokenize(input)
	pager := newPager()

	// Index jmp targets
	stack := newStack()
	jmpMap := make(map[int]int)
	for i, ch := range code {
		if ch == '[' {
			stack.push(i)
		}
		if ch == ']' {
			loc, err := stack.pop()
			if err != nil {
				panic("Invalid program: Unmatched brackets.")
			}
			// From opening to closing and from closing to opening
			jmpMap[loc] = i
			jmpMap[i] = loc
		}
	}

	// Memory pointer and instruction pointer
	mp, ip := 0, 0
	for {
		switch code[ip] {
		case '<':
			mp--
		case '>':
			mp++
		case '.':
			// Only ASCII is supported
			fmt.Print(string(pager.getValue(mp)))
		case ',':
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter character: ")
			text, _ := reader.ReadString('\n')
			pager.setValue(mp, text[0])
		case '[':
			if pager.getValue(mp) == 0 {
				ip = jmpMap[ip]
			}
		case ']':
			if pager.getValue(mp) != 0 {
				ip = jmpMap[ip]
			}
		case '+':
			pager.increment(mp)
		case '-':
			pager.decrement(mp)
		}

		if ip < len(code)-1 {
			ip++
		} else {
			// Program finished
			break
		}
	}
}
