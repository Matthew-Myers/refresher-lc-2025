package largestcolorvalue

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	// Build adjacency list
	adj := make([][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
	}

	// visited states: 0 = unvisited, 1 = in current path, 2 = completely visited
	visited := make([]int, n)
	// dp[i][c] represents max count of color c in paths starting from node i
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26) // 26 lowercase letters
	}

	maxCount := 1 // Initialize to 1 since any single node is a valid path
	hasCycle := false

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if visited[node] == 1 {
			// Found a cycle
			hasCycle = true
			return false
		}
		if visited[node] == 2 {
			return true
		}

		visited[node] = 1 // Mark as being visited in current path

		// Get current node's color index
		color := colors[node] - 'a'

		// Initialize current node's color count
		dp[node][color] = 1

		// If node has no outgoing edges, we still need to count its color
		if len(adj[node]) == 0 {
			maxCount = max(maxCount, dp[node][color])
			visited[node] = 2
			return true
		}

		// Visit all neighbors
		for _, next := range adj[node] {
			if !dfs(next) {
				return false
			}
			// Update color counts for all colors
			for c := 0; c < 26; c++ {
				count := dp[next][c]
				if c == int(color) {
					count++ // Add current node's color
				}
				if count > dp[node][c] {
					dp[node][c] = count
				}
				if count > maxCount {
					maxCount = count
				}
			}
		}

		visited[node] = 2 // Mark as completely visited
		return true
	}

	// Try DFS from each unvisited node
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			if !dfs(i) {
				return -1
			}
		}
	}

	if hasCycle {
		return -1
	}

	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
