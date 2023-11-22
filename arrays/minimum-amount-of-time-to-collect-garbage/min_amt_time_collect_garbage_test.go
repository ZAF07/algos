package minimumamountoftimetocollectgarbage

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
2391. Minimum Amount of Time to Collect Garbage
Medium

You are given a 0-indexed array of strings garbage where garbage[i] represents the assortment of garbage at the ith house. garbage[i] consists only of the characters 'M', 'P' and 'G' representing one unit of metal, paper and glass garbage respectively. Picking up one unit of any type of garbage takes 1 minute.
You are also given a 0-indexed integer array travel where travel[i] is the number of minutes needed to go from house i to house i + 1.
There are three garbage trucks in the city, each responsible for picking up one type of garbage. Each garbage truck starts at house 0 and must visit each house in order; however, they do not need to visit every house.
Only one garbage truck may be used at any given moment. While one truck is driving or picking up garbage, the other two trucks cannot do anything.
Return the minimum number of minutes needed to pick up all the garbage.

Example 1:
Input: garbage = ["G","P","GP","GG"], travel = [2,4,3]
Output: 21
Explanation:
The paper garbage truck:
1. Travels from house 0 to house 1
2. Collects the paper garbage at house 1
3. Travels from house 1 to house 2
4. Collects the paper garbage at house 2
Altogether, it takes 8 minutes to pick up all the paper garbage.
The glass garbage truck:
1. Collects the glass garbage at house 0
2. Travels from house 0 to house 1
3. Travels from house 1 to house 2
4. Collects the glass garbage at house 2
5. Travels from house 2 to house 3
6. Collects the glass garbage at house 3
Altogether, it takes 13 minutes to pick up all the glass garbage.
Since there is no metal garbage, we do not need to consider the metal garbage truck.
Therefore, it takes a total of 8 + 13 = 21 minutes to collect all the garbage.

Example 2:
Input: garbage = ["MMM","PGM","GP"], travel = [3,10]
Output: 37
Explanation:
The metal garbage truck takes 7 minutes to pick up all the metal garbage.
The paper garbage truck takes 15 minutes to pick up all the paper garbage.
The glass garbage truck takes 15 minutes to pick up all the glass garbage.
It takes a total of 7 + 15 + 15 = 37 minutes to collect all the garbage.

Constraints:
2 <= garbage.length <= 105
garbage[i] consists of only the letters 'M', 'P', and 'G'.
1 <= garbage[i].length <= 10
travel.length == garbage.length - 1
1 <= travel[i] <= 100
*/

func TestGarbageCollection(t *testing.T) {
	tests := []struct {
		garbage []string
		travel  []int
		want    int
	}{
		{[]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}, 21},
		{[]string{"MMM", "PGM", "GP"}, []int{3, 10}, 37},
		{[]string{"MMMGP", "PGM", "GP"}, []int{3, 10}, 39},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v && %v || want: %v", tt.garbage, tt.travel, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := garbageCollection(tt.garbage, tt.travel)
			want := tt.want
			assert.Equal(t, got, want)
		})
	}
}

func garbageCollection(garbage []string, travel []int) int {
	/*
		We know that it takes x amount of time to travel to a specific house regardless of how much rubbish there is to collect.
		So even if at house 1 (garbage[0]), there is no metal garbage to collect but 1 metal garbage to collect at house 4 (garbage[3]),
		the amount of time to travel still remains at x while the garbage collection time is just 1

		So with this knowledge, we can simply loop the garbage array once, and at each iteration, we calculate the total time it takes to travel from house 0 to current house
		At each house, we simply add the time it takes to collect all garbage (meaning the length of the string representing the garbage type on garbage array)
		Also, loop the garbage types and update the travel time for each garbage type
	*/

	// constant space for storing total garbage collection time for each garbage type at each house
	var totalGarbageCollectionTime int
	var metalTravelTime int
	var paperTravelTime int
	var glassTravelTime int
	var travelTime int

	// Loop the garbage array representing the trucks going house to house
	for i, g := range garbage {

		// If current house is the first house (garbage[0]) we do not need to calculate travel time. As all trucks start at the first house
		if i != 0 {
			travelTime += travel[i-1]
		}

		// We do not need to care about the type of garbage being collected at each house, we can simply add the length of the string (garbage types at each house) to the totalGarbageCollectionTime
		// This is because we are looking for the total time it takes for all trucks to complete garbage collection (not individually)
		totalGarbageCollectionTime += len(g)

		// Loop through the garbage type at current house and update the travel time it takes to get to current house
		for _, t := range g {
			switch string(t) {
			case "M":
				metalTravelTime = travelTime

			case "P":
				paperTravelTime = travelTime
			case "G":
				glassTravelTime = travelTime
			default:
				fmt.Printf("Invalid garbage type: %v @ house: %v", string(t), i)
				continue
			}

		}
	}

	return totalGarbageCollectionTime + metalTravelTime + glassTravelTime + paperTravelTime
}
