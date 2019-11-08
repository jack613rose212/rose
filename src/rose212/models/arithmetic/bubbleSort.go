package arithmetic

import "fmt"

//冒泡排序
func BubleSort() {

	num := []int{0, 9, 999, 564, 4574, 878, 25, 4, 212, 5654}
	for i := 1; i < len(num); i++ { //控制次数  10个数字  控制九次就可以
		for j := 0; j < len(num)-i; j++ { //初始化0切片下标  控制置换次数(置换次数与控制次数相关)
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	fmt.Println("冒泡排序从小到大的结果是:", num)
}
