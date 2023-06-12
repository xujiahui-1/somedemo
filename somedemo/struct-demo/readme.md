# 空结构体探究
空结构体特点:
1. 零内存占用
2. 地址相同
3. 无状态

## 为什么空结构体零内存占用？
在golang中，存在一个变量叫`zerobase`，这是一个 `uintptr` 全局变量，占用 8 个字节。
当你定义了 内存 size 为 0 的内存分配，那么golang就会把这个`&zerobase`这个地址付给他
所以无论你有多少个空结构体，他底层都是这个`zerobase`
## 内存管理特殊处理
>译器在编译期间，识别到 struct {} 这种特殊类型的内存分配，会统统分配出 runtime.zerobase 的地址出去，这个代码逻辑是在 mallocgc 函数里面
```go
func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
    // 分配 size 为 0 的结构体，把全局变量 zerobase 的地址给出去即可；
	if size == 0 {
		return unsafe.Pointer(&zerobase)
	}
    // ... 
```
## 偏移量计算 
1.结构体的成员变量，第一个成员变量的偏移量为 0。往后的每个成员变量的对齐值必须为编译器默认对齐长度（#pragma pack(n)）或当前成员变量类型的长度（unsafe.Sizeof），取最小值作为当前类型的对齐值。
>其偏移量必须为对齐值的整数倍

2。**结构体本身，对齐值必须为编译器默认对齐长度**（#pragma pack(n)）或结构体的所有成员变量类型中的最大长度，取最大数的最小整数倍作为对齐值

