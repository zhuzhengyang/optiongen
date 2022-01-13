package optiongen

import (
	"fmt"
	"strings"
)

//go:generate optiongen --option_with_struct_name=false --new_func=NewTestConfig --xconf=true --empty_composite_nil=true --usage_tag_name=usage
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"OptionPrefix":         "",                           // annotation@OptionPrefix(comment="option func name prefix, like: With, WithRedis")
		"OptionWithStructName": false,                        // annotation@OptionWithStructName(comment="should the option func with struct name?")
		"OptionReturnPrevious": true,                         // annotation@OptionReturnPrevious(comment="should option func return the previous ones?")
		"NewFunc":              "",                           // annotation@NewFunc(comment="new function name")
		"NewFuncReturn":        string(NewFuncReturnPointer), // annotation@NewFuncReturn(comment="valid data: pointer,interface,visitor")
		"Verbose":              false,                        // annotation@Verbose(xconf="v",deprecated="use --debug instead")
		"UsageTagName":         "",                           // annotation@UsageTagName(comment="usage tag name,if not empty,will gen usage support for xconf/xflag")
		"EmptyCompositeNil":    false,                        // annotation@EmptyCompositeNil(comment="should empty slice or map to be nil? otherwise will be make(XXXX,0)")
		"Debug":                false,                        // annotation@Debug(comment="debug will print more detail info")
		"XConf":                false,                        // annotation@XConf(xconf="xconf",comment="should gen xconf tag support?")
		"XConfTrimPrefix":      "",                           // annotation@XConfTrimPrefix(comment="if enable xconf tag, the tag value will trim prefix [XConfTrimPrefix]")
	}
}

const NewFuncReturnPointer = "pointer"
const NewFuncReturnVisitor = "visitor"
const NewFuncReturnInterface = "interface"

var newFuncReturnAs = []string{NewFuncReturnPointer, NewFuncReturnVisitor, NewFuncReturnInterface}

func fixConfig(cc *Config) {
	if cc.GetVerbose() {
		cc.ApplyOption(WithDebug(true))
	}
	if !containStringEqualFold(newFuncReturnAs, cc.GetNewFuncReturn()) {
		panic(fmt.Sprintf("new_func_return: %s, value not valid, valid list: %v", cc.GetNewFuncReturn(), newFuncReturnAs))
	}
}

func init() {
	InstallConfigWatchDog(fixConfig)
	InstallCallbackOnAtomicConfigSet(func(cc ConfigInterface) bool {
		fixConfig(cc.(*Config))
		return true
	})
}

func containStringEqualFold(s []string, v string) bool {
	for _, vv := range s {
		if strings.EqualFold(vv, v) {
			return true
		}
	}
	return false
}
