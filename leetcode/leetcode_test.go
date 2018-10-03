package leetcode

import (
	"fmt"
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

func TestProducerConsumer(t *testing.T) {
	data := make(chan int, 5)
	ProducerWG.Add(1)
	go Producer("P1", data)
	ConsumerWG.Add(1)
	go Consumer("U1", data)
	ProducerWG.Add(1)
	go Producer("P2", data)
	ConsumerWG.Add(1)
	go Consumer("U2", data)

	ProducerWG.Wait()
	close(data)
	ConsumerWG.Wait()
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

func Test198RobHouses(t *testing.T) {
	if n := rob([]int{1, 2, 3, 1}); n != 4 {
		t.Errorf("expected 4 and got %d\n", n)
	}
	if n := rob([]int{2, 7, 9, 3, 1}); n != 12 {
		t.Errorf("expected 12 and got %d\n", n)
	}
	if n := rob([]int{6, 6, 4, 8, 4, 3, 3, 10}); n != 27 {
		t.Errorf("expected 27 and got %d\n", n)
	}
	if n := rob([]int{2, 4, 8, 9, 9, 3}); n != 19 {
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
func Test746ClimbStairs(t *testing.T) {
	if n := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}); n != 6 {
		t.Errorf("expected 6 and got %d\n", n)
	}
}
