package regexp_demo

import (
	"fmt"
	"regexp"
	"testing"
)

//regexp 正则表达式 demo

// Match 查询相对简单的正则表达式,MatchReader 则是查询reader中，MatchString则是查询string中
func Test_Match(t *testing.T) {
	matched, err := regexp.Match(`foo.*`, []byte(`searfood`))
	fmt.Println(matched, err)

	matched, err = regexp.MatchString(`foo.*`, "seafoeoed")
	fmt.Println(matched, err)

	fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`))

}

// regexp结构体内部操作是线程安全的
func Test_Regexp(t *testing.T) {
	text := "Hello  世界!123 go."

	// Compile 输入正则表达式，进行解析
	// MastCompile 和Compile一样，但是当正则不合法时，直接paincs了
	compile, err := regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println("正则出错了:", err)
	}

	// FindAllString 返回所有匹配结果的切片,n=-1时返回全部，其他数字时就返回几个
	// FindAllIndex 返回匹配项的下标数组
	allString := compile.FindAllString(text, -1)
	fmt.Println(allString)

	//FindStringSubmatch 带match的就是把子匹配项也都匹配进去
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))

	// ReplaceAllString 查询并替换 所有匹配项
	replaceAllString := compile.ReplaceAllString(text, "aaaaaa")
	fmt.Println(replaceAllString)

	// Longest 方法让你的regexp返回最左最长的匹配项
	re = regexp.MustCompile(`a(|b)`)
	fmt.Println(re.FindString("ab"))
	re.Longest()
	fmt.Println(re.FindString("ab"))

}

// Expand 为你的pattern添加模版，根据定制的规则修改你的结果
func Test_Expand(t *testing.T) {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`)

	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := []byte("$key=$value\n")

	result := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))

}
