// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	examples := [][]string{
		{
			`[1,2,2,3,1]`, 
			`true`,
		},
		{
			`[5,1,8,8,1,5]`, 
			`true`,
		},
		{
			`[1,2,3,4]`, 
			`false`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, isPalindrome, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/cnunionpay-2022spring/problems/D7rekZ/
