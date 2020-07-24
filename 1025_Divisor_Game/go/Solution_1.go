package main

// 找规律
// 如果N是奇数，因为奇数的所有因数都是奇数，因此 N 进行一次 N-x 的操作结果一定是偶数，
// 所以如果 Alice 拿到了一个奇数，那么轮到 Bob 的时候，Bob 拿到的肯定是偶数，
// 这个时候 Bob 只要进行 -1， 还给 Alice 一个奇数，那么这样 Bob 就会一直拿到偶数，
// 到最后 Bob 一定会拿到最小偶数2，Alice 就输了。
// 所以,如果游戏开始时 Alice 拿到 N 为奇数，那么她必输，也就是false。
// 如果N为偶数，Alice 只用 -1，让 Bob 拿到奇数，则最后 Bob必输，结果就是true。

// 时间复杂度：O(1)  空间复杂度：O(1)

func divisorGame_1(N int) bool {

	return N%2==0
}