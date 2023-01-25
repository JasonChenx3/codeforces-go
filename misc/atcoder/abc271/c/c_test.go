// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc271/submit?taskScreenName=abc271_c
// 对拍：https://atcoder.jp/contests/abc271/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc271_c&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`6
1 2 4 6 7 271`,
			`4`,
		},
		{
			`10
1 1 1 1 1 1 1 1 1 1`,
			`5`,
		},
		{
			`1
5`,
			`0`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc271/tasks/abc271_c
