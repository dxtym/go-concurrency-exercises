//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
	mu sync.RWMutex
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	checkForPoor := func() bool {
		u.mu.RLock()
		defer u.mu.RUnlock()
		if u.TimeUsed > 10 && !u.IsPremium {
			return false
		}
		return true
	}
	
	if !checkForPoor() {
		return false
	}
	
	start := time.Now()
	process()
	
	u.mu.Lock()
	u.TimeUsed += int64(time.Since(start).Seconds())
	u.mu.Unlock()
	
	return checkForPoor()	
}

func main() {
	RunMockServer()
}
