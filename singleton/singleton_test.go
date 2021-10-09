package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	assert := assert.New(t)
	singleInstance1 := GetSingleInstance(1)
	singleInstance2 := GetSingleInstance(2)
	assert.Equal(singleInstance1, singleInstance2)
}
