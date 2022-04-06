package store

import "time"

type cacheDB interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
	Delete(k string)
	Flush()
}
