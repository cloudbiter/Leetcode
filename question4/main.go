package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3, 5}
	nums2 := []int{2, 4}
	modian := findMedianSortedArrays(nums1, nums2)
	fmt.Println(modian)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	findIndex := (len(nums1) + len(nums2) - 1) / 2
	// 使用数组的指针，主要是提升性能，不必值拷贝
	med := dividedAndConquer(&nums1, &nums2, (m+n)%2, findIndex)
	return med
}

func dividedAndConquer(nums1 *[]int, nums2 *[]int, odd, findIndex int) float64 {
	//为了简化判断 ，假设nums1 始终是两数据中最短的
	if len(*nums1) > len(*nums2) {
		nums1, nums2 = nums2, nums1
	}
	// 停止条件 1 ： 当其中一个数组无值时，配合前述nums1长度最短，只需要判断 nums2即可
	if len(*nums1) == 0 {
		if odd == 1 {
			return float64((*nums2)[findIndex])
		} else {
			return float64(((*nums2)[findIndex] + (*nums2)[findIndex+1])) / 2
		}
	} else if findIndex == 0 {
		// 停止条件 2 ： 当不需要再剔除数据时
		if odd == 1 {
			if (*nums1)[0] >= (*nums2)[0] {
				return float64((*nums2)[0])
			} else {
				return float64((*nums1)[0])
			}
		} else {
			if len(*nums1) == 1 && len(*nums2) == 1 {
				return float64((*nums1)[0]+(*nums2)[0]) / 2
			} else if len(*nums2) > 1 && len(*nums1) == 1 {
				return min2Sum((*nums2)[0], (*nums2)[1], (*nums1)[0], 1000001) / 2
			} else {
				return min2Sum((*nums2)[0], (*nums2)[1], (*nums1)[0], (*nums1)[1]) / 2
			}
		}
	}
	// 递归条件， 为了防止数据越界，需要取不超过最短数组长度的剔除序列
	halfIndex := min(int((findIndex+1)/2), len(*nums1))
	if (*nums1)[halfIndex-1] <= (*nums2)[halfIndex-1] {
		tmp := (*nums1)[halfIndex:]
		return dividedAndConquer(&tmp, nums2, odd, findIndex-halfIndex)
	} else {
		tmp := (*nums2)[halfIndex:]
		return dividedAndConquer(nums1, &tmp, odd, findIndex-halfIndex)
	}
}

//当 m + n 为偶数时，需要取最小的两个数，此四个参数的函数可能复用
func min2Sum(a, b, c, d int) float64 {
	if a <= d && c <= b {
		return float64(a + c)
	} else if a > d {
		return float64(c + d)
	}
	return float64(a + b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
