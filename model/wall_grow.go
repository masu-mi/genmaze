package model

import (
	"time"

	"math/rand"
)

const maxLength = 30

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateMazeWithWallGrow(x, y int) (m *Maze, err error) {
	m, err = NewMaze(x, y)
	if err != nil {
		return nil, err
	}
	candidates := set{}
	for i := 1; i <= x/2; i++ {
		for j := 1; j <= y/2; j++ {
			candidates[pair{i * 2, j * 2}] = mark
		}
	}

	q := []pair{}
	for k := range candidates {
		q = append(q, k)
	}
	rand.Shuffle(len(q), func(i, j int) {
		q[i], q[j] = q[j], q[i]
	})
	for len(candidates) > 0 {
		n := q[0]
		q = q[1:]
		if _, ok := candidates[n]; !ok {
			continue
		}

		var (
			l    int
			last bool
		)
		adding := set{}
		for l < maxLength && !last {
			m.filed[n.x][n.y] = wall
			delete(candidates, n)
			adding[n] = mark

			next := selectNextPos(n, adding)
			if m.filed[next.x][next.y] == wall {
				last = true
			}
			mid := middle(n, next)
			m.filed[mid.x][mid.y] = wall
			n = next
		}
	}
	return m, nil
}

func selectNextPos(cur pair, adding map[pair]none) (next pair) {
	diffs := []pair{pair{2, 0}, pair{-2, 0}, pair{0, 2}, pair{0, -2}}
	rand.Shuffle(len(diffs), func(i, j int) {
		diffs[i], diffs[j] = diffs[j], diffs[i]
	})
	for _, df := range diffs {
		next = pair{cur.x + df.x, cur.y + df.y}
		if _, ok := adding[next]; !ok {
			return next
		}
	}
	// no path
	return pair{}
}

func middle(s, t pair) pair {
	return pair{
		x: (s.x + t.x) / 2,
		y: (s.y + t.y) / 2,
	}
}

type pair struct {
	x, y int
}

type none struct{}

var mark none

type set map[pair]none
