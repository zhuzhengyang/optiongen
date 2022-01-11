package optiongen

//go:generate optiongen --option_with_struct_name=false --new_func=NewTestConfig --xconf=true --empty_composite_nil=true --usage_tag_name=usage
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"OptionPrefix":         "",    // @MethodComment(option func name prefix, like: With, WithRedis)
		"OptionWithStructName": false, // @MethodComment(should the option func with struct name?)
		"OptionReturnPrevious": true,  // @MethodComment(生成的Option方法是否返回之前的标签数据)
		"NewFunc":              "",    // @MethodComment(new function name)
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

func init() {
	InstallConfigWatchDog(func(cc *Config) {
		if cc.Verbose {
			cc.Debug = true
		}
	})
}
