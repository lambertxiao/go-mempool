package mempool

type Locker interface {
	Lock()
	Unlock()
}
