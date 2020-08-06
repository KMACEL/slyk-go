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
	case []*getWalletActivityWithIDFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletBalanceWithIDFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletMovementFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletTransactionstFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletActivityFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getWalletBalanceFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*geTransactionstFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getaddressFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getassetFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getRateFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getPaymentMethodFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getMovementFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	case []*getInviteFilter:
		for i := range getMap {
			for k, v := range *getMap[i] {
				mmap[k] = v
			}
		}
	}
	return mmap
}
