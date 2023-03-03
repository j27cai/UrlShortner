package common

import (
	"math/rand"
)

// I know this is not completely random or when there are enough keys generated, there will be more collisions
// However, this can be replaced with a key generator service to ensure there are no collisions within the algorithm
// One way to guarantee no collisions is to have a consistent service sequentially generate strings using the alphanumeric alphabet 
// You would start with something like 0000000000 and then each digit would just increase from 0-z and end up with zzzzzzzzzz. 
// This would gurarantee at least 36^10 options which is a lot
// then hashing before distributing those keys elsewhere
func GenerateRandomNumber() int64 {
	return rand.Int63()
}