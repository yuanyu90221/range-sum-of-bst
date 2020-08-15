# range sum of  bst

## 題目解讀：

### 題目來源:
[range-sum-of-bst](https://leetcode.com/problems/range-sum-of-bst/)

### 原文:

Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.
### 解讀：

給定一個二元搜尋樹root, 給兩個正整數 L跟R

求出介於L跟R之間的所有元素的和(包含L,R)

## 初步解法:
### 初步觀察:

這題的關鍵在於

要透過二元搜尋樹去traverse介於L跟R的值

所以需要理解二元搜尋樹的分佈規則：

遇到一個數值 會先放在節點

下一個遇到的數值若是比節點大 會放在節點右邊子節點

反之, 則會放在左節點
 
初步想法是 先紀錄 找尋到L,R的所有節點 到一個Set S

然後從L往下找可能大於的L的右subtree

從R往下找左subtree

對Set 檢查元素 把符合條件得做相加

### 初步設計:


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
package range_sum_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	result := 0
	result = transversal(root, L, R)
	return result
}

func transversal(root *TreeNode, L int, R int) int {
	// put unique val
	count := 0
	if root.Val <= R && root.Val >= L {
		count += root.Val
	}
	if root.Left != nil && root.Val != L {
		count += transversal(root.Left, L, R)
	}
	if root.Right != nil && root.Val != R {
		count += transversal(root.Right, L, R)
	}
	return count
}

```
參考別人的作法
```golang
package range_sum_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	result := 0
	result = transversal(root, L, R)
	return result
}

func transversal(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	switch {
	case root.Val > R:
		return transversal(root.Left, L, R)
	case root.Val < L:
		return transversal(root.Right, L, R)
	default:
		return transversal(root.Left, L, R) + root.Val + transversal(root.Right, L, R)
	}
}

```
## 測資的撰寫
```golang
package range_sum_bst

import "testing"

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		L    int
		R    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:  15,
						Left: nil,
						Right: &TreeNode{
							Val:   18,
							Left:  nil,
							Right: nil,
						},
					},
				},
				L: 7,
				R: 15,
			},
			want: 32,
		},
		{
			name: "Example2",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
						Right: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val:   6,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 15,
						Left: &TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   18,
							Left:  nil,
							Right: nil,
						},
					},
				},
				L: 6,
				R: 10,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.L, tt.args.R); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)