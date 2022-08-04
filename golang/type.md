# 值类型和引用类型

### 类型特点

#### 值类型

**基本数据类型：**`int`，`float`，`bool`，`string`，以及`数组`和`struct`

**特点：**变量直接存储值，内存通常在栈中分配，栈在函数调用完会被释放

#### 引用类型

**基本数据类型：** `指针`，`slice`，`map`，`channel`都是引用类型

**特点：**变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配，通过GC回收。

### 区别

 值类型在传参数时，是传的值的Copy，如：

```go
package main

import "fmt"

func main() {
	a := 2
	func01(2)
	fmt.Printf("a的值为:%d\n", a)

}

func func01(a int) {
	a = 7
}

//输出:
//	a的值为:2

```

引用类型传参时，传的是地址，因为引用类型本身存的就是一个地址，所有在函数中对变量的操作都会影响外部的值:

```go
package main

import "fmt"

func main() {
	m := []map[string]string{{"type": "10", "msg": "hello."}}
	checkMissNode(m)
	fmt.Printf("print mString:%s\n", m)
}

func checkMissNode(items []map[string]string) {
	items[0]["one"] = `6`
}

//输出：
//print mString:[map[msg:hello. one:6 type:10]]

```

变量`m`和`items`都存了相同的地址,如下所示：

```go
package main

import "fmt"

func main() {
	m := []map[string]string{{"type": "10", "msg": "hello."}}
	checkMissNode(m)
	fmt.Printf("m的地址为:%x\n", m)
}

func checkMissNode(items []map[string]string) {
	items[0]["one"] = `6`
	fmt.Printf("items的地址为:%x\n", items)

}

/**
* 输出 ：
* items的地址为:[map[6d7367:68656c6c6f2e 6f6e65:36 74797065:3130]]
*     m的地址为:[map[6d7367:68656c6c6f2e 6f6e65:36 74797065:3130]]
*
 */

```

但是，下面的代码，在函数中，又将m2的地址赋值给了items，所以对items的修改不会影响外部值

```go
package main

import "fmt"

func main() {
	m := map[string]string{"type": "10", "msg": "hello."}
	checkMissNode(m)
	fmt.Printf("m的地址为:%x\n", m)
}

func checkMissNode(items map[string]string) {
	m2 := make(map[string]string, 0)
	items = m2
	items["one"] = "1"
	fmt.Printf("items的地址为:%x\n", items)

}

//输出
//items的地址为:map[6f6e65:31]
//m的地址为:map[6d7367:68656c6c6f2e 74797065:3130]

```
