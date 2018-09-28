package interviews

/*
小Q所在的王国有n个城市，城市之间有m条单向道路连接起来。
对于一个城市V，从城市V出发可以到达的城市数量为X，从某个城市出发可以达到城市V的城市数量为Y，如果Y>X ，则城市V是一个重要城市（间接到达也算可以到达）。
小Q希望你能帮他计算一下王国中一共有多少个重要城市。

输入描述
输入包括m+1行
第一行包括两个数n和m（1<=n,m<=1000）,分别表示城市数和道路数
接下来m行，每行两个数u,v(1<=u,v<=n)，表示一条从u到v的有向道路，输入可能包含重边和自环

input:
4 3
2 1
3 2
4 3

输出描述
输出一个数，重要节点的个数。
output:
2
*/

// GetImportantCitys returns the important city number
func GetImportantCitys(graph [][]int, n, m int) int {
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		dfs(graph, i, i, degree, visited)
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if degree[i] > 0 {
			cnt++
		}
	}
	return cnt
}

func dfs(graph [][]int, origin, from int, degree []int, visited []bool) {
	for i, e := range graph[from] {
		if i != from && e == 1 && !visited[i] {
			degree[origin]--
			degree[i]++
			visited[i] = true
			dfs(graph, origin, i, degree, visited)
		}
	}
}
