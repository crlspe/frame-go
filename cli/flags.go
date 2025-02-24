package cli

type Flags struct {
    values *map[string]any
}

func (f Flags) Get(flagName string) any {
	return getBool(f.values, flagName)
}

func getBool(flags *map[string]any, flagName string) bool {
    if  val, ok := (*flags)[flagName]; ok {
        return *val.(*bool)
    }
    return false
}
