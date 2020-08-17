# count-good-triplets

## 題目解讀：

### 題目來源:
[count-good-triplets](https://leetcode.com/problems/count-good-triplets/)

### 原文:
Given an array of integers arr, and three integers a, b and c. You need to find the number of good triplets.

A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are true:

0 <= i < j < k < arr.length
|arr[i] - arr[j]| <= a
|arr[j] - arr[k]| <= b
|arr[i] - arr[k]| <= c
Where |x| denotes the absolute value of x.

Return the number of good triplets.


### 解讀：
給定一個正整數陣列 arr, 還有三個正整數a, b, c

找尋數對 (arr[i], arr[j], arr[k])符合以下條件

1. 0 <= i < j < k < arr.lenth

2. |arr[i] - arr[j]| <= a
   |arr[j] - arr[k]| <= b
   |arr[i] - arr[k]| <= c
   
找尋所有符合條件的數對個數

## 初步解法:
### 初步觀察:
首先 
對於數對 (arr[i], arr[j], arr[k])

 0 <= i < j < k < arr.length
 
 因此只要固定 i 剩下兩個值 就可以往比i大的直去縮小範圍比對
 
 每次 先固定i 從 0 到 arr.length - 3
 
 j 則會從 i+1 到 arr.length -2
 
 k 則是從 j+1 到 arr.length -1

然後不斷去檢驗條件

### 初步設計:
Given an integer array arr , 3 integer a,b, c

Step 0: let i=0 , count = 0

Step 1: if i > arr.length - 3 go to step 10

Step 2: let j= i + 1

Step 3: if j > arr.length - 2 go to step 1

step 4: let k = j + 1

step 5: if k > arr.length - 1 go to step 3

step 6: check |arr[i] - arr[j]| <= a
              |arr[j] - arr[k]| <= b
              |arr[i] - arr[k]| <= c satisfy then set count += 1
step 7: set k = k + 1 go to step 5

step 8: set j = j + 1 go to step 3

step 9: set i = i + 1 go to step 1

step 10: return count

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package count_good_triplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	lengthOfArr := len(arr)
	for i := 0; i <= lengthOfArr-3; i++ {
		for j := i + 1; j <= lengthOfArr-2; j++ {
			for k := j + 1; k <= lengthOfArr-1; k++ {
				if checkRange(arr[i]-arr[j], arr[j]-arr[k], arr[i]-arr[k], a, b, c) {
					count += 1
				}
			}
		}
	}

	return count
}

func checkRange(d1 int, d2 int, d3 int, a int, b int, c int) bool {
	range1 := abs(d1)
	range2 := abs(d2)
	range3 := abs(d3)
	return range1 <= a && range2 <= b && range3 <= c
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

```
## 測資的撰寫
```golang
package count_good_triplets

import "testing"

func Test_countGoodTriplets(t *testing.T) {
	type args struct {
		arr []int
		a   int
		b   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				arr: []int{3, 0, 1, 1, 9, 7},
				a:   7,
				b:   2,
				c:   3,
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				arr: []int{1, 1, 2, 2, 3},
				a:   0,
				b:   0,
				c:   1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodTriplets(tt.args.arr, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("countGoodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang leetcode 30 20thday](https://hackmd.io/PSGrOqwESUetOjU4ZAUHRA?view)

## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)