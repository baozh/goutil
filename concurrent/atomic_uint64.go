package concurrent

import "sync/atomic"
//this port of atomic function is not very good, see uber's outstanding work, refer to "go.uber.org/atomic"

// AtomicInteger is a uint64 wrapper fo atomic
type AtomicUint64 uint64

// IncrementAndGet increment wrapped uint64 with 1 and return new value.
func (i *AtomicUint64) IncrementAndGet() uint64 {
	return atomic.AddUint64((*uint64)(i), uint64(1))
}

// GetAndIncrement increment wrapped uint64 with 1 and return old value.
func (i *AtomicUint64) GetAndIncrement() uint64 {
	ret := atomic.LoadUint64((*uint64)(i))
	atomic.AddUint64((*uint64)(i), uint64(1))
	return ret
}

// Get current value
func (i *AtomicUint64) Get() uint64 {
	return atomic.LoadUint64((*uint64)(i))
}

//set value
func (i *AtomicUint64) Set(val uint64) {
	atomic.StoreUint64((*uint64)(i), val)
}

func (i *AtomicUint64) AddAndGet(num uint64) uint64 {
	return atomic.AddUint64((*uint64)(i), num)
}
