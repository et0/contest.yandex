package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const eps = 1e-12

type Model struct {
	index      int
	x, y       float64
	vx, vy     float64
	finish     bool
	finishTime float64
}

type Collision struct {
	time float64
	i, j int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	var L, W float64
	fmt.Fscan(in, &N, &L, &W)

	models := make([]Model, N)
	for i := range N {
		fmt.Fscan(in, &models[i].x, &models[i].y, &models[i].vx, &models[i].vy)
		models[i].index = i + 1
		models[i].finish = false
		models[i].finishTime = math.Inf(1)
	}

	active := make([]bool, N)
	for i := range N {
		active[i] = true

		if models[i].vx > 0 {
			tFinish := (L - models[i].x) / models[i].vx
			if tFinish >= -eps {
				collidesWithWall := false

				if models[i].vy != 0 {
					tBottom := -models[i].y / models[i].vy
					if tBottom > -eps && tBottom < tFinish+eps {
						collidesWithWall = true
					}
				}

				if !collidesWithWall && models[i].vy != 0 {
					tTop := (W - models[i].y) / models[i].vy
					if tTop > -eps && tTop < tFinish+eps {
						collidesWithWall = true
					}
				}

				if !collidesWithWall {
					yFinal := models[i].y + models[i].vy*tFinish
					if yFinal > eps && yFinal < W-eps {
						models[i].finish = true
						models[i].finishTime = tFinish
					}
				}
			}
		}
	}

	var collisions []Collision
	for i := range N {
		if !active[i] {
			continue
		}
		for j := i + 1; j < N; j++ {
			if !active[j] {
				continue
			}
			if t := getCollisionTime(&models[i], &models[j]); t >= -eps {
				maxValidTime := math.Inf(1)
				if models[i].finish {
					maxValidTime = math.Min(maxValidTime, models[i].finishTime)
				}
				if models[j].finish {
					maxValidTime = math.Min(maxValidTime, models[j].finishTime)
				}

				if t < maxValidTime+eps {
					collisions = append(collisions, Collision{t, i, j})
				}
			}
		}
	}

	sort.Slice(collisions, func(i, j int) bool {
		return collisions[i].time < collisions[j].time
	})

	i := 0
	for i < len(collisions) {
		currentTime := collisions[i].time

		collisionGroups := make(map[int]bool)
		j := i
		for j < len(collisions) && math.Abs(collisions[j].time-currentTime) < eps {
			if active[collisions[j].i] && active[collisions[j].j] {
				collisionGroups[collisions[j].i] = true
				collisionGroups[collisions[j].j] = true
			}
			j++
		}

		for modelIdx := range collisionGroups {
			active[modelIdx] = false
			models[modelIdx].finish = false
		}

		i = j
	}

	minTime := math.Inf(1)
	winners := []int{}

	for i := range N {
		if !active[i] || !models[i].finish {
			continue
		}

		if models[i].finishTime < minTime-eps {
			minTime = models[i].finishTime
			winners = []int{models[i].index}
		} else if math.Abs(models[i].finishTime-minTime) < eps {
			winners = append(winners, models[i].index)
		}
	}

	sort.Ints(winners)

	fmt.Fprintln(out, len(winners))
	for _, winner := range winners {
		fmt.Fprint(out, winner, " ")
	}
}

func getCollisionTime(a, b *Model) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dvx := a.vx - b.vx
	dvy := a.vy - b.vy

	if math.Abs(dvx) > eps {
		t := -dx / dvx
		if t < -eps {
			return -1
		}

		y1 := a.y + a.vy*t
		y2 := b.y + b.vy*t
		if math.Abs(y1-y2) < eps {
			return t
		}
	} else if math.Abs(dvy) > eps {
		t := -dy / dvy
		if t < -eps {
			return -1
		}

		x1 := a.x + a.vx*t
		x2 := b.x + b.vx*t
		if math.Abs(x1-x2) < eps {
			return t
		}
	} else {
		if math.Abs(dx) < eps && math.Abs(dy) < eps {
			return 0
		}
	}

	return -1
}
