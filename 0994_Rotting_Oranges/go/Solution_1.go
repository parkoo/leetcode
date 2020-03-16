package main

// 腐烂过程模拟法：构造方法 'rotting()' 模拟一次腐烂过程.
//              若腐烂过程后新鲜橘子数小于腐烂过程前新鲜橘子数，则分钟数'times'加一并继续腐烂过程.
//              直至某次腐烂过程前后，新鲜橘子数不变.
//              此时,若新鲜橘子数为0，则全部橘子被腐烂，返回分钟数'time',否则返回＇-1＇    
// 时间复杂度：O(kmn)  空间复杂度：O(nm)  其中，k为腐烂过程所需的分钟数

var M int  // 行数
var N int  // 列数
var dir = [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}  // 方向
var flag [][]int  // 防止在一次腐烂过程中出现多次传递，即在本次腐烂过程中被腐烂的橘子不能去腐烂它四周的橘子
                  // flag[i][j]=1 表示本次腐烂过程中,（i,j)处的橘子是被腐烂的，不可腐烂其四周的橘子

func orangesRotting_1(grid [][]int) int {
	M = len(grid)
	N = len(grid[0])
	flag = make([][]int, M)
	for i:=0; i<M; i++ {
		flag[i] = make([]int, N)
	}

	times := 0  
	fresh_oranges_num := getNumOfFreshOranges(grid)
	rotting(grid)
	next_fresh_oranges_num := getNumOfFreshOranges(grid)

	for next_fresh_oranges_num<fresh_oranges_num {
		times++
		fresh_oranges_num = next_fresh_oranges_num
		rotting(grid)
		next_fresh_oranges_num = getNumOfFreshOranges(grid)
	}
	
	if getNumOfFreshOranges(grid)>0 {
		return -1
	}

	return times
}

func getNumOfFreshOranges(grid [][]int) int {
	cnt := 0
	for i:=0; i<M; i++ {
		for j:=0; j<N; j++ {
			if grid[i][j]==1 {
				cnt++
			}
		}
	}

	return cnt
}

// 模拟一次腐烂过程，注意在本次腐烂过程中被腐烂的橘子不可向其四周传递腐烂
func rotting(grid [][]int) {
	for i:=0; i<M; i++ {
		for j:=0; j<N; j++ {
			if grid[i][j]==2 && flag[i][j]==0 {  // flag[i][j]=0时，才可向四周腐烂
				for k:=0; k<4; k++ {
					next_i := i + dir[k][0]
					next_j := j + dir[k][1]

					if inArea(next_i, next_j) && grid[next_i][next_j]==1 {
						grid[next_i][next_j] = 2
						flag[next_i][next_j] = 1  // 标记在本次腐烂过程中，该橘子是被腐烂的状态
					}
				}
			}
		}
	}

	// 重置flag,防止影响下一次腐烂过程
	for i:=0; i<M; i++ {
		for j:=0; j<N; j++ {
			flag[i][j] = 0
		}
	}
}

func inArea(i int, j int) bool {
	return i>=0 && i<M && j>=0 && j<N;
}