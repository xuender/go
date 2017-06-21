package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type Alook struct {
	old []int
	i   int
	b   int  //上次B出的
	w   bool //上次胜利
}

type Blook struct {
	old []int
	all []int
	a   int //这把A出的
	i   int
}

func a1(look Alook) int {
	// A随便出
	return r.Intn(2)
}

func b1(look Blook) int {
	// B和A出的一样，胜率49%
	return look.a
}

func a2(look Alook) int {
	// 单数随便出， 偶数看B单数如何出
	if look.i/2 == 0 {
		return r.Intn(2)
	}
	return look.b
}

func b2(look Blook) int {
	// 单数出下局，偶数出当前 胜率依然是49
	if look.i/2 == 0 {
		return look.all[look.i+1]
	}
	return look.all[look.i]
}

func a3(look Alook) int {
	// 第一把随便出，如果输了就看B出什么下次出
	if look.i == 0 || look.w {
		return r.Intn(2)
	}
	return look.b
}

func b3(look Blook) int {
	// 如果A出的对，就出一样的，如果A出的错，就出下把的 胜率 66
	if look.a == look.all[look.i] {
		return look.a
	}
	return look.all[look.i+1]
}

// http://www.matrix67.com/blog/archives/5523
// 有时间做剩下的
func main() {
	fmt.Println("庄家的秘密序列  -- 开始")
	const max = 900000
	var win = 0
	var 序列 = make([]int, max)
	for i := 0; i < max; i++ {
		序列[i] = r.Intn(2)
	}
	al := new(Alook)
	bl := new(Blook)
	bl.all = 序列
	for i := 0; i < max; i++ {
		al.i = i
		al.old = 序列[0:i]
		a := a3(*al)

		bl.i = i
		bl.old = 序列[0:i]
		bl.a = a
		b := b3(*bl)
		al.b = b
		if 序列[i] == a && a == b {
			win += 1
			al.w = true
		} else {
			al.w = false
		}
	}
	fmt.Println("胜利:", win)
	fmt.Println("胜率:", win*100.0/max)
}
