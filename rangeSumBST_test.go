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
