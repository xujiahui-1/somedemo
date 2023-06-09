package fuzzing_demo

import "testing"

func TestEqual(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Slices have the same length",
			args: args{a: []byte{102, 97, 108, 99}, b: []byte{102, 97, 108, 99}},
			want: true,
		},
		{
			name: "Slices don’t have the  same length",
			args: args{a: []byte{102, 97, 99}, b: []byte{102, 97, 108, 99}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

//模糊测试
/* 规范:
您的函数名需要以 Fuzz 开头。
您应该接受 *testing.F 类型到您的函数。
您应该只针对一个函数。
*/
//可以使用go test -fuzz命令，也可以直接点run
func FuzzEqual(f *testing.F) {
	// 只要测试一个函数
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		// a、b 的值将被自动生成并传递
		Equal(a, b) //这里对Equal函数进行测试
		// 如果两个自动生成的值都匹配，则 Equal 匹配。如果不匹配，则返回 False。
	})
}
