// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optiongen

package example

import (
	"sync/atomic"
	"time"
	"unsafe"
)

// AllConfig should use NewAllConfig to initialize it
type AllConfig struct {
	TypeBool          bool            `xconf:"type_bool"`
	TypeString        string          `xconf:"type_string"`
	TypeDuration      time.Duration   `xconf:"type_duration"`
	TypeFloat32       float32         `xconf:"type_float32"`
	TypeFloat64       float32         `xconf:"type_float64"`
	TypeInt           int             `xconf:"type_int"`
	TypeUint          int             `xconf:"type_uint"`
	TypeInt8          int8            `xconf:"type_int8"`
	TypeUint8         uint8           `xconf:"type_uint8"`
	TypeInt16         int16           `xconf:"type_int16"`
	TypeUint16        uint16          `xconf:"type_uint16"`
	TypeInt32         int32           `xconf:"type_int32"`
	TypeUint32        uint32          `xconf:"type_uint32"`
	TypeInt64         int64           `xconf:"type_int64"`
	TypeUint64        uint64          `xconf:"type_uint64"`
	TypeSliceInt      []int           `xconf:"type_slice_int"`
	TypeSliceUint     []uint          `xconf:"type_slice_uint"`
	TypeSliceInt8     []int8          `xconf:"type_slice_int8"`
	TypeSliceUint8    []uint8         `xconf:"type_slice_uint8"`
	TypeSliceInt16    []int16         `xconf:"type_slice_int16"`
	TypeSliceUin16    []uint16        `xconf:"type_slice_uin16"`
	TypeSliceInt32    []int32         `xconf:"type_slice_int32"`
	TypeSliceUint32   []uint32        `xconf:"type_slice_uint32"`
	TypeSliceInt64    []int64         `xconf:"type_slice_int64"`
	TypeSliceUint64   []uint64        `xconf:"type_slice_uint64"`
	TypeSliceString   []string        `xconf:"type_slice_string"`
	TypeSliceFloat32  []float32       `xconf:"type_slice_float32"`
	TypeSliceFloat64  []float64       `xconf:"type_slice_float64"`
	TypeSliceDuratuon []time.Duration `xconf:"type_slice_duratuon"`
	// annotation@TypeMapStringIntNotLeaf(xconf="type_map_string_int_not_leaf,notleaf")
	TypeMapStringIntNotLeaf map[string]int           `xconf:"type_map_string_int_not_leaf,notleaf"`
	TypeMapStringInt        map[string]int           `xconf:"type_map_string_int"`
	TypeMapIntString        map[int]string           `xconf:"type_map_int_string"`
	TypeMapStringString     map[string]string        `xconf:"type_map_string_string"`
	TypeMapIntInt           map[int]int              `xconf:"type_map_int_int"`
	TypeMapStringDuration   map[string]time.Duration `xconf:"type_map_string_duration"`
	// annotation@Redis(getter="RedisVisitor")
	Redis         *Redis      `xconf:"redis"`
	ETCD          *ETCD       `xconf:"etcd"`
	TestInterface interface{} `xconf:"test_interface"`
}

// NewAllConfig new AllConfig
func NewAllConfig(opts ...AllConfigOption) *AllConfig {
	cc := newDefaultAllConfig()
	for _, opt := range opts {
		opt(cc)
	}
	if watchDogAllConfig != nil {
		watchDogAllConfig(cc)
	}
	return cc
}

// ApplyOption apply multiple new option and return the old ones
// sample:
// old := cc.ApplyOption(WithTimeout(time.Second))
// defer cc.ApplyOption(old...)
func (cc *AllConfig) ApplyOption(opts ...AllConfigOption) []AllConfigOption {
	var previous []AllConfigOption
	for _, opt := range opts {
		previous = append(previous, opt(cc))
	}
	return previous
}

// AllConfigOption option func
type AllConfigOption func(cc *AllConfig) AllConfigOption

// WithTypeBool option func for filed TypeBool
func WithTypeBool(v bool) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeBool
		cc.TypeBool = v
		return WithTypeBool(previous)
	}
}

// WithTypeString option func for filed TypeString
func WithTypeString(v string) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeString
		cc.TypeString = v
		return WithTypeString(previous)
	}
}

// WithTypeDuration option func for filed TypeDuration
func WithTypeDuration(v time.Duration) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeDuration
		cc.TypeDuration = v
		return WithTypeDuration(previous)
	}
}

// WithTypeFloat32 option func for filed TypeFloat32
func WithTypeFloat32(v float32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeFloat32
		cc.TypeFloat32 = v
		return WithTypeFloat32(previous)
	}
}

// WithTypeFloat64 option func for filed TypeFloat64
func WithTypeFloat64(v float32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeFloat64
		cc.TypeFloat64 = v
		return WithTypeFloat64(previous)
	}
}

// WithTypeInt option func for filed TypeInt
func WithTypeInt(v int) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeInt
		cc.TypeInt = v
		return WithTypeInt(previous)
	}
}

// WithTypeUint option func for filed TypeUint
func WithTypeUint(v int) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeUint
		cc.TypeUint = v
		return WithTypeUint(previous)
	}
}

// WithTypeInt8 option func for filed TypeInt8
func WithTypeInt8(v int8) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeInt8
		cc.TypeInt8 = v
		return WithTypeInt8(previous)
	}
}

// WithTypeUint8 option func for filed TypeUint8
func WithTypeUint8(v uint8) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeUint8
		cc.TypeUint8 = v
		return WithTypeUint8(previous)
	}
}

// WithTypeInt16 option func for filed TypeInt16
func WithTypeInt16(v int16) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeInt16
		cc.TypeInt16 = v
		return WithTypeInt16(previous)
	}
}

// WithTypeUint16 option func for filed TypeUint16
func WithTypeUint16(v uint16) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeUint16
		cc.TypeUint16 = v
		return WithTypeUint16(previous)
	}
}

// WithTypeInt32 option func for filed TypeInt32
func WithTypeInt32(v int32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeInt32
		cc.TypeInt32 = v
		return WithTypeInt32(previous)
	}
}

// WithTypeUint32 option func for filed TypeUint32
func WithTypeUint32(v uint32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeUint32
		cc.TypeUint32 = v
		return WithTypeUint32(previous)
	}
}

// WithTypeInt64 option func for filed TypeInt64
func WithTypeInt64(v int64) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeInt64
		cc.TypeInt64 = v
		return WithTypeInt64(previous)
	}
}

// WithTypeUint64 option func for filed TypeUint64
func WithTypeUint64(v uint64) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeUint64
		cc.TypeUint64 = v
		return WithTypeUint64(previous)
	}
}

// WithTypeSliceInt option func for filed TypeSliceInt
func WithTypeSliceInt(v ...int) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt
		cc.TypeSliceInt = v
		return WithTypeSliceInt(previous...)
	}
}

// AppendTypeSliceInt append func for filed TypeSliceInt
func AppendTypeSliceInt(v ...int) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt
		cc.TypeSliceInt = append(cc.TypeSliceInt, v...)
		return WithTypeSliceInt(previous...)
	}
}

// WithTypeSliceUint option func for filed TypeSliceUint
func WithTypeSliceUint(v ...uint) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUint
		cc.TypeSliceUint = v
		return WithTypeSliceUint(previous...)
	}
}

// AppendTypeSliceUint append func for filed TypeSliceUint
func AppendTypeSliceUint(v ...uint) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUint
		cc.TypeSliceUint = append(cc.TypeSliceUint, v...)
		return WithTypeSliceUint(previous...)
	}
}

// WithTypeSliceInt8 option func for filed TypeSliceInt8
func WithTypeSliceInt8(v ...int8) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt8
		cc.TypeSliceInt8 = v
		return WithTypeSliceInt8(previous...)
	}
}

// AppendTypeSliceInt8 append func for filed TypeSliceInt8
func AppendTypeSliceInt8(v ...int8) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt8
		cc.TypeSliceInt8 = append(cc.TypeSliceInt8, v...)
		return WithTypeSliceInt8(previous...)
	}
}

// WithTypeSliceUint8 option func for filed TypeSliceUint8
func WithTypeSliceUint8(v ...uint8) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUint8
		cc.TypeSliceUint8 = v
		return WithTypeSliceUint8(previous...)
	}
}

// AppendTypeSliceUint8 append func for filed TypeSliceUint8
func AppendTypeSliceUint8(v ...uint8) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUint8
		cc.TypeSliceUint8 = append(cc.TypeSliceUint8, v...)
		return WithTypeSliceUint8(previous...)
	}
}

// WithTypeSliceInt16 option func for filed TypeSliceInt16
func WithTypeSliceInt16(v ...int16) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt16
		cc.TypeSliceInt16 = v
		return WithTypeSliceInt16(previous...)
	}
}

// AppendTypeSliceInt16 append func for filed TypeSliceInt16
func AppendTypeSliceInt16(v ...int16) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt16
		cc.TypeSliceInt16 = append(cc.TypeSliceInt16, v...)
		return WithTypeSliceInt16(previous...)
	}
}

// WithTypeSliceUin16 option func for filed TypeSliceUin16
func WithTypeSliceUin16(v ...uint16) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUin16
		cc.TypeSliceUin16 = v
		return WithTypeSliceUin16(previous...)
	}
}

// AppendTypeSliceUin16 append func for filed TypeSliceUin16
func AppendTypeSliceUin16(v ...uint16) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUin16
		cc.TypeSliceUin16 = append(cc.TypeSliceUin16, v...)
		return WithTypeSliceUin16(previous...)
	}
}

// WithTypeSliceInt32 option func for filed TypeSliceInt32
func WithTypeSliceInt32(v ...int32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt32
		cc.TypeSliceInt32 = v
		return WithTypeSliceInt32(previous...)
	}
}

// AppendTypeSliceInt32 append func for filed TypeSliceInt32
func AppendTypeSliceInt32(v ...int32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt32
		cc.TypeSliceInt32 = append(cc.TypeSliceInt32, v...)
		return WithTypeSliceInt32(previous...)
	}
}

// WithTypeSliceUint32 option func for filed TypeSliceUint32
func WithTypeSliceUint32(v ...uint32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUint32
		cc.TypeSliceUint32 = v
		return WithTypeSliceUint32(previous...)
	}
}

// AppendTypeSliceUint32 append func for filed TypeSliceUint32
func AppendTypeSliceUint32(v ...uint32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUint32
		cc.TypeSliceUint32 = append(cc.TypeSliceUint32, v...)
		return WithTypeSliceUint32(previous...)
	}
}

// WithTypeSliceInt64 option func for filed TypeSliceInt64
func WithTypeSliceInt64(v ...int64) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt64
		cc.TypeSliceInt64 = v
		return WithTypeSliceInt64(previous...)
	}
}

// AppendTypeSliceInt64 append func for filed TypeSliceInt64
func AppendTypeSliceInt64(v ...int64) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceInt64
		cc.TypeSliceInt64 = append(cc.TypeSliceInt64, v...)
		return WithTypeSliceInt64(previous...)
	}
}

// WithTypeSliceUint64 option func for filed TypeSliceUint64
func WithTypeSliceUint64(v ...uint64) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUint64
		cc.TypeSliceUint64 = v
		return WithTypeSliceUint64(previous...)
	}
}

// AppendTypeSliceUint64 append func for filed TypeSliceUint64
func AppendTypeSliceUint64(v ...uint64) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceUint64
		cc.TypeSliceUint64 = append(cc.TypeSliceUint64, v...)
		return WithTypeSliceUint64(previous...)
	}
}

// WithTypeSliceString option func for filed TypeSliceString
func WithTypeSliceString(v ...string) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceString
		cc.TypeSliceString = v
		return WithTypeSliceString(previous...)
	}
}

// AppendTypeSliceString append func for filed TypeSliceString
func AppendTypeSliceString(v ...string) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceString
		cc.TypeSliceString = append(cc.TypeSliceString, v...)
		return WithTypeSliceString(previous...)
	}
}

// WithTypeSliceFloat32 option func for filed TypeSliceFloat32
func WithTypeSliceFloat32(v ...float32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceFloat32
		cc.TypeSliceFloat32 = v
		return WithTypeSliceFloat32(previous...)
	}
}

// AppendTypeSliceFloat32 append func for filed TypeSliceFloat32
func AppendTypeSliceFloat32(v ...float32) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceFloat32
		cc.TypeSliceFloat32 = append(cc.TypeSliceFloat32, v...)
		return WithTypeSliceFloat32(previous...)
	}
}

// WithTypeSliceFloat64 option func for filed TypeSliceFloat64
func WithTypeSliceFloat64(v ...float64) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceFloat64
		cc.TypeSliceFloat64 = v
		return WithTypeSliceFloat64(previous...)
	}
}

// AppendTypeSliceFloat64 append func for filed TypeSliceFloat64
func AppendTypeSliceFloat64(v ...float64) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceFloat64
		cc.TypeSliceFloat64 = append(cc.TypeSliceFloat64, v...)
		return WithTypeSliceFloat64(previous...)
	}
}

// WithTypeSliceDuratuon option func for filed TypeSliceDuratuon
func WithTypeSliceDuratuon(v ...time.Duration) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceDuratuon
		cc.TypeSliceDuratuon = v
		return WithTypeSliceDuratuon(previous...)
	}
}

// AppendTypeSliceDuratuon append func for filed TypeSliceDuratuon
func AppendTypeSliceDuratuon(v ...time.Duration) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeSliceDuratuon
		cc.TypeSliceDuratuon = append(cc.TypeSliceDuratuon, v...)
		return WithTypeSliceDuratuon(previous...)
	}
}

// WithTypeMapStringIntNotLeaf option func for filed TypeMapStringIntNotLeaf
func WithTypeMapStringIntNotLeaf(v map[string]int) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeMapStringIntNotLeaf
		cc.TypeMapStringIntNotLeaf = v
		return WithTypeMapStringIntNotLeaf(previous)
	}
}

// WithTypeMapStringInt option func for filed TypeMapStringInt
func WithTypeMapStringInt(v map[string]int) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeMapStringInt
		cc.TypeMapStringInt = v
		return WithTypeMapStringInt(previous)
	}
}

// WithTypeMapIntString option func for filed TypeMapIntString
func WithTypeMapIntString(v map[int]string) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeMapIntString
		cc.TypeMapIntString = v
		return WithTypeMapIntString(previous)
	}
}

// WithTypeMapStringString option func for filed TypeMapStringString
func WithTypeMapStringString(v map[string]string) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeMapStringString
		cc.TypeMapStringString = v
		return WithTypeMapStringString(previous)
	}
}

// WithTypeMapIntInt option func for filed TypeMapIntInt
func WithTypeMapIntInt(v map[int]int) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeMapIntInt
		cc.TypeMapIntInt = v
		return WithTypeMapIntInt(previous)
	}
}

// WithTypeMapStringDuration option func for filed TypeMapStringDuration
func WithTypeMapStringDuration(v map[string]time.Duration) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TypeMapStringDuration
		cc.TypeMapStringDuration = v
		return WithTypeMapStringDuration(previous)
	}
}

// WithRedis option func for filed Redis
func WithRedis(v *Redis) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.Redis
		cc.Redis = v
		return WithRedis(previous)
	}
}

// WithETCD option func for filed ETCD
func WithETCD(v *ETCD) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.ETCD
		cc.ETCD = v
		return WithETCD(previous)
	}
}

// WithTestInterface option func for filed TestInterface
func WithTestInterface(v interface{}) AllConfigOption {
	return func(cc *AllConfig) AllConfigOption {
		previous := cc.TestInterface
		cc.TestInterface = v
		return WithTestInterface(previous)
	}
}

// InstallAllConfigWatchDog the installed func will called when NewAllConfig  called
func InstallAllConfigWatchDog(dog func(cc *AllConfig)) { watchDogAllConfig = dog }

// watchDogAllConfig global watch dog
var watchDogAllConfig func(cc *AllConfig)

// setAllConfigDefaultValue default AllConfig value
func setAllConfigDefaultValue(cc *AllConfig) {
	for _, opt := range [...]AllConfigOption{
		WithTypeBool(false),
		WithTypeString("a"),
		WithTypeDuration(time.Second),
		WithTypeFloat32(32.32),
		WithTypeFloat64(64.64),
		WithTypeInt(32),
		WithTypeUint(32),
		WithTypeInt8(8),
		WithTypeUint8(8),
		WithTypeInt16(16),
		WithTypeUint16(16),
		WithTypeInt32(32),
		WithTypeUint32(32),
		WithTypeInt64(64),
		WithTypeUint64(64),
		WithTypeSliceInt([]int{1, 2, 3, 4}...),
		WithTypeSliceUint([]uint{1, 2, 3, 4}...),
		WithTypeSliceInt8([]int8{1, 2, 3, 4}...),
		WithTypeSliceUint8([]uint8{1, 2, 3, 4}...),
		WithTypeSliceInt16([]int16{1, 2, 3, 4}...),
		WithTypeSliceUin16([]uint16{1, 2, 3, 4}...),
		WithTypeSliceInt32([]int32{1, 2, 3, 4}...),
		WithTypeSliceUint32([]uint32{1, 2, 3, 4}...),
		WithTypeSliceInt64([]int64{1, 2, 3, 4}...),
		WithTypeSliceUint64([]uint64{1, 2, 3, 4}...),
		WithTypeSliceString([]string{"a", "b", "c"}...),
		WithTypeSliceFloat32([]float32{1.32, 2.32, 3.32, 4.32}...),
		WithTypeSliceFloat64([]float64{1.64, 2.64, 3.64, 4.64}...),
		WithTypeSliceDuratuon([]time.Duration{time.Second, time.Minute, time.Hour}...),
		WithTypeMapStringIntNotLeaf(map[string]int{"a": 1, "b": 2}),
		WithTypeMapStringInt(map[string]int{"a": 1, "b": 2}),
		WithTypeMapIntString(map[int]string{1: "a", 2: "b"}),
		WithTypeMapStringString(map[string]string{"a": "a", "b": "b"}),
		WithTypeMapIntInt(map[int]int{1: 1, 2: 2}),
		WithTypeMapStringDuration(map[string]time.Duration{"read": time.Second, "write": time.Second * 5}),
		WithRedis(NewRedis()),
		WithETCD(NewETCD(time.Second)),
		WithTestInterface(nil),
	} {
		opt(cc)
	}
}

// newDefaultAllConfig new default AllConfig
func newDefaultAllConfig() *AllConfig {
	cc := &AllConfig{}
	setAllConfigDefaultValue(cc)
	return cc
}

// AtomicSetFunc used for XConf
func (cc *AllConfig) AtomicSetFunc() func(interface{}) { return AtomicAllConfigSet }

// atomicAllConfig global *AllConfig holder
var atomicAllConfig unsafe.Pointer

// onAtomicAllConfigSet global call back when  AtomicAllConfigSet called by XConf.
// use AllConfigInterface.ApplyOption to modify the updated cc
// if passed in cc not valid, then return false, cc will not set to atomicAllConfig
var onAtomicAllConfigSet func(cc AllConfigInterface) bool

// InstallCallbackOnAtomicAllConfigSet install callback
func InstallCallbackOnAtomicAllConfigSet(callback func(cc AllConfigInterface) bool) {
	onAtomicAllConfigSet = callback
}

// AtomicAllConfigSet atomic setter for *AllConfig
func AtomicAllConfigSet(update interface{}) {
	cc := update.(*AllConfig)
	if onAtomicAllConfigSet != nil && !onAtomicAllConfigSet(cc) {
		return
	}
	atomic.StorePointer(&atomicAllConfig, (unsafe.Pointer)(cc))
}

// AtomicAllConfig return atomic *AllConfigVisitor
func AtomicAllConfig() AllConfigVisitor {
	current := (*AllConfig)(atomic.LoadPointer(&atomicAllConfig))
	if current == nil {
		defaultOne := newDefaultAllConfig()
		if watchDogAllConfig != nil {
			watchDogAllConfig(defaultOne)
		}
		atomic.CompareAndSwapPointer(&atomicAllConfig, nil, (unsafe.Pointer)(defaultOne))
		return (*AllConfig)(atomic.LoadPointer(&atomicAllConfig))
	}
	return current
}

// all getter func
func (cc *AllConfig) GetTypeBool() bool                          { return cc.TypeBool }
func (cc *AllConfig) GetTypeString() string                      { return cc.TypeString }
func (cc *AllConfig) GetTypeDuration() time.Duration             { return cc.TypeDuration }
func (cc *AllConfig) GetTypeFloat32() float32                    { return cc.TypeFloat32 }
func (cc *AllConfig) GetTypeFloat64() float32                    { return cc.TypeFloat64 }
func (cc *AllConfig) GetTypeInt() int                            { return cc.TypeInt }
func (cc *AllConfig) GetTypeUint() int                           { return cc.TypeUint }
func (cc *AllConfig) GetTypeInt8() int8                          { return cc.TypeInt8 }
func (cc *AllConfig) GetTypeUint8() uint8                        { return cc.TypeUint8 }
func (cc *AllConfig) GetTypeInt16() int16                        { return cc.TypeInt16 }
func (cc *AllConfig) GetTypeUint16() uint16                      { return cc.TypeUint16 }
func (cc *AllConfig) GetTypeInt32() int32                        { return cc.TypeInt32 }
func (cc *AllConfig) GetTypeUint32() uint32                      { return cc.TypeUint32 }
func (cc *AllConfig) GetTypeInt64() int64                        { return cc.TypeInt64 }
func (cc *AllConfig) GetTypeUint64() uint64                      { return cc.TypeUint64 }
func (cc *AllConfig) GetTypeSliceInt() []int                     { return cc.TypeSliceInt }
func (cc *AllConfig) GetTypeSliceUint() []uint                   { return cc.TypeSliceUint }
func (cc *AllConfig) GetTypeSliceInt8() []int8                   { return cc.TypeSliceInt8 }
func (cc *AllConfig) GetTypeSliceUint8() []uint8                 { return cc.TypeSliceUint8 }
func (cc *AllConfig) GetTypeSliceInt16() []int16                 { return cc.TypeSliceInt16 }
func (cc *AllConfig) GetTypeSliceUin16() []uint16                { return cc.TypeSliceUin16 }
func (cc *AllConfig) GetTypeSliceInt32() []int32                 { return cc.TypeSliceInt32 }
func (cc *AllConfig) GetTypeSliceUint32() []uint32               { return cc.TypeSliceUint32 }
func (cc *AllConfig) GetTypeSliceInt64() []int64                 { return cc.TypeSliceInt64 }
func (cc *AllConfig) GetTypeSliceUint64() []uint64               { return cc.TypeSliceUint64 }
func (cc *AllConfig) GetTypeSliceString() []string               { return cc.TypeSliceString }
func (cc *AllConfig) GetTypeSliceFloat32() []float32             { return cc.TypeSliceFloat32 }
func (cc *AllConfig) GetTypeSliceFloat64() []float64             { return cc.TypeSliceFloat64 }
func (cc *AllConfig) GetTypeSliceDuratuon() []time.Duration      { return cc.TypeSliceDuratuon }
func (cc *AllConfig) GetTypeMapStringIntNotLeaf() map[string]int { return cc.TypeMapStringIntNotLeaf }
func (cc *AllConfig) GetTypeMapStringInt() map[string]int        { return cc.TypeMapStringInt }
func (cc *AllConfig) GetTypeMapIntString() map[int]string        { return cc.TypeMapIntString }
func (cc *AllConfig) GetTypeMapStringString() map[string]string  { return cc.TypeMapStringString }
func (cc *AllConfig) GetTypeMapIntInt() map[int]int              { return cc.TypeMapIntInt }
func (cc *AllConfig) GetTypeMapStringDuration() map[string]time.Duration {
	return cc.TypeMapStringDuration
}
func (cc *AllConfig) GetRedis() RedisVisitor        { return cc.Redis }
func (cc *AllConfig) GetETCD() *ETCD                { return cc.ETCD }
func (cc *AllConfig) GetTestInterface() interface{} { return cc.TestInterface }

// AllConfigVisitor visitor interface for AllConfig
type AllConfigVisitor interface {
	GetTypeBool() bool
	GetTypeString() string
	GetTypeDuration() time.Duration
	GetTypeFloat32() float32
	GetTypeFloat64() float32
	GetTypeInt() int
	GetTypeUint() int
	GetTypeInt8() int8
	GetTypeUint8() uint8
	GetTypeInt16() int16
	GetTypeUint16() uint16
	GetTypeInt32() int32
	GetTypeUint32() uint32
	GetTypeInt64() int64
	GetTypeUint64() uint64
	GetTypeSliceInt() []int
	GetTypeSliceUint() []uint
	GetTypeSliceInt8() []int8
	GetTypeSliceUint8() []uint8
	GetTypeSliceInt16() []int16
	GetTypeSliceUin16() []uint16
	GetTypeSliceInt32() []int32
	GetTypeSliceUint32() []uint32
	GetTypeSliceInt64() []int64
	GetTypeSliceUint64() []uint64
	GetTypeSliceString() []string
	GetTypeSliceFloat32() []float32
	GetTypeSliceFloat64() []float64
	GetTypeSliceDuratuon() []time.Duration
	GetTypeMapStringIntNotLeaf() map[string]int
	GetTypeMapStringInt() map[string]int
	GetTypeMapIntString() map[int]string
	GetTypeMapStringString() map[string]string
	GetTypeMapIntInt() map[int]int
	GetTypeMapStringDuration() map[string]time.Duration
	GetRedis() RedisVisitor
	GetETCD() *ETCD
	GetTestInterface() interface{}
}

// AllConfigInterface visitor + ApplyOption interface for AllConfig
type AllConfigInterface interface {
	AllConfigVisitor
	ApplyOption(...AllConfigOption) []AllConfigOption
}
