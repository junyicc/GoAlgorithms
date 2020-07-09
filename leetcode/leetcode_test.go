package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	result := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)

	result = ThreeSum([]int{0, 0, 0})
	fmt.Println(result)
}

func TestThreeSumClosest(t *testing.T) {
	if result := ThreeSumClosest([]int{-1, 2, 1, 4}, 1); result != 2 {
		t.Errorf("expected 2 and got %d", result)
	}
}

func TestFourSum(t *testing.T) {
	result := FourSum([]int{0, 0, 0, 0}, 0)
	fmt.Println(result)
}

func Test45JumpGames(t *testing.T) {
	input := []int{2, 3, 1, 1, 4}
	expected := 2
	if result := Jump(input); result != expected {
		t.Errorf("expected %d and got %d", expected, result)
	}

	input = []int{2}
	expected = 0
	if result := Jump(input); result != expected {
		t.Errorf("expected %d and got %d", expected, result)
	}

	input = []int{0} // false
	expected = 0
	if result := Jump(input); result != expected {
		t.Errorf("expected %d and got %d", expected, result)
	}

	input = []int{} // false
	expected = 0
	if result := Jump(input); result != expected {
		t.Errorf("expected %d and got %d", expected, result)
	}
}

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "jump games",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test2AddNumbers(t *testing.T) {
	l1_3 := ListNode{Val: 3}
	l1_2 := ListNode{Val: 4, Next: &l1_3}
	l1_1 := ListNode{Val: 2, Next: &l1_2}
	l2_3 := ListNode{Val: 4}
	l2_2 := ListNode{Val: 6, Next: &l2_3}
	l2_1 := ListNode{Val: 5, Next: &l2_2}
	l := addTwoNumbers(&l1_1, &l2_1)
	if l.Val != 7 || l.Next.Val != 0 || l.Next.Next.Val != 8 {
		t.Errorf("failed to add two linked list")
	}
}

func Test3LenOfLCS(t *testing.T) {
	if n := lengthOfLongestSubstring("abcabcbb"); n != 3 {
		t.Errorf("expected 3 and got %d\n", n)
	}

	if n := lengthOfLongestSubstring("bbbbbbb"); n != 1 {
		t.Errorf("expected 1 and got %d\n", n)
	}

	if n := lengthOfLongestSubstring("pwwkew"); n != 3 {
		t.Errorf("expected 3 and got %d\n", n)
	}
}

func Test5LongestPalindrome(t *testing.T) {
	if subString := longestPalindrome("abbd"); subString != "bb" {
		t.Errorf("expected bb and got %s\n", subString)
	}

	if subString := longestPalindrome("babad"); subString != "aba" {
		t.Errorf("expected aba and got %s\n", subString)
	}

	if subString := longestPalindrome("ac"); subString != "c" {
		t.Errorf("expected c and got %s\n", subString)
	}
}

func Test7ReverseInt(t *testing.T) {
	if d := reverse(123); d != 321 {
		t.Errorf("expected 321 and got %d\n", d)
	}
	if d := reverse(-123); d != -321 {
		t.Errorf("expected -321 and got %d\n", d)
	}
	if d := reverse(120); d != 21 {
		t.Errorf("expected 21 and got %d\n", d)
	}
	if d := reverse(0); d != 0 {
		t.Errorf("expected 0 and got %d\n", d)
	}
}

func Test8Atoi(t *testing.T) {
	if n := myAtoi("42"); n != 42 {
		t.Errorf("expected 42 and got %d", n)
	}
	if n := myAtoi("    -42"); n != -42 {
		t.Errorf("expected -42 and got %d", n)
	}
	if n := myAtoi("4193 with words"); n != 4193 {
		t.Errorf("expected 4193 and got %d", n)
	}
	if n := myAtoi("words and 987"); n != 0 {
		t.Errorf("expected 0 and got %d", n)
	}
	if n := myAtoi("-91283472332"); n != -2147483648 {
		t.Errorf("expected -2147483648 and got %d", n)
	}
	if n := myAtoi("-91283472332"); n != -2147483648 {
		t.Errorf("expected -2147483648 and got %d", n)
	}
	if n := myAtoi("  0000000001234"); n != 1234 {
		t.Errorf("expected 1234 and got %d", n)
	}
	if n := myAtoi("010"); n != 10 {
		t.Errorf("expected 10 and got %d", n)
	}
}

func Test9PalindromeNum(t *testing.T) {
	if b := isPalindrome(121); !b {
		t.Errorf("expected true and got %t", b)
	}
	if b := isPalindrome(-121); b {
		t.Errorf("expected false and got %t", b)
	}
	if b := isPalindrome(10); b {
		t.Errorf("expected false and got %t", b)
	}
}

func Test12IntToRoman(t *testing.T) {
	if s := intToRoman(3); s != "III" {
		t.Errorf("expected III and got %s", s)
	}
	if s := intToRoman(4); s != "IV" {
		t.Errorf("expected IV and got %s", s)
	}
	if s := intToRoman(9); s != "IX" {
		t.Errorf("expected IX and got %s", s)
	}
	if s := intToRoman(58); s != "LVIII" {
		t.Errorf("expected LVIII and got %s", s)
	}
	if s := intToRoman(1994); s != "MCMXCIV" {
		t.Errorf("expected MCMXCIV and got %s", s)
	}
}
func Test13RomanToInt(t *testing.T) {
	if n := romanToInt("III"); n != 3 {
		t.Errorf("expected 3 and got %d", n)
	}
	if n := romanToInt("IV"); n != 4 {
		t.Errorf("expected 4 and got %d", n)
	}
	if n := romanToInt("IX"); n != 9 {
		t.Errorf("expected 9 and got %d", n)
	}
	if n := romanToInt("LVIII"); n != 58 {
		t.Errorf("expected 58 and got %d", n)
	}
	if n := romanToInt("MCMXCIV"); n != 1994 {
		t.Errorf("expected 1994 and got %d", n)
	}
}

func Test14LongestCommonPrefix(t *testing.T) {
	if s := longestCommonPrefix([]string{"flower", "flow", "flight"}); s != "fl" {
		t.Errorf("expected fl and got %s", s)
	}

	if s := longestCommonPrefix([]string{"dog", "racecar", "car"}); s != "" {
		t.Errorf("expected fl and got %s", s)
	}
}

func Test28StrStr(t *testing.T) {
	if n := strStr("mississippi", "issip"); n != 4 {
		t.Errorf("expected 4 and got %d\n", n)
	}
}

func Test70ClimbStairs(t *testing.T) {
	if n := climbStairs(2); n != 2 {
		t.Errorf("expected 2 and got %d\n", n)
	}

	if n := climbStairs(3); n != 3 {
		t.Errorf("expected 3 and got %d\n", n)
	}
}

func Test121MaxProfit(t *testing.T) {
	if n := maxProfit([]int{7, 1, 5, 3, 6, 4}); n != 5 {
		t.Errorf("expected 5 and got %d\n", n)
	}
	if n := maxProfit([]int{7, 6, 4, 3, 1}); n != 0 {
		t.Errorf("expected 0 and got %d\n", n)
	}
	if n := maxProfit([]int{2, 4, 1}); n != 2 {
		t.Errorf("expected 2 and got %d\n", n)
	}
}

func Test125ValidPalindrome(t *testing.T) {
	if b := isPalindromeString(" "); !b {
		t.Errorf("expected false and got %t\n", b)
	}
	if b := isPalindromeString("A man, a plan, a canal: Panama"); !b {
		t.Errorf("expected true and got %t\n", b)
	}
	if b := isPalindromeString("0P"); b {
		t.Errorf("expected false and got %t\n", b)
	}
}

func Test155MinStack(t *testing.T) {
	obj := StackConstructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	if m := obj.GetMin(); m != -3 {
		t.Errorf("expected -3 and got %d\n", m)
	}
	obj.Pop()
	if top := obj.Top(); top != 0 {
		t.Errorf("expected 0 and got %d\n", top)
	}
	if m := obj.GetMin(); m != -2 {
		t.Errorf("expected -2 and got %d\n", m)
	}
}

func Test166FindFirstCommonNode(t *testing.T) {
	n7 := ListNode{Val: 7, Next: nil}
	n6 := ListNode{Val: 6, Next: &n7}
	n5 := ListNode{Val: 5, Next: &n6}
	h2 := ListNode{Val: 4, Next: &n5}
	n3 := ListNode{Val: 3, Next: &n6}
	n2 := ListNode{Val: 2, Next: &n3}
	h1 := ListNode{Val: 1, Next: &n2}

	if n := getIntersectionNode(&h1, &h2); n.Val != 6 {
		t.Errorf("expected 6 and got %v", n.Val)
	}

	h1 = ListNode{Val: 7, Next: nil}

	if n := getIntersectionNode(&h1, &h1); n.Val != 7 {
		t.Errorf("expected 6 and got %v", n.Val)
	}
}

func Test198RobHouses(t *testing.T) {
	if n := rob3([]int{1, 2, 3, 1}); n != 4 {
		t.Errorf("expected 4 and got %d\n", n)
	}
	if n := rob3([]int{2, 7, 9, 3, 1}); n != 12 {
		t.Errorf("expected 12 and got %d\n", n)
	}
	if n := rob3([]int{6, 6, 4, 8, 4, 3, 3, 10}); n != 27 {
		t.Errorf("expected 27 and got %d\n", n)
	}
	if n := rob3([]int{2, 4, 8, 9, 9, 3}); n != 19 {
		t.Errorf("expected 19 and got %d\n", n)
	}
}

func Test202HappyNum(t *testing.T) {
	if b := isHappy(18); b {
		t.Errorf("expected false and got %t\n", b)
	}
	if b := isHappy(19); !b {
		t.Errorf("expected true and got %t\n", b)
	}
}

func Test326PowerOfThree(t *testing.T) {
	if b := isPowerOfThree(81); !b {
		t.Errorf("expected true and got %t\n", b)
	}
	if b := isPowerOfThree(14348908); b {
		t.Errorf("expected false and got %t\n", b)
	}
}
func Test326PowerOfFour(t *testing.T) {
	if b := isPowerOfFour(16); !b {
		t.Errorf("expected true and got %t\n", b)
	}
	if b := isPowerOfFour(2); b {
		t.Errorf("expected false and got %t\n", b)
	}
	if b := isPowerOfFour(8); b {
		t.Errorf("expected false and got %t\n", b)
	}
	if b := isPowerOfFour(4194304); !b {
		t.Errorf("expected true and got %t\n", b)
	}
}
func Test746ClimbStairs(t *testing.T) {
	if n := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}); n != 6 {
		t.Errorf("expected 6 and got %d\n", n)
	}
}

func Test_mergeSortedArr(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "merge sorted arr",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSortedArr(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}

func Test_majorityElementII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "marjority element",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: []int{3},
		}, {
			name: "marjority element2",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElementII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first missing positive",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: 3,
		}, {
			name: "first missing positive",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		}, {
			name: "first missing positive",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		}, {
			name: "first missing positive",
			args: args{
				nums: []int{1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "max sliding windown",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrtWithBinary(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "my sqrt",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "my sqrt",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "my sqrt",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "my sqrt",
			args: args{
				x: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrtWithBinary(tt.args.x); got != tt.want {
				t.Errorf("mySqrtWithBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverse words",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "reverse words",
			args: args{
				s: "  hello world!  ",
			},
			want: "world! hello",
		},
		{
			name: "reverse words",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
		{
			name: "reverse words",
			args: args{
				s: " ",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "num of island",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 1,
		}, {
			name: "num of island",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			want: 3,
		}, {
			name: "num of island",
			args: args{
				grid: [][]byte{
					{'1', '1', '1'},
					{'0', '1', '0'},
					{'1', '1', '1'},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest continuous increasing subsequence",
			args: args{
				nums: []int{1, 3, 5, 7},
			},
			want: 4,
		}, {
			name: "longest continuous increasing subsequence",
			args: args{
				nums: []int{2, 2, 2, 2},
			},
			want: 1,
		}, {
			name: "longest continuous increasing subsequence",
			args: args{
				nums: []int{3, 0, 6, 2, 4, 7, 0, 0},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lengthOfLIS",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		}, {
			name: "lengthOfLIS",
			args: args{
				nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			want: 6,
		}, {
			name: "lengthOfLIS",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		}, {
			name: "lengthOfLIS",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNumberOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "number of LIS",
			args: args{
				nums: []int{1, 3, 5, 4, 7},
			},
			want: 2,
		}, {
			name: "number of LIS",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
			},
			want: 5,
		}, {
			name: "number of LIS",
			args: args{
				nums: []int{1, 2, 4, 3, 5, 4, 7, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "min path sum",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "coin change",
			args: args{
				coins:  []int{186, 419, 83, 408},
				amount: 6249,
			},
			want: 20,
		}, {
			name: "coin change",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		}, {
			name: "coin change",
			args: args{
				coins:  []int{1, 3, 5},
				amount: 9,
			},
			want: 3,
		}, {
			name: "coin change",
			args: args{
				coins:  []int{2, 5, 10, 1},
				amount: 27,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "max product",
			args: args{
				nums: []int{2, -5},
			},
			want: 2,
		}, {
			name: "max product",
			args: args{
				nums: []int{0, 2},
			},
			want: 2,
		}, {
			name: "max product",
			args: args{
				nums: []int{2, -5, -2, -4, 3},
			},
			want: 24,
		}, {
			name: "max product",
			args: args{
				nums: []int{-2, 0, -4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "min triangle path sum",
			args: args{
				triangle: [][]int{
					{2},
					{3, 4},
					{6, 5, 7},
					{4, 1, 8, 3},
				},
			},
			want: 11,
		}, {
			name: "min triangle path sum",
			args: args{
				triangle: [][]int{
					{-1},
					{2, 3},
					{1, -1, -3},
				},
			},
			want: -1,
		}, {
			name: "min triangle path sum",
			args: args{
				triangle: [][]int{
					{-1},
					{3, 2},
					{-3, 1, -1},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "min edit distance",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		}, {
			name: "min edit distance",
			args: args{
				word1: "horse",
				word2: "r",
			},
			want: 4,
		}, {
			name: "min edit distance",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partitionLinkedList(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "partition linkedlist",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  2,
										Next: nil,
									},
								},
							},
						},
					},
				},
				x: 3,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLinkedList(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid sudo",
			args: args{
				board: [][]byte{
					{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
					{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "eval rpn",
			args: args{
				tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			},
			want: 22,
		}, {
			name: "eval rpn",
			args: args{
				tokens: []string{"2", "1", "+", "3", "*"},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "decode string",
			args: args{
				s: "3[a]12[bc]",
			},
			want: "aaabcbcbcbcbcbcbcbcbcbcbcbc",
		}, {
			name: "decode string",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		}, {
			name: "decode string",
			args: args{
				s: "abc3[cd]xyz",
			},
			want: "abccdcdcdxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "largest react area",
			args: args{
				heights: []int{2},
			},
			want: 2,
		}, {
			name: "largest react area",
			args: args{
				heights: []int{2, 1, 5, 6, 2, 3},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01 matrix",
			args: args{
				matrix: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumberII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single number II",
			args: args{
				nums: []int{1, 2, 2, 2, 5, 5, 5},
			},
			want: 1,
		}, {
			name: "single number II",
			args: args{
				nums: []int{2, 2, 3, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberII(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bitwise and",
			args: args{
				m: 5,
				n: 7,
			},
			want: 4,
		}, {
			name: "bitwise and",
			args: args{
				m: 0,
				n: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find min in rotated sorted nums",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		}, {
			name: "find min in rotated sorted nums",
			args: args{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find min in rotated sorted nums with duplication",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 1,
		}, {
			name: "find min in rotated sorted nums",
			args: args{
				nums: []int{1, 0, 1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinII(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInRotatedSortedArray(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search in rotated sorted array",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		}, {
			name: "search in rotated sorted array",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInRotatedSortedArray(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInRotatedSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "can jump",
			args: args{
				nums: []int{0, 2, 3},
			},
			want: false,
		}, {
			name: "can jump",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "next greater element in cycle",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{2, -1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
