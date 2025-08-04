function numEnclaves(grid: number[][]): number {
    const m = grid.length;
    const n = grid[0].length;
    let sum = 0
    function dfs(r: number, c: number): ({inBounds: boolean, tiles: number}) {
        if (r < 0 || r >= m || c < 0 || c >= n) {
            // boundary, therefore must set island to invalid
            return {inBounds: false, tiles: 0}
        }
        if (grid[r][c] === 0) {
            // we're ocean, don't count
            return {inBounds: true, tiles:0}
        }
        if (grid[r][c] === 2) {
            // we've visited and counted
            return {inBounds: true, tiles: 0}
        }
        grid[r][c] = 2
        let retTiles = 1

        const adjArr = [dfs(r -1, c), dfs(r + 1, c), dfs(r, c -1), dfs(r, c + 1)]
        for (const a of adjArr) {
            if (!a.inBounds) {
                return {inBounds: false, tiles: 0}
            }
            retTiles += a.tiles
        }
        return {inBounds: true, tiles: retTiles}
    }
    for (let i = 0 ; i < m; i ++) {
        for (let j = 0 ; j < n; j++) {
            if (grid[i][j] === 1) {
                const {inBounds, tiles} = dfs(i,j)
                if (inBounds) {
                    sum += tiles
                }
            }
        }
    }
    return sum
};