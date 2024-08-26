package main

// 思路: 可变数组 + 游标

// 时间复杂度: O(?)    空间复杂度: O(?)

type BrowserHistory struct {
	list []string
	cur  int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		list: []string{homepage},
		cur:  0,
	}
}

// 若当前不是最新页面，这里需要作覆盖
func (this *BrowserHistory) Visit(url string) {
	if this.cur < len(this.list)-1 {
		this.list = this.list[:this.cur+1]
	}

	this.list = append(this.list, url)
	this.cur++
}

func (this *BrowserHistory) Back(steps int) string {
	if steps > this.cur {
		this.cur = 0
	} else {
		this.cur -= steps
	}

	return this.list[this.cur]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.cur+steps >= len(this.list) {
		this.cur = len(this.list) - 1
	} else {
		this.cur += steps
	}

	return this.list[this.cur]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
