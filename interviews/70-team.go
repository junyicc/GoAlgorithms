package interviews

/*
10,10
0,0,0,0,0,0,0,0,0,0
0,0,0,1,1,0,1,0,0,0
0,1,0,0,0,0,0,1,0,0
1,0,0,0,0,0,0,0,1,1
0,0,0,1,1,1,0,0,0,1
0,0,0,0,0,0,1,0,1,1
0,1,1,0,0,0,0,0,0,0
0,0,0,1,0,1,0,0,0,0
0,0,1,0,0,1,0,0,0,0
0,1,0,0,0,0,0,0,0,0

output:
6,7

Teamates of (i,j) have the same status (1) in eight directions around (i,j)
*/

// GetTeam returns the count of teams and the members of the max team
func GetTeam(nums [][]int, m, n int) (int, int) {
	team := make([][]int, m)
	for i := 0; i < m; i++ {
		team[i] = make([]int, n)
	}
	teamCnt := 0
	maxTeamMem := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mem := checkTeam(nums, team, i, j, m, n)
			if mem > 0 {
				teamCnt++
				if mem > maxTeamMem {
					maxTeamMem = mem
				}
			}
		}
	}
	return teamCnt, maxTeamMem
}

func checkTeam(nums, team [][]int, i, j, m, n int) int {
	if nums[i][j] == 0 {
		return 0
	}
	teamMem := 0
	if team[i][j] == 0 {
		teamMem = 1
		team[i][j] = 1
	}

	for y := i - 1; y <= i+1; y++ {
		for x := j - 1; x <= j+1; x++ {
			if y == i && x == j {
				continue
			}
			if x >= 0 && x < n && y >= 0 && y < m {
				if team[y][x] == 0 && nums[y][x] == 1 {
					team[y][x] = 1
					teamMem++
					mem := checkTeam(nums, team, y, x, m, n)
					teamMem += mem
				}
			}
		}
	}
	return teamMem
}
