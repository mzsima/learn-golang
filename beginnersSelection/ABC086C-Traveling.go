package main
import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

type Plan struct {
	t int
	vertex Vertex
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	plans := []Plan{ Plan{0 , Vertex{0, 0}}}
	for i := 0; i < n; i++ {
		var t, x, y int
		fmt.Scanf("%d%d%d", &t, &x, &y)
		plans = append(plans, Plan{t, Vertex{x, y} } )
	}

    for i := 0; i < len(plans) - 1; i++ {
		cur := plans[i]
		next := plans[i + 1]
		if !check(cur.vertex, next.vertex, next.t - cur.t) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func check(cur Vertex, next Vertex, step int) bool {
	minstep := Abs(next.x - cur.x) + Abs(next.y - cur.y)
	if minstep > step {
		return false
	} else if minstep % 2 == step % 2 {
		return true
	}
	return false
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}