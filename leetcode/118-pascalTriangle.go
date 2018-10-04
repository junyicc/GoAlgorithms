package leetcode

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	tri := [][]int{}
	for i := 1; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for j := 1; j < i-1; j++ {
			row[j] = tri[i-2][j] + tri[i-2][j-1]
		}
		tri = append(tri, row)
	}
	return tri
}
