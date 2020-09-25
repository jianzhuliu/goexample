/*
无符号整数类型T值的原子减法操作
1、第二个实参为类型为T的一个变量值v。 因为-v在Go中是合法的，所以-v可以直接被用做AddT调用的第二个实参。
2、第二个实参为一个正整数常量c，这时-c在Go中是编译不通过的，所以它不能被用做AddT调用的第二个实参。 这时我们可以使用^T(c-1)（仍为一个正数）做为AddT调用的第二个实参

此^T(v-1)小技巧对于无符号类型的变量v也是适用的，但是^T(v-1)比T(-v)的效率要低。

对于这个^T(c-1)小技巧，如果c是一个类型确定值并且它的类型确实就是T，则它的表示形式可以简化为^(c-1)。

*/

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var (
		n uint64 = 97
		m uint64 = 1
		k int    = 2
	)
	const (
		a        = 3
		b uint64 = 4
		c uint32 = 5
		d int    = 6
	)

	show := fmt.Println
	atomic.AddUint64(&n, -m)
	show(n) // 96 (97 - 1)
	atomic.AddUint64(&n, -uint64(k))
	show(n) // 94 (96 - 2)
	atomic.AddUint64(&n, ^uint64(a-1))
	show(n) // 91 (94 - 3)
	atomic.AddUint64(&n, ^(b - 1))
	show(n) // 87 (91 - 4)
	atomic.AddUint64(&n, ^uint64(c-1))
	show(n) // 82 (87 - 5)
	atomic.AddUint64(&n, ^uint64(d-1))
	show(n) // 76 (82 - 6)
	x := b
	atomic.AddUint64(&n, -x)
	show(n) // 72 (76 - 4)
	atomic.AddUint64(&n, ^(m - 1))
	show(n) // 71 (72 - 1)
	atomic.AddUint64(&n, ^uint64(k-1))
	show(n) // 69 (71 - 2)
}
