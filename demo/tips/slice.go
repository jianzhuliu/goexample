package main

//实现不被推荐的原因是append函数调用可能会修改输入实参opts的底层潜在Option元素序列
func NewX(opts ...Option) *X {
	options := append(opts, defaultOpts...)
	// 使用合并后选项来创建一个X值并返回其指针。
	// ...
}

//避免输入实参的底层Option元素序列被修改，我们应该使用下面的实现方法
func NewX(opts ...Option) *X {
	// 改用三下标子切片格式。
	opts = append(opts[:len(opts):len(opts)], defaultOpts...)
	// 使用合并后选项来创建一个X值并返回其指针。
	// ...
}
