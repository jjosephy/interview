package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"sync"
)

// Util type for various utility methods
type Util struct {
}

// InstanceUtil singleton instance
var InstanceUtil *Util

// NewUtil initializes the singleton instance of Util
func NewUtil() {
	if InstanceUtil == nil {
		var mutex = &sync.Mutex{}
		mutex.Lock()
		if InstanceUtil == nil {
			InstanceUtil = &Util{}
		}
		mutex.Unlock()
	}
}

// NewUUID generates a random UUID according to RFC 4122
func (u *Util) NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
