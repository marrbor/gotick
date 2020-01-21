package gotick

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func callback() {
	fmt.Println("callback")
}

func TestTick_ChangeInterval2(t *testing.T) {
	tick := NewTick(time.Second, callback)
	assert.NotNil(t, tick)
	assert.EqualValues(t, time.Second, tick.interval)
	assert.Nil(t, tick.ticker)
	tick.ChangeInterval(time.Minute)
	assert.EqualValues(t, time.Minute, tick.interval)
}
