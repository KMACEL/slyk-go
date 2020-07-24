package slyk

// merge is
func merge(m interface{}) map[string]string {
	mmap := make(map[string]string)
	switch getMap := m.(type) {
	case []*getUserFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	}
	return mmap
}
