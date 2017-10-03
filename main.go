package main

import (
	"engo.io/engo/math"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	white = " " //"□"
	black = "x" //"■"
	wt    = "△"
	wr    = "▷"
	wb    = "▽"
	wl    = "◁"
	bt    = "▲"
	br    = "▶"
	bb    = "▼"
	bl    = "◀"
)

type dir int

const (
	t dir = iota
	r
	b
	l
)

type CellMap struct {
	w, h        int
	idx, nx, ny int
	d           dir
	cnt         string
	cells       []string
}

func NewMap(w, h int) *CellMap {
	//rs := rand.NewSource(time.Now().Unix())
	//r := rand.New(rs)
	nbCells := w * h
	m := math.Trunc(float32(nbCells/2+w/2)) - 1
	cells := make([]string, nbCells)
	for k, _ := range cells {
		if float32(k) == m {
			cells[k] = wt
		} else {
			cells[k] = white
			/*if r.Intn(15) == 0 {
				cells[k] = black
			} else {
				cells[k] = white
			}*/
		}
	}
	return &CellMap{w: w, h: h, cells: cells}
}

func (m *CellMap) String() string {
	str := ""
	nbCells := len(m.cells)
	for i := 0; i < nbCells; i += 1 {
		if i%m.w == 0 {
			str += "\n"
		}
		str += m.cells[i]
	}
	return str
}

func (m *CellMap) redraw() {
	if m.cnt == white {
		m.cells[m.idx] = black
	} else {
		m.cells[m.idx] = white
	}
	nidx := m.ny*m.w + m.nx
	switch m.d {
	case t:
		if m.cells[nidx] == white {
			m.cells[nidx] = wt
		} else {
			m.cells[nidx] = bt
		}
		break
	case r:
		if m.cells[nidx] == white {
			m.cells[nidx] = wr
		} else {
			m.cells[nidx] = br
		}
		break
	case b:
		if m.cells[nidx] == white {
			m.cells[nidx] = wb
		} else {
			m.cells[nidx] = bb
		}
		break
	case l:
		if m.cells[nidx] == white {
			m.cells[nidx] = wl
		} else {
			m.cells[nidx] = bl
		}
		break
	}
}

func (m *CellMap) move() bool {
	switch m.d {
	case t:
		if m.ny-1 == -1 {
			return false
		}
		m.d = t
		m.ny -= 1
		break
	case r:
		if m.nx+1 == m.w {
			return false
		}
		m.d = r
		m.nx += 1
		break
	case b:
		if m.ny+1 == m.h {
			return false
		}
		m.d = b
		m.ny += 1
		break
	case l:
		if m.nx-1 == -1 {
			return false
		}
		m.d = l
		m.nx -= 1
		break
	}
	m.redraw()
	return true
}

func (m *CellMap) ComputeNext() bool {
	m.idx = -1
	for y := 0; y < m.h; y += 1 {
		for x := 0; x < m.w; x += 1 {
			m.nx = x
			m.ny = y
			m.idx += 1
			c := m.cells[m.idx]
			if c == white || c == black {
				continue
			}
			switch c {
			case wt:
				m.d = r
				m.cnt = white
				return m.move()
			case wr:
				m.d = b
				m.cnt = white
				return m.move()
			case wb:
				m.d = l
				m.cnt = white
				return m.move()
			case wl:
				m.d = t
				m.cnt = white
				return m.move()
			case bt:
				m.d = l
				m.cnt = black
				return m.move()
			case br:
				m.d = t
				m.cnt = black
				return m.move()
			case bb:
				m.d = r
				m.cnt = black
				return m.move()
			case bl:
				m.d = b
				m.cnt = black
				return m.move()
			}
		}
	}
	return true
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	frame := 0
	m := NewMap(200, 70)
	for {
		frame += 1
		clear()
		fmt.Print(frame)
		fmt.Println(m.String())
		if !m.ComputeNext() {
			break
		}
		time.Sleep(150 * time.Millisecond)
	}
}
