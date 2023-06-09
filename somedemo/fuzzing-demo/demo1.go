package fuzzing_demo

// 测试使用golang自带的模糊测试功能来对下面这个方法进行测试

// Equal 一个对比切片是否完全一致的方法
func Equal(a []byte, b []byte) bool {
	// 检查切片的长度是否相同
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		// 检查同一索引中的元素是否相同
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//单元测试方法
//command +n -->Test for funcation --> 测试文件生成后，使用go test 命令进行测试
