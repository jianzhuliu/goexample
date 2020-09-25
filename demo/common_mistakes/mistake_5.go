/*
如果一个Timer值已经过期或者已经被终止（stopped），则相应的Stop方法调用返回false。 在此Timer值尚未终止的时候，Stop方法调用返回false只能意味着此Timer值已经过期

一个Timer值被终止之后，它的通道字段C最多只能含有一个过期的通知

在一个Timer终止（stopped）之后并且在重置和重用此Timer值之前，我们应该确保此Timer值中肯定不存在过期的通知

一个*Timer值的Reset方法必须在对应Timer值过期或者终止之后才能被调用； 否则，此Reset方法调用和一个可能的向此Timer值的C通道字段的发送通知操作产生数据竞争

*/
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer := time.NewTimer(time.Second / 2)
	select {
	case <-timer.C:
	default:
		time.Sleep(time.Second) // 此分支被选中的可能性较大
	}
	timer.Reset(time.Second * 10) // 可能数据竞争
	<-timer.C
	fmt.Println(time.Since(start)) // 大约1s
}
