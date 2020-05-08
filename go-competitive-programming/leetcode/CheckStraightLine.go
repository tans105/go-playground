package main

import "fmt"
/*
https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3323/

Check If It Is a Straight Line
You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.

Example 1:

Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true
Example 2:

Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false

Constraints:

2 <= coordinates.length <= 1000
coordinates[i].length == 2
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
coordinates contains no duplicate point.
*/
func main() {
	arr := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}
	fmt.Println(checkStraightLine(arr))

}

func checkStraightLine(coordinates [][]int) bool {
	var prevSlope float64;
	var slope float64;

	for i := 1; i < len(coordinates); i++ {
		y := coordinates[i][1] - coordinates[i - 1][1];
		x := coordinates[i][0] - coordinates[i - 1][0];
		slope = float64(y / x);
		
		if i != 1 && slope != prevSlope {
			return false;
		}
		prevSlope = slope;
	}

	return true;
}
