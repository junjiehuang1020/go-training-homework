package hystrix

import (
	"sync"
	"time"
)

type Bucket struct {
	sync.RWMutex

	Value     int
	Timestamp time.Time
}

func NewBucket() *Bucket {
	return &Bucket{
		Timestamp: time.Now(),
	}
}
