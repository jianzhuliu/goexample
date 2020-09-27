/*
1、单向接收通道无法被关闭。
2、发送一个值到一个已关闭的通道被视为一个非阻塞操作，该操作会导致恐慌。

*/

package main

func main() {
	var c = make(chan bool)
	close(c)
	select {
	case <-c:
	case c <- true: // panic: 向已关闭的通道发送数据
	default:
	}
}

/*
func foo(c <-chan int) {
	close(c) // error: 不能关闭单向接收通道
}
*/
