package interpreter

import (
	"fmt"
	"bufio"
	"errors"
	"io"
)

// Execute runs a brainfuck program.
func Execute(program string, in io.Reader, out io.Writer) error {
	tokens := Tokenize(program)
	pager := newPager()

	// Index jmp targets
	stack := newStack()
	jmpMap := make(map[int]int)
	for i, ch := range tokens {
		if ch == '[' {
			stack.push(i)
		}
		if ch == ']' {
			loc, err := stack.pop()
			if err != nil {
				return errors.New("Invalid code: unmatched brackets")
			}
			// From opening to closing and from closing to opening
			jmpMap[loc] = i
			jmpMap[i] = loc
		}
	}
	if !stack.isEmpty() {
		return errors.New("Invalid code: unmatched brackets")
	}

	// Memory pointer and instruction pointer
	mp, ip := 0, 0
	for {
		switch tokens[ip] {
		case '<':
			mp--
		case '>':
			mp++
		case '.':
			// Only ASCII is supported
			fmt.Fprint(out, string(pager.getValue(mp)))
		case ',':
			reader := bufio.NewReader(in)
			fmt.Fprint(out, "Enter character: ")
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

		if ip < len(tokens)-1 {
			ip++
		} else {
			// Program finished
			break
		}
	}
	return nil
}
