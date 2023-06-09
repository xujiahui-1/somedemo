package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// 序列化返序列化的基本使用
func main0() {
	person1 := Person{
		Name:   "xujiahui",
		Age:    24,
		Weight: 80,
	}
	//使用Go标准库将结构体转化成json串
	j, _ := json.Marshal(person1)
	fmt.Printf("str:%s\n", j)

	//反序列化并将值复给person2
	person2 := &Person{}
	json.Unmarshal(j, person2)
	fmt.Println(person2.Name)
}

//结构体tag的使用
/*
1.给结构体打tag以后，就可以在运行时通过反射机制，读取到tag信息
2.我们给结构体的name字段打了json的tag。那么通过结构体生成json串的时候，就会显示小写的name
3.给结构体的某个字段打tag - 。那么序列化的时候就会忽略这个字段
4.给属性添加 omitempty tag，让序列化时忽略空值
5.当我们想要结构体中的匿名嵌套结构体也忽略空值时，即需要打tag，又需要将该属性变成指针类型
*/
func main1() {
	person1 := Person{
		Name: "xujiahui",
		Age:  24,
	}
	marshal, _ := json.Marshal(person1)
	fmt.Printf("str:%s\n", marshal)
}

// 优雅处理字符串格式的数字
/*
前端经常会传输过来string类型的数字数据，我们想使用int接受它们，
我们可以给结构体属性添加 string tag.优雅的接受他

*/
func main2() {
	jsons := `{"id":"1234","number":"12321421"}`
	card1 := &card{}
	err := json.Unmarshal([]byte(jsons), card1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(card1)
	fmt.Printf("%T\n", card1.ID)
	fmt.Printf(" %T\n", card1.Number)
}

// 使用decode返序列化来避免Unmarshal返序列化后数字变成float64类型
func main4() {
	newp := Person{
		Name:   "xujiahui",
		Age:    16,
		Weight: 30,
	}
	marshal, _ := json.Marshal(newp)
	p2 := &Person{}
	decoder := json.NewDecoder(bytes.NewReader(marshal))
	decoder.UseNumber()
	decoder.Decode(p2)
}

// 自定义解析时间字段
/*
2023-06-07 15:13:09    这是我们常用的时间格式
2023-06-07T15:13:09.162263+09:00   这是json包的默认格式
我们需要自己实现接口来自定义json的时间格式
*/
type Post struct {
	CreateTime time.Time `json:"create_time"`
}

const layout = "2006-01-02 15:04:05"

// 实现MarshalJSON接口
func (p *Post) MarshalJSON() ([]byte, error) {
	type TempPost Post
	return json.Marshal(struct {
		CreateTime string `json:"create_time"`
		*TempPost
	}{
		CreateTime: p.CreateTime.Format(layout),
		TempPost:   (*TempPost)(p),
	})
}

func main() {
	p1 := Post{CreateTime: time.Now()}
	b, _ := json.Marshal(&p1)
	fmt.Printf("%s\n", b) //这里输出的是RFC3339格式  {"create_time":"2023-06-07T15:13:09.162263+09:00"}
	//实现MarshalJSON接口后变成 {"create_time":"2023-06-07 15:29:47"}
	//jsonStr := `{"create_time":"2023-06-07 15:13:09"}`
	//var p2 Post
	//err := json.Unmarshal([]byte(jsonStr), &p2)
	//fmt.Println(err)

}
