package cli

type Flags struct {
    values *map[string]any
}

func (f Flags) Get(flagName string) any {
    if val, ok := (*f.values)[flagName]; ok {
        return val
    }
    return nil
}
