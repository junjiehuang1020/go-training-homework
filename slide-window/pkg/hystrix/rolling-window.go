package hystrix

import (
	"fmt"
	"sync"
)

type RollingWindow struct {
	sync.RWMutex

	size    int
	buckets []*Bucket
}

func NewBuckets(size int) *RollingWindow {
	return &RollingWindow{
		size:    size,
		buckets: make([]*Bucket, 0, size),
	}
}

func (r *RollingWindow) AppendBucket() {

	r.Lock()

	defer r.Unlock()

	r.buckets = append(r.buckets, NewBucket())
	if !(len(r.buckets) < r.size+1) {
		r.buckets = r.buckets[1:]
	}

}

func (r *RollingWindow) GetBucket() *Bucket {
	if len(r.buckets) == 0 {
		r.AppendBucket()
	}
	return r.buckets[len(r.buckets)-1]
}

func (r *RollingWindow) ShowAllBuckets() {
	for _, v := range r.buckets {
		fmt.Printf("| id: [%v] | value: [%d] |\n", v.Timestamp, v.Value)
	}
}

func (r *RollingWindow) Increment(num int) {
	if num == 0 {
		return
	}

	r.Lock()

	defer r.Unlock()

	b := r.GetBucket()
	b.Value += num

}
