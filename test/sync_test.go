/**
 * @Author: shershon
 * @Description:
 * @Date: 2022/03/20 12:33 PM
 */

package test

import (
	"shershon1991/go-tools/app/syncpkg"
	"testing"
)

func TestWaitGrp(t *testing.T) {
	//syncpkg.NoWaitGroup()
	syncpkg.UseWaitGroup()
}

func TestMutex(t *testing.T) {
	//syncpkg.NoAsStructFieldMutex()
	syncpkg.AsStructFieldMutex()
}

func TestRWMutex(t *testing.T) {
	//syncpkg.NoAsStructFieldRWMutex()
	syncpkg.AsStructFieldRWMutex()
}

func TestCond(t *testing.T) {
	syncpkg.TestCond()
}

func TestOnce(t *testing.T) {
	syncpkg.TestOnce()
}

func TestPool(t *testing.T) {
	syncpkg.TestPool()
}

func TestMap(t *testing.T) {
	//syncpkg.TestMap1()
	//syncpkg.TestMap2()
	syncpkg.TestMap3()
}
