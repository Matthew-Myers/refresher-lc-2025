function closedIsland(grid: number[][]): number {
    const m = grid.length;
    const n = grid[0].length;
    let closedIslands = 0;
    
    // DFS to check if an island is closed and mark visited cells
    function dfs(row: number, col: number): boolean {
        // Out of bounds - not a valid cell
        if (row < 0 || row >= m || col < 0 || col >= n) {
            return false;
        }
        
        // If we hit water (1), this direction is closed
        if (grid[row][col] === 1) {
            return true;
        }
        
        // If already visited (2), skip
        if (grid[row][col] === 2) {
            return true;
        }
        
        // Mark as visited
        grid[row][col] = 2;
        
        // Check all 4 directions - ALL must be closed for the island to be closed
        const up = dfs(row - 1, col);
        const down = dfs(row + 1, col);
        const left = dfs(row, col - 1);
        const right = dfs(row, col + 1);
        
        return up && down && left && right;
    }
    
    // Check all cells
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            // If we find an unvisited land cell (0)
            if (grid[i][j] === 0) {
                if (dfs(i, j)) {
                    closedIslands++;
                }
            }
        }
    }
    
    return closedIslands;
}