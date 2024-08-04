package main

import (
	"fmt"
	"strings"
)

type game struct {
	state map[rune][]int
	moves int
}

func main() {
	game := newGame(6)

	fmt.Printf("Towers of Hanoi! Move the disks on stack A to stack C\n\n")
	fmt.Printf("Make a move. Example: 'AB'\nmoves the top disk from stack A to stack B")
	game.print()
	for {
		line := ""
		fmt.Scanln(&line)
		if len(line) != 2 || line[0] == ' ' || line[1] == ' ' {
			fmt.Println("Invalid move")
			continue
		}
		err := game.move(rune(line[0]), rune(line[1]))
		if err != nil {
			fmt.Print(err)
			continue
		}
		game.moves++
		game.print()

		if game.done() {
			fmt.Printf("\nYou successfully solved the Towers of Hanoi in %d moves! Congratulations!\n", game.moves)
			break
		}
	}
}

func (g game) done() bool {
	if len(g.state['A']) > 0 {
		return false
	}
	if len(g.state['B']) > 0 {
		return false
	}
	return true 
}

func (g *game) move(a, b rune) error {
	if a > 90 {
		a -= 32
	}
	if b > 90 {
		b -= 32
	}
	if _, ok := g.state[a]; !ok {
		return fmt.Errorf("Invalid move %c\n", a) 
	}
	if _, ok := g.state[b]; !ok {
		return fmt.Errorf("Invalid move %c\n", b) 
	}

	l := len(g.state[a])
	if l == 0 {
		return fmt.Errorf("No disk to move on stack %c\n", a) 
	}

	top := g.state[a][l-1]
	if !g.canMove(top, b) {
		return fmt.Errorf("Invalid move: disk on stack %c is too big to move to stack %c\n", a, b)
	}

	g.state[a] = g.state[a][:l-1] 
	g.state[b] = append(g.state[b], top)
	return nil
}

func (g *game) canMove(disk int, dest rune) bool {
	l := len(g.state[dest])
	if l == 0 {
		return true
	}
	top := g.state[dest][l-1]
	if top < disk {
		return false
	}
	return true 
}

func newGame(rods int) game {
	a := make([]int, rods)
	for i := range rods {
		a[i] = rods - i 
	}
	return game{
		state: map[rune][]int{
			'A': a,
			'B': {},
			'C': {},
		},
	}
}

func (g game) print() {
	a, b, c := g.state['A'], g.state['B'], g.state['C']
	m := len(a) + len(b) + len(c)

	fmt.Printf("\n")
	fmt.Printf("%sA%s", strings.Repeat(" ", m), strings.Repeat(" ", m))
	fmt.Printf("%sB%s", strings.Repeat(" ", m), strings.Repeat(" ", m))
	fmt.Printf("%sC%s", strings.Repeat(" ", m), strings.Repeat(" ", m))
	fmt.Printf("\n")
	for i := m-1; i >= 0; i-- {
		printRod(i, m, a)
		printRod(i, m, b)
		printRod(i, m, c)
		fmt.Println()
	}
	fmt.Printf("%s\n", strings.Repeat("#", m*6))
}

func printRod(idx, gap int, r []int) {
	if idx < len(r) {
		fmt.Printf("%s%s|%s%s", 
		strings.Repeat(" ", gap - r[idx]), 
		strings.Repeat("-", r[idx]),
		strings.Repeat("-", r[idx]),
		strings.Repeat(" ", gap - r[idx]))
	} else {
		fmt.Printf("%s|%s", strings.Repeat(" ", gap), strings.Repeat(" ", gap))
	}
}
