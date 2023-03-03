package common

import (
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
    result := GenerateRandomNumber()

    if result == 0 {
        t.Errorf("expected non-zero result, but got %d", result)
    }
}