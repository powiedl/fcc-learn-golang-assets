package main


// #region CH13.1
// #region do not change
type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}
// #endregion

// #region change here
func (sc safeCounter) inc(key string) {
	// ?
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	// ?
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// #endregion

// #region do not change 
func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}
// #endregion
// #endregion

// CH13.2 - mutex refers to Mutual Exclusion
// C13.4 - to safely access shared ressources concurrently

// #region CH13.5
func (sc safeCounter) val(key string) int {
	sc.mu.RLock() 																				// added R 
	defer sc.mu.RUnlock()																	// added R
	return sc.counts[key]
}
// #endregion

// CH13.6 - one writer can access a RWMutex at once
// CH13.7 - infinite readers can access a RWMutex at once
// CH13.8 - no, readers and writers can not lock a RWMutex at once