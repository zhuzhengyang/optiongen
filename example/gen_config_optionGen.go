// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optiongen

package example

import (
	"log"
	"sync/atomic"
	"unsafe"
)

// Google Public DNS provides two distinct DoH APIs at these endpoints
// Using the GET method can reduce latency, as it is cached more effectively.
// RFC 8484 GET requests must have a ?dns= query parameter with a Base64Url encoded DNS message. The GET method is the only method supported for the JSON API.

// Config struct
type Config struct {
	// test comment 1
	// test comment 2
	TestNil             interface{}       `xconf:"re3"` // test comment 3
	TestInt             int               `xconf:"test_int" usage:"这里是函数注释1,\"test\"  这里是函数注释2"`
	TestInt64           int64             `xconf:"test_int64"`
	TestSliceInt        []int             `xconf:"test_slice_int"`
	TestSliceInt64      []int64           `xconf:"test_slice_int64"`
	TestSliceString     []string          `xconf:"test_slice_string"`
	TestSliceBool       []bool            `xconf:"test_slice_bool"`
	TestSliceIntNil     []int             `xconf:"test_slice_int_nil"`
	TestSliceByte       []byte            `xconf:"test_slice_byte"`
	TestSliceIntEmpty   []int             `xconf:"test_slice_int_empty"`
	TestHTTPPort        string            `xconf:"test_http_port"`
	TestEmptyMap        map[int]int       `xconf:"test_empty_map"`
	TestMapIntInt       map[int]int       `xconf:"test_map_int_int"`
	TestMapIntString    map[int]string    `xconf:"test_map_int_string"`
	TestMapStringInt    map[string]int    `xconf:"test_map_string_int"`
	TestMapStringString map[string]string `xconf:"test_map_string_string"`
	TestString          string            `xconf:"test_string"`
	Food                *string           `xconf:"food"`
	Walk                func()            `xconf:"walk"`
	TestNilFunc         func()            `xconf:"test_nil_func"` // 中文1
	SubTest             *SubTest          `xconf:"sub_test"`
	FOO                 *FOO              `xconf:"foo"`
	TestProtected       []byte            `xconf:"test_protected"`
	SpecSub             *Spec             `xconf:"spec_sub"`
	TestParamterInt     bool              `xconf:"test_paramter_int"` // reserved parameter 1
	TestParamterStr     string            `xconf:"test_paramter_str"` // reserved parameter 2
}

// SetOption apply single option
func (cc *Config) SetOption(opt ConfigOption) {
	_ = opt(cc)
}

// ApplyOption apply mutiple options
func (cc *Config) ApplyOption(opts ...ConfigOption) {
	for _, opt := range opts {
		_ = opt(cc)
	}
}

// GetSetOption apply new option and return the old optuon
// sample:
// old := cc.GetSetOption(WithTimeout(time.Second))
// defer cc.SetOption(old)
func (cc *Config) GetSetOption(opt ConfigOption) ConfigOption {
	return opt(cc)
}

// ConfigOption option func
type ConfigOption func(cc *Config) ConfigOption

// WithTestNil option func for TestNil
func WithTestNil(v interface{}) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestNil
		cc.TestNil = v
		return WithTestNil(previous)
	}
}

// 这里是函数注释1,&#34;test&#34;
// 这里是函数注释2
// WithTestInt option func for TestInt
func WithTestInt(v int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestInt
		cc.TestInt = v
		return WithTestInt(previous)
	}
}

// WithTestInt64 option func for TestInt64
func WithTestInt64(v int64) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestInt64
		cc.TestInt64 = v
		return WithTestInt64(previous)
	}
}

// WithTestSliceInt option func for TestSliceInt
func WithTestSliceInt(v ...int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceInt
		cc.TestSliceInt = v
		return WithTestSliceInt(previous...)
	}
}

// WithTestSliceInt64 option func for TestSliceInt64
func WithTestSliceInt64(v ...int64) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceInt64
		cc.TestSliceInt64 = v
		return WithTestSliceInt64(previous...)
	}
}

// WithTestSliceString option func for TestSliceString
func WithTestSliceString(v ...string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceString
		cc.TestSliceString = v
		return WithTestSliceString(previous...)
	}
}

// WithTestSliceBool option func for TestSliceBool
func WithTestSliceBool(v ...bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceBool
		cc.TestSliceBool = v
		return WithTestSliceBool(previous...)
	}
}

// WithTestSliceIntNil option func for TestSliceIntNil
func WithTestSliceIntNil(v ...int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceIntNil
		cc.TestSliceIntNil = v
		return WithTestSliceIntNil(previous...)
	}
}

// WithTestSliceByte option func for TestSliceByte
func WithTestSliceByte(v []byte) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceByte
		cc.TestSliceByte = v
		return WithTestSliceByte(previous)
	}
}

// WithTestSliceIntEmpty option func for TestSliceIntEmpty
func WithTestSliceIntEmpty(v ...int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestSliceIntEmpty
		cc.TestSliceIntEmpty = v
		return WithTestSliceIntEmpty(previous...)
	}
}

// WithTestHTTPPort option func for TestHTTPPort
func WithTestHTTPPort(v string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestHTTPPort
		cc.TestHTTPPort = v
		return WithTestHTTPPort(previous)
	}
}

// WithTestEmptyMap option func for TestEmptyMap
func WithTestEmptyMap(v map[int]int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestEmptyMap
		cc.TestEmptyMap = v
		return WithTestEmptyMap(previous)
	}
}

// WithTestMapIntInt option func for TestMapIntInt
func WithTestMapIntInt(v map[int]int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestMapIntInt
		cc.TestMapIntInt = v
		return WithTestMapIntInt(previous)
	}
}

// WithTestMapIntString option func for TestMapIntString
func WithTestMapIntString(v map[int]string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestMapIntString
		cc.TestMapIntString = v
		return WithTestMapIntString(previous)
	}
}

// WithTestMapStringInt option func for TestMapStringInt
func WithTestMapStringInt(v map[string]int) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestMapStringInt
		cc.TestMapStringInt = v
		return WithTestMapStringInt(previous)
	}
}

// WithTestMapStringString option func for TestMapStringString
func WithTestMapStringString(v map[string]string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestMapStringString
		cc.TestMapStringString = v
		return WithTestMapStringString(previous)
	}
}

// WithTestString option func for TestString
func WithTestString(v string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestString
		cc.TestString = v
		return WithTestString(previous)
	}
}

// WithFood option func for Food
func WithFood(v *string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.Food
		cc.Food = v
		return WithFood(previous)
	}
}

// WithWalk option func for Walk
func WithWalk(v func()) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.Walk
		cc.Walk = v
		return WithWalk(previous)
	}
}

// WithTestNilFunc option func for TestNilFunc
func WithTestNilFunc(v func()) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.TestNilFunc
		cc.TestNilFunc = v
		return WithTestNilFunc(previous)
	}
}

// WithSubTest option func for SubTest
func WithSubTest(v *SubTest) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.SubTest
		cc.SubTest = v
		return WithSubTest(previous)
	}
}

// WithFOO option func for FOO
func WithFOO(v *FOO) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.FOO
		cc.FOO = v
		return WithFOO(previous)
	}
}

// WithSpecSub option func for SpecSub
func WithSpecSub(v *Spec) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.SpecSub
		cc.SpecSub = v
		return WithSpecSub(previous)
	}
}

// NewFuncNameSpecified(testParamterInt bool,testParamterStr string, opts... ConfigOption) new Config
func NewFuncNameSpecified(testParamterInt bool, testParamterStr string, opts ...ConfigOption) *Config {
	cc := newDefaultConfig()
	cc.TestParamterInt = testParamterInt
	cc.TestParamterStr = testParamterStr

	for _, opt := range opts {
		_ = opt(cc)
	}
	if watchDogConfig != nil {
		watchDogConfig(cc)
	}
	return cc
}

// InstallConfigWatchDog the installed func will called when NewFuncNameSpecified(testParamterInt bool,testParamterStr string, opts... ConfigOption)  called
func InstallConfigWatchDog(dog func(cc *Config)) {
	watchDogConfig = dog
}

// watchDogConfig global watch dog
var watchDogConfig func(cc *Config)

// newDefaultConfig new default Config
func newDefaultConfig() *Config {
	cc := &Config{
		TestProtected:   nil,
		TestParamterInt: false,
		TestParamterStr: "",
	}

	for _, opt := range [...]ConfigOption{
		WithTestNil(nil),
		WithTestInt(32),
		WithTestInt64(32),
		WithTestSliceInt([]int{1, 2, 3}...),
		WithTestSliceInt64([]int64{1, 2, 3}...),
		WithTestSliceString([]string{"test1", "test2"}...),
		WithTestSliceBool([]bool{false, true}...),
		WithTestSliceIntNil(nil...),
		WithTestSliceByte(nil),
		WithTestSliceIntEmpty(make([]int, 0)...),
		WithTestHTTPPort(""),
		WithTestEmptyMap(make(map[int]int, 0)),
		WithTestMapIntInt(map[int]int{1: 1, 2: 2, 3: 3}),
		WithTestMapIntString(map[int]string{1: "test"}),
		WithTestMapStringInt(map[string]int{"test": 1}),
		WithTestMapStringString(map[string]string{"test": "test"}),
		WithTestString("Meow"),
		WithFood(nil),
		WithWalk(func() {
			log.Println("Walking")
		}),
		WithTestNilFunc(nil),
		WithSubTest(&SubTest{}),
		WithFOO(nil),
		WithSpecSub(NewSpec()),
	} {
		_ = opt(cc)
	}

	return cc
}

// AtomicSetFunc used for XConf
func (cc *Config) AtomicSetFunc() func(interface{}) { return AtomicConfigSet }

// atomicConfig global *Config holder
var atomicConfig unsafe.Pointer

// AtomicConfigSet atomic setter for *Config
func AtomicConfigSet(update interface{}) {
	atomic.StorePointer(&atomicConfig, (unsafe.Pointer)(update.(*Config)))
}

// AtomicConfig return atomic *Config visitor
func AtomicConfig() ConfigVisitor {
	current := (*Config)(atomic.LoadPointer(&atomicConfig))
	if current == nil {
		atomic.CompareAndSwapPointer(&atomicConfig, nil, (unsafe.Pointer)(newDefaultConfig()))
		return (*Config)(atomic.LoadPointer(&atomicConfig))
	}
	return current
}

// all getter func
// GetTestNil return interface{}
func (cc *Config) GetTestNil() interface{} { return cc.TestNil }

// GetTestInt return int
func (cc *Config) GetTestInt() int { return cc.TestInt }

// GetTestInt64 return int64
func (cc *Config) GetTestInt64() int64 { return cc.TestInt64 }

// GetTestSliceInt return []int
func (cc *Config) GetTestSliceInt() []int { return cc.TestSliceInt }

// GetTestSliceInt64 return []int64
func (cc *Config) GetTestSliceInt64() []int64 { return cc.TestSliceInt64 }

// GetTestSliceString return []string
func (cc *Config) GetTestSliceString() []string { return cc.TestSliceString }

// GetTestSliceBool return []bool
func (cc *Config) GetTestSliceBool() []bool { return cc.TestSliceBool }

// GetTestSliceIntNil return []int
func (cc *Config) GetTestSliceIntNil() []int { return cc.TestSliceIntNil }

// GetTestSliceByte return []byte
func (cc *Config) GetTestSliceByte() []byte { return cc.TestSliceByte }

// GetTestSliceIntEmpty return []int
func (cc *Config) GetTestSliceIntEmpty() []int { return cc.TestSliceIntEmpty }

// GetTestHTTPPort return string
func (cc *Config) GetTestHTTPPort() string { return cc.TestHTTPPort }

// GetTestEmptyMap return map[int]int
func (cc *Config) GetTestEmptyMap() map[int]int { return cc.TestEmptyMap }

// GetTestMapIntInt return map[int]int
func (cc *Config) GetTestMapIntInt() map[int]int { return cc.TestMapIntInt }

// GetTestMapIntString return map[int]string
func (cc *Config) GetTestMapIntString() map[int]string { return cc.TestMapIntString }

// GetTestMapStringInt return map[string]int
func (cc *Config) GetTestMapStringInt() map[string]int { return cc.TestMapStringInt }

// GetTestMapStringString return map[string]string
func (cc *Config) GetTestMapStringString() map[string]string { return cc.TestMapStringString }

// GetTestString return string
func (cc *Config) GetTestString() string { return cc.TestString }

// GetFood return *string
func (cc *Config) GetFood() *string { return cc.Food }

// GetWalk return func()
func (cc *Config) GetWalk() func() { return cc.Walk }

// GetTestNilFunc return func()
func (cc *Config) GetTestNilFunc() func() { return cc.TestNilFunc }

// GetSubTest return *SubTest
func (cc *Config) GetSubTest() *SubTest { return cc.SubTest }

// GetFOO return *FOO
func (cc *Config) GetFOO() *FOO { return cc.FOO }

// GetTestProtected return []byte
func (cc *Config) GetTestProtected() []byte { return cc.TestProtected }

// GetSpecSub return SpecVisitor
func (cc *Config) GetSpecSub() SpecVisitor { return cc.SpecSub }

// GetTestParamterInt return bool
func (cc *Config) GetTestParamterInt() bool { return cc.TestParamterInt }

// GetTestParamterStr return string
func (cc *Config) GetTestParamterStr() string { return cc.TestParamterStr }

// ConfigVisitor visitor interface for Config
type ConfigVisitor interface {
	GetTestNil() interface{}
	GetTestInt() int
	GetTestInt64() int64
	GetTestSliceInt() []int
	GetTestSliceInt64() []int64
	GetTestSliceString() []string
	GetTestSliceBool() []bool
	GetTestSliceIntNil() []int
	GetTestSliceByte() []byte
	GetTestSliceIntEmpty() []int
	GetTestHTTPPort() string
	GetTestEmptyMap() map[int]int
	GetTestMapIntInt() map[int]int
	GetTestMapIntString() map[int]string
	GetTestMapStringInt() map[string]int
	GetTestMapStringString() map[string]string
	GetTestString() string
	GetFood() *string
	GetWalk() func()
	GetTestNilFunc() func()
	GetSubTest() *SubTest
	GetFOO() *FOO
	GetTestProtected() []byte
	GetSpecSub() SpecVisitor
	GetTestParamterInt() bool
	GetTestParamterStr() string
}
