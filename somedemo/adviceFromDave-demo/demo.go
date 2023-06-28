package adviceFromDave_demo

import (
	"net/http"
	"testing"
)

/*
1.避免使用任何类似变量类型的后缀。如
	var usersMap map[string]*User
应写为
	var users map[string]*User
此建议也适用于函数参数
	func WriteConfig(w io.Writer, config *Config)
应写为
	func WriteConfig(w io.Writer, conf或c *Config)


2.向变量或常量添加注释时，该注释应描述变量内容，而不是变量目的
	const randomNumber = 6 // determined from an unbiased die
	对于没有初始值的变量，注释应描述谁负责初始化此变量

3.好的代码是最好的文档。 在即将添加注释时，请问下自己，“如何改进代码以便不需要此注释？' 改进代码使其更清晰

4.我建议改进 utils 或 helpers 包的名称是分析它们的调用位置，
如果可能的话，将相关的函数移动到调用者的包中。即使这涉及复制一些 helper 程序代码，这也比在两个程序包之间引入导入依赖项更好。

5.在 Go 语言中有两种很好的方法可以实现松散耦合
	使用接口来描述函数或方法所需的行为。
	避免使用全局状态

6.选择更少，更大的包 你应该做的是不创建新的程序包。 这将导致太多类型被公开，为你的包创建一个宽而浅的API。

7.APIs should be easy to use and hard to misuse. (API 应该易于使用且难以被误用) — Josh Bloch [3]

8.首选可变参数函数而非 []T 参数
	func anyPositive(first int, rest ...int) bool {
	}

*/

func Test_dave(t *testing.T) {
	http.ListenAndServe("sdsafsa", nil)
}
