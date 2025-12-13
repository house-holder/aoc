package main

import (
	"fmt"
	"log"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

var debug = true

type Machine struct {
	voltages []uint16
	buttons  []uint16
	target   uint16
	state    uint16
}

func NewMachine(line string) Machine {
	var tgtState uint16 = 0
	accB, accV := []uint16{}, []uint16{}

	for cmp := range strings.SplitSeq(line, " ") {
		switch leader := cmp[0]; leader {
		case '[':
			tCmp := strings.Trim(cmp, "[]")
			for i, c := range tCmp {
				if c == '#' {
					tgtState |= 1 << i
				}
			}

		case '(':
			var newB uint16 = 0
			tCmp := strings.Trim(cmp, "()")
			for c := range strings.SplitSeq(tCmp, ",") {
				n := parseU16(c, "button parse")
				newB |= 1 << n
			}
			accB = append(accB, newB)

		case '{':
			tCmp := strings.Trim(cmp, "{}")
			for c := range strings.SplitSeq(tCmp, ",") {
				newV := parseU16(c, "voltage parse")
				accV = append(accV, newV)
			}
		}
	}
	return Machine{
		voltages: accV,
		buttons:  accB,
		target:   tgtState,
		state:    0,
	}
}

func (m *Machine) operate() (minOps int, cycles int) {
	k := len(m.buttons)
	limit := 1 << k

	minOps = math.MaxInt
	cycles = 0

	for subset := 1; subset < limit; subset++ {
		cycles++
		ops := bits.OnesCount(uint(subset))
		if ops >= minOps {
			continue
		}
		m.state = 0
		for i := range k {
			if subset&(1<<i) != 0 {
				m.state ^= m.buttons[i]
			}
		}
		if m.state == m.target {
			minOps = ops
		}
	}
	return minOps, cycles
}

func Dbg(format string, a ...any) {
	if debug {
		fmt.Printf(format, a...)
	}
}

func parseU16(c string, msg string) uint16 {
	num, err := strconv.Atoi(c)
	if err != nil {
		log.Fatalf("%s failed. %v", msg, err)
	}
	return uint16(num)
}

func evalPart1(lines []string) int {
	operations := 0
	machines := []Machine{}

	for _, line := range lines {
		if line != "" {
			newM := NewMachine(line)
			machines = append(machines, newM)
		}
	}

	cycles := 0
	for _, machine := range machines {
		ops, cyc := machine.operate()
		cycles += cyc
		operations += ops
	}

	Dbg("Total cycles: %d\n", cycles)
	return operations
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	lines := strings.Split(body, "\n")

	result1 := evalPart1(lines)
	// result2 := evalPart2(lines)
	fmt.Println("Part 1:", result1)
	// fmt.Println("Part 2:", result2)
}
