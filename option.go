package optiongen

//go:generate optiongen --option_with_struct_name=false --new_func=NewTestConfig --xconf=true --empty_composite_nil=true --usage_tag_name=usage
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"OptionWithStructName": false, // @MethodComment(should the option func with struct name?)
		"NewFunc":              "",    // @MethodComment(new function name)
		// annotation@XConf(xconf="xconf")
		"XConf": false, // @MethodComment(should gen xconf tag?)
		// annotation@Verbose(xconf="v")
		"Verbose":           false, // @MethodComment(Deprecated: use --debug instead)
		"UsageTagName":      "",    // @MethodComment(usage tag name)
		"EmptyCompositeNil": false, // @MethodComment(should empty slice or map to be nil default?)
		"Debug":             false, // @MethodComment(debug will print more detail info)
		"XConfTrimPrefix":   "",    // @MethodComment(生成xconf标签时自动trim前缀)
	}
}

func init() {
	InstallConfigWatchDog(func(cc *Config) {
		if cc.Verbose {
			cc.Debug = true
		}
	})
}
