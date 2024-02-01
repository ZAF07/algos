package depthfirstsearch

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// var rll = &Node{
// 	Val:   1,
// 	Left:  nil,
// 	Right: nil,
// }

// var rlr = &Node{
// 	Val:   6,
// 	Left:  nil,
// 	Right: nil,
// }

// var rl = &Node{
// 	Val:   2,
// 	Left:  rll,
// 	Right: rlr,
// }

// var rr = &Node{
// 	Val:   4,
// 	Left:  rl,
// 	Right: nil,
// }

// var rootNode = &Node{
// 	Val:   3,
// 	Left:  rl,
// 	Right: rr,
// }

func TestDFS(t *testing.T) {
	tests := []struct {
		leftRange  int
		rightRange int
		root       *Node
		want       int
	}{
		{
			leftRange:  1,
			rightRange: 5,
			root: &Node{ // Root
				Val: 4,
				Left: &Node{
					Val: 3,
					Left: &Node{
						Val:  2,
						Left: nil,
						Right: &Node{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &Node{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
			},
			want: 9,
		},
		{
			leftRange:  4,
			rightRange: 10,
			root: &Node{ // Root
				Val: 4,
				Left: &Node{
					Val: 3,
					Left: &Node{
						Val:  2,
						Left: nil,
						Right: &Node{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &Node{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &Node{
					Val: 6,
					Left: &Node{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: 37,
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("Given: %v, || Want: %v", tt.root, tt.want)
		t.Run(name, func(t *testing.T) {
			got := bfs(tt.root, tt.leftRange, tt.rightRange)
			assert.Equal(t, tt.want, got)
		})
	}

}

func bfs(n *Node, l, r int) int {
	var sum int

	// Check current node's val, add to sun if falls in range
	if n.Val >= l && n.Val <= r {
		sum += n.Val
	}

	//  if there is a Left node, recursively call the bfs func, simply repeating the above steps
	if n.Left != nil {
		sum += bfs(n.Left, l, r)
	}

	//  if there is a Right node, recursively call the bfs func, simply repeating the above steps
	if n.Right != nil {
		sum += bfs(n.Right, l, r)
	}

	return sum
}
