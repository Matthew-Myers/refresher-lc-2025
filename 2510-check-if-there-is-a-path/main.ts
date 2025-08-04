function isThereAPath(grid: number[][]): boolean {
    const m = grid.length;
    const n = grid[0].length;
    const max = Math.floor((m + n) / 2);
    
    // If total path length is odd, it's impossible to have equal 0s and 1s
    if ((m + n - 1) % 2 !== 0) return false;
    
    // Create memo map using string key for state
    const memo = new Map<string, boolean>();
    
    function explore(x: number, y: number, diff: number): boolean {
        // Out of bounds or impossible difference
        if (x >= m || y >= n || Math.abs(diff) > max) return false;
        
        // Update difference (using diff instead of separate counts)
        diff += grid[x][y] === 1 ? 1 : -1;
        
        // Reached end - check if counts are equal (diff = 0)
        if (x === m - 1 && y === n - 1) return diff === 0;
        
        // Create unique key for current state
        const key = `${x},${y},${diff}`;
        if (memo.has(key)) return memo.get(key)!;
        
        // Try both directions
        const result = explore(x + 1, y, diff) || explore(x, y + 1, diff);
        memo.set(key, result);
        return result;
    }
    
    return explore(0, 0, 0);
}
