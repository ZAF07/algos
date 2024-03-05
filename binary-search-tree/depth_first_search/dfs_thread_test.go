package depthfirstsearch

import (
	"fmt"
	"sync"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type NodeThread struct {
	Val         int
	Left, Right *NodeThread
}

var tests = []struct {
	leftRange  int
	rightRange int
	root       *NodeThread
	want       int
}{
	{
		leftRange:  1,
		rightRange: 5,
		root: &NodeThread{ // Root
			Val: 4,
			Left: &NodeThread{
				Val: 3,
				Left: &NodeThread{
					Val:  2,
					Left: nil,
					Right: &NodeThread{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &NodeThread{
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
		root: &NodeThread{ // Root
			Val: 4,
			Left: &NodeThread{
				Val: 3,
				Left: &NodeThread{
					Val:  2,
					Left: nil,
					Right: &NodeThread{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &NodeThread{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &NodeThread{
				Val: 6,
				Left: &NodeThread{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &NodeThread{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		want: 37,
	},
}

func TestDfsThread(t *testing.T) {

	for _, tt := range tests {
		testName := fmt.Sprintf("Given %v || Want: %v", tt.root, tt.want)

		t.Run(testName, func(t *testing.T) {
			var res int

			// init the channel & the waitGroup
			resChan := make(chan int)
			wg := sync.WaitGroup{}

			// Add 1 to the waitgroup and start the recursive call to dfs
			wg.Add(1)
			go dfsThread(tt.root, tt.leftRange, tt.rightRange, &wg, resChan)

			// Place the wg.wait and channel close inside a goroutien so that we dont block the main routine
			go func() {
				wg.Wait()
				close(resChan)
			}()

			// now receive any values from the channel and add to res
			for r := range resChan {
				res += r
			}

			assert.Equal(t, res, tt.want)
		})
	}
}

func BenchmarkDFSThread(b *testing.B) {
	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v || Want: %v", tt.root, tt.want)
		b.Run(testName, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				benchDFSThread(tt.root, tt.leftRange, tt.rightRange)
			}
		})
	}
}

// Benchmark helper
func benchDFSThread(node *NodeThread, l, r int) {
	var res int
	resChan := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go dfsThread(node, l, r, &wg, resChan)

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for r := range resChan {
		res += r
	}

}

// Actual implementation
func dfsThread(NodeThread *NodeThread, l, r int, wg *sync.WaitGroup, resChan chan<- int) {
	defer wg.Done()
	sum := 0
	if NodeThread.Val >= l && NodeThread.Val <= r {
		sum += NodeThread.Val
	}

	if NodeThread.Left != nil {
		wg.Add(1)
		go dfsThread(NodeThread.Left, l, r, wg, resChan)
	}

	if NodeThread.Right != nil {
		wg.Add(1)
		go dfsThread(NodeThread.Right, l, r, wg, resChan)
	}
	resChan <- sum
}
