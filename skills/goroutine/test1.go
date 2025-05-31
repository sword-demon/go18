package main

import (
	"fmt"
	"time"
)

func makeBuns(filling string) {
	fmt.Printf("开始做%s馅的包子\n", filling)
	fmt.Printf("开始剁%s馅...\n", filling)
	fmt.Println("开始擀皮")
	time.Sleep(time.Second)
	fmt.Printf("开始包%s馅的包子...\n", filling)
	fmt.Printf("开始蒸%s的馅的包子...\n", filling)
	// 蒸好了
	time.Sleep(time.Second)
	fmt.Printf("%s的馅的包子已经蒸好了\n", filling)
	// 上菜 通过通道来上菜
	// 在这个位置把包好的包子放到通道内
}

func waiter() {
	// 把蒸好的包子拿出来去上菜
}

func main() {
	fillings := []string{"韭菜", "鸡蛋", "猪肉", "西葫芦"}
	for _, v := range fillings {
		go makeBuns(v) // 主程序创建了协程，并不会等待所有的协程执行成功，需要主程序等待协程处理完成之后
	}

	time.Sleep(10 * time.Second) // 不建议使用
}
