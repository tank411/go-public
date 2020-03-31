package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

//判断是否相等，是否符合预期，不符合直接退出，
func TestRequire(t *testing.T) {
	name := "Bob"
	age := 10

	require.Equal(t, "bob", name)
	require.Equal(t, 20, age)

}

//判断是否相等，是否符合预期，
func TestAssert(t *testing.T) {
	name := "Bob"
	age := 10

	assert.Equal(t, "Bob", name)
	assert.Equal(t, 10, age)

}

//判断是否相等，类型和值都要一致，是否符合预期，
func TestExactly(t *testing.T) {
	assert.Exactly(t, int32(123), int64(123))

}

//判断是否包含子字符串
func TestContains(t *testing.T) {
	assert.Contains(t, "Hello World", "World")

}

//判断是否包含子集
func TestSubset(t *testing.T) {
	assert.Subset(t, []int{1, 2, 3}, []int{1, 2}, "But [1, 2, 3] does contain [1, 2]")
}

//没有子集的判断
func TestNotSubset(t *testing.T) {
	assert.NotSubset(t, []int{1, 3, 4}, []int{1, 2}, "But [1, 3, 4] does not contain [1, 2]")

}
func TestElementsMatch(t *testing.T) {
	assert.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
}

//证明Assert函数内有异常
func TestThatItPanicsOnFalseConditions(t *testing.T) {
	assert.Panics(t, func() {
		Assert(false)
	}, "Calling Assert() with a false condition should panic.")
}

//证明Assert函数内没有异常
func TestThatItDoesNotPanicOnTrueConditions(t *testing.T) {
	assert.NotPanics(t, func() {
		Assert(true)
	}, "Calling Assert() with a true condition should not panic.")
}
