/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/05/17 14:30
 */

package test

import (
	"shershon1991/go-tools/app/funkpkg"
	"testing"
)

// 切片(slice)操作
// 1.判断元素是否存在
func TestExist(t *testing.T) {
	funkpkg.Exist()
}

// 2.查找元素第一次出现位置
func TestIndexOf(t *testing.T) {
	funkpkg.IndexOf()
}

// 3.查找元素最后一次出现位置
func TestLastIndexOf(t *testing.T) {
	funkpkg.LastIndexOf()
}

// 4.批量查找(都有则True)
func TestEvery(t *testing.T) {
	funkpkg.Every()
}

// 5.批量查找(有一则True)
func TestSome(t *testing.T) {
	funkpkg.Some()
}

// 6.获取最后或第一个元素
func TestFirstOrLast(t *testing.T) {
	funkpkg.LastOrFirst()
}

// 7.用元素填充切片
func TestFill(t *testing.T) {
	funkpkg.Fill()
}

// 8.取两个切片共同元素结果集
func TestJoin(t *testing.T) {
	funkpkg.Join()
}

// 9.获取去掉两切片共同元素结果集
func TestDiffSlice(t *testing.T) {
	funkpkg.DiffSlice()
}
