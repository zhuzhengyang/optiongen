package example

import (
	"log"
)

// Google Public DNS provides two distinct DoH APIs at these endpoints
// Using the GET method can reduce latency, as it is cached more effectively.
// RFC 8484 GET requests must have a ?dns= query parameter with a Base64Url encoded DNS message. The GET method is the only method supported for the JSON API.

//go:generate optiongen --option_with_struct_name=false --v=true
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		// test comment 1
		// test comment 2
		"TestNil":           nil,   // test comment 3
		"TestBool":          false, // test comment 4
		"TestInt":           32,    // @MethodComment(这里是函数注释1) @MethodComment(这里是函数注释2)
		"TestInt64":         int64(32),
		"TestSliceInt":      []int{1, 2, 3},
		"TestSliceInt64":    []int64{1, 2, 3},
		"TestSliceString":   []string{"test1", "test2"},
		"TestSliceBool":     []bool{false, true},
		"TestSliceIntNil":   []int(nil),
		"TestSliceByte":     []byte(nil),
		"TestSliceIntEmpty": []int{},

		"TestMapIntInt":       map[int]int{1: 1, 2: 2, 3: 3},
		"TestMapIntString":    map[int]string{1: "test"},
		"TestMapStringInt":    map[string]int{"test": 1},
		"TestMapStringString": map[string]string{"test": "test"},

		"TestString": "Meow",
		"Food":       (*string)(nil),
		"Walk": func() {
			log.Println("Walking")
		},
		"TestNilFunc":    (func())(nil), // 中文1
		"TestReserved1_": []byte(nil),   // 在调优或者运行阶段，我们可能需要动态查看连接池中的一些指标，
		// 来判断设置的值是否合理，或者检测连接池是否有异常情况出现
		"TestReserved2Inner": 1,
	}
}

// HTTP parsing and communication with DNS resolver was successful, and the response body content is a DNS response in either binary or JSON encoding,
// depending on the query endpoint, Accept header and GET parameters.

//go:generate optiongen --option_with_struct_name=false --v=true
func SpecOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		// test comment 5
		// test comment 6
		"TestNil1":  nil,   // test comment 1
		"TestBool1": false, // test comment 2
		"TestInt1":  32,    // @MethodComment(这里是函数注释3) @MethodComment(这里是函数注释4)

		"TestNilFunc1":   (func())(nil), // 中文2
		"TestReserved2_": []byte(nil),   // sql.DB对外暴露出了其运行时的状态db.DBStats，sql.DB在关闭，创建，释放连接时候，会维护更新这个状态。
		// 我们可以通过prometheus来收集连接池状态，然后在grafana面板上配置指标，使指标可以动态的展示。
		"TestReserved2Inner1": 1,
	}
}
