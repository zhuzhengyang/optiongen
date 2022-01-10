// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optiongen

package optiongen

import (
	"sync/atomic"
	"unsafe"
)

// Config struct
type Config struct {
	OptionWithStructName bool   `xconf:"option_with_struct_name" usage:"should the option func with struct name?"`
	NewFunc              string `xconf:"new_func" usage:"new function name"`
	// annotation@XConf(xconf="xconf")
	XConf bool `xconf:"xconf" usage:"should gen xconf tag?"`
	// annotation@Verbose(xconf="v")
	Verbose           bool   `xconf:"v" usage:"Deprecated: use --debug instead"`
	UsageTagName      string `xconf:"usage_tag_name" usage:"usage tag name"`
	EmptyCompositeNil bool   `xconf:"empty_composite_nil" usage:"should empty slice or map to be nil default?"`
	Debug             bool   `xconf:"debug" usage:"debug will print more detail info"`
	XConfTrimPrefix   string `xconf:"x_conf_trim_prefix" usage:"生成xconf标签时自动trim前缀"`
}

// SetOption apply single option
// Deprecated: use ApplyOption instead
func (cc *Config) SetOption(opt ConfigOption) {
	cc.ApplyOption(opt)
}

// ApplyOption apply mutiple new option and return the old mutiple optuons
// sample:
// old := cc.ApplyOption(WithTimeout(time.Second))
// defer cc.ApplyOption(old...)
func (cc *Config) ApplyOption(opts ...ConfigOption) []ConfigOption {
	var previous []ConfigOption
	for _, opt := range opts {
		previous = append(previous, opt(cc))
	}
	return previous
}

// GetSetOption apply new option and return the old optuon
// sample:
// old := cc.GetSetOption(WithTimeout(time.Second))
// defer cc.SetOption(old)
// Deprecated: use ApplyOption instead
func (cc *Config) GetSetOption(opt ConfigOption) ConfigOption {
	return opt(cc)
}

// ConfigOption option func
type ConfigOption func(cc *Config) ConfigOption

// should the option func with struct name?
// WithOptionWithStructName option func for OptionWithStructName
func WithOptionWithStructName(v bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.OptionWithStructName
		cc.OptionWithStructName = v
		return WithOptionWithStructName(previous)
	}
}

// new function name
// WithNewFunc option func for NewFunc
func WithNewFunc(v string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.NewFunc
		cc.NewFunc = v
		return WithNewFunc(previous)
	}
}

// should gen xconf tag?
// WithXConf option func for XConf
func WithXConf(v bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.XConf
		cc.XConf = v
		return WithXConf(previous)
	}
}

// Deprecated: use --debug instead
// WithVerbose option func for Verbose
func WithVerbose(v bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.Verbose
		cc.Verbose = v
		return WithVerbose(previous)
	}
}

// usage tag name
// WithUsageTagName option func for UsageTagName
func WithUsageTagName(v string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.UsageTagName
		cc.UsageTagName = v
		return WithUsageTagName(previous)
	}
}

// should empty slice or map to be nil default?
// WithEmptyCompositeNil option func for EmptyCompositeNil
func WithEmptyCompositeNil(v bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.EmptyCompositeNil
		cc.EmptyCompositeNil = v
		return WithEmptyCompositeNil(previous)
	}
}

// debug will print more detail info
// WithDebug option func for Debug
func WithDebug(v bool) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.Debug
		cc.Debug = v
		return WithDebug(previous)
	}
}

// 生成xconf标签时自动trim前缀
// WithXConfTrimPrefix option func for XConfTrimPrefix
func WithXConfTrimPrefix(v string) ConfigOption {
	return func(cc *Config) ConfigOption {
		previous := cc.XConfTrimPrefix
		cc.XConfTrimPrefix = v
		return WithXConfTrimPrefix(previous)
	}
}

// NewTestConfig(opts... ConfigOption) new Config
func NewTestConfig(opts ...ConfigOption) *Config {
	cc := newDefaultConfig()

	for _, opt := range opts {
		_ = opt(cc)
	}
	if watchDogConfig != nil {
		watchDogConfig(cc)
	}
	return cc
}

// InstallConfigWatchDog the installed func will called when NewTestConfig(opts... ConfigOption)  called
func InstallConfigWatchDog(dog func(cc *Config)) {
	watchDogConfig = dog
}

// watchDogConfig global watch dog
var watchDogConfig func(cc *Config)

// newDefaultConfig new default Config
func newDefaultConfig() *Config {
	cc := &Config{}

	for _, opt := range [...]ConfigOption{
		WithOptionWithStructName(false),
		WithNewFunc(""),
		WithXConf(false),
		WithVerbose(false),
		WithUsageTagName(""),
		WithEmptyCompositeNil(false),
		WithDebug(false),
		WithXConfTrimPrefix(""),
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
// GetOptionWithStructName return struct field: OptionWithStructName
func (cc *Config) GetOptionWithStructName() bool { return cc.OptionWithStructName }

// GetNewFunc return struct field: NewFunc
func (cc *Config) GetNewFunc() string { return cc.NewFunc }

// GetXConf return struct field: XConf
func (cc *Config) GetXConf() bool { return cc.XConf }

// GetVerbose return struct field: Verbose
func (cc *Config) GetVerbose() bool { return cc.Verbose }

// GetUsageTagName return struct field: UsageTagName
func (cc *Config) GetUsageTagName() string { return cc.UsageTagName }

// GetEmptyCompositeNil return struct field: EmptyCompositeNil
func (cc *Config) GetEmptyCompositeNil() bool { return cc.EmptyCompositeNil }

// GetDebug return struct field: Debug
func (cc *Config) GetDebug() bool { return cc.Debug }

// GetXConfTrimPrefix return struct field: XConfTrimPrefix
func (cc *Config) GetXConfTrimPrefix() string { return cc.XConfTrimPrefix }

// ConfigVisitor visitor interface for Config
type ConfigVisitor interface {
	GetOptionWithStructName() bool
	GetNewFunc() string
	GetXConf() bool
	GetVerbose() bool
	GetUsageTagName() string
	GetEmptyCompositeNil() bool
	GetDebug() bool
	GetXConfTrimPrefix() string
}

type ConfigInterface interface {
	ConfigVisitor
	ApplyOption(...ConfigOption) []ConfigOption
}
