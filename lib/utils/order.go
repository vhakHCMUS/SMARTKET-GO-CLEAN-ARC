package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOrderCode generates a unique order code
func GenerateOrderCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("ORD%d%04d", time.Now().Unix(), rand.Intn(10000))
}
