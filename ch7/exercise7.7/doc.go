// 练习 7.7： 解释为什么帮助信息在它的默认值是20.0
// 没有包含°C的情况下输出了°C。

// 因为 temp 的类型是 *Celsius，而 *Celsius 实现了 String() 方法
// 所以在调用 fmt.Println(*temp) 会调用 String() 方法
package doc
