package main 

//　暴力法
//  时间复杂度：O(sqr(G))  空间复杂度：O(1)  其中，G为糖果数目，s<=sqr(2G)

//  如：n=4时
//  第一轮：01 02 03 04  ->  0*n*n+(1+n)*n/2=10
//  第二轮：05 06 07 08  ->  1*n*n+(1+n)*n/2=16
//  第三轮：09 10 11 12  ->  2*n*n+(1+n)*n/2=42

func distributeCandies_1(candies int, num_people int) []int {
	
	people := make([]int, num_people)

	i := 0
	for candies>0 {
		for j:=1; j<=num_people; j++ {
			if candies>=j+i*num_people {
				people[j-1] = people[j-1] + j + i*num_people
				candies -= j+i*num_people
			}else {
				people[j-1] += candies
				candies -= candies
				break
			}		
		}
		i++
	} 

	return people
}