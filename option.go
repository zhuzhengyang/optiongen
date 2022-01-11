package optiongen

import (
	"fmt"
	"strings"
)

//go:generate optiongen --option_with_struct_name=false --new_func=NewTestConfig --xconf=true --empty_composite_nil=true --usage_tag_name=usage
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"OptionPrefix":         "",                           // @MethodComment(option func name prefix, like: With, WithRedis)
		"OptionWithStructName": false,                        // @MethodComment(should the option func with struct name?)
		"OptionReturnPrevious": true,                         // @MethodComment(生成的Option方法是否返回之前的标签数据)
		"NewFunc":              "",                           // @MethodComment(new function name)
		"NewFuncReturn":        string(NewFuncReturnPointer), // @MethodComment(valid data: pointer,interface,visitor)
		// annotation@Verbose(xconf="v")
		"Verbose":           false, // @MethodComment(Deprecated: use --debug instead)
		"UsageTagName":      "",    // @MethodComment(usage tag name,if not empty,will gen usage support for xconf/xflag)
		"EmptyCompositeNil": false, // @MethodComment(should empty slice or map to be nil? otherwise will be make XXXX,0 )
		"Debug":             false, // @MethodComment(debug will print more detail info)
		// annotation@XConf(xconf="xconf")
		"XConf":           false, // @MethodComment(should gen xconf tag support?)
		"XConfTrimPrefix": "",    // @MethodComment(if enable xconf tag, the tag value will trim prefix [XConfTrimPrefix])
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
