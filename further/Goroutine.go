package main

import (
	"fmt"
	"runtime"
)

/*現在開始弄攜程,可以進行更有效的並行運算，攜程和操作系統線程之間無一對一的關係。線程是輕量級的*/

/*這裡把概率搞懂
go語音是go語音提供的一種用戶態的一種線程

調度器主要分成4個部分，分別是m,g,p,sched

m m代表系統的線程
p p銜接m和g的調度上下文
g g是攜程的實體
*/

func showGOruntime() {
	var Cpu = runtime.NumCPU()
	var Number = runtime.NumGoroutine()
	fmt.Println(Cpu, Number)
}

/*攜程的使用很簡單，直接加上go就行了*/

/*func main() {
	showGOruntime()
}*/
