package concurrent

import "sync/atomic"

// AtomicInteger is a int32 wrapper fo atomic
type AtomicInt32 int32

// IncrementAndGet increment wrapped int32 with 1 and return new value.
func (i *AtomicInt32) IncrementAndGet() int32 {
	return atomic.AddInt32((*int32)(i), int32(1))
}

// GetAndIncrement increment wrapped int32 with 1 and return old value.
func (i *AtomicInt32) GetAndIncrement() int32 {
	ret := atomic.LoadInt32((*int32)(i))
	atomic.AddInt32((*int32)(i), int32(1))
	return ret
}

// DecrementAndGet decrement wrapped int32 with 1 and return new value.
func (i *AtomicInt32) DecrementAndGet() int32 {
	return atomic.AddInt32((*int32)(i), int32(-1))
}

// GetAndDecrement decrement wrapped int32 with 1 and return old value.
func (i *AtomicInt32) GetAndDecrement() int32 {
	ret := atomic.LoadInt32((*int32)(i))
	atomic.AddInt32((*int32)(i), int32(-1))
	return ret
}

// Get current value
func (i *AtomicInt32) Get() int32 {
	return atomic.LoadInt32((*int32)(i))
}

//set value
func (i *AtomicInt32) Set(val int32) {
	atomic.StoreInt32((*int32)(i), val)
}
