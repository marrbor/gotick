package gotick_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/marrbor/gotick"
	"github.com/stretchr/testify/assert"
)

var tick *gotick.Tick

func callback() {
	path := fmt.Sprintf("%s/%s.log", os.Getenv("HOME"), time.Now().Format("2006_01_02_15_04_05"))
	f, _ := os.Create(path)
	if f != nil {
		_ = f.Close()
	}
}

func TestNewTick(t *testing.T) {
	tick = gotick.NewTick(1*time.Second, callback)
	assert.NotNil(t, tick)
}

func TestTick_Start(t *testing.T) {
	err := tick.Start()
	assert.NoError(t, err)
	err = tick.Start()
	assert.Error(t, err)
}

func TestTick_Stop(t *testing.T) {
	time.Sleep(10 * time.Second)
	err := tick.Stop()
	assert.NoError(t, err)
	err = tick.Stop()
	assert.Error(t, err)
}

func TestTick_ChangeInterval(t *testing.T) {
	err := tick.Start()
	assert.NoError(t, err)
	time.Sleep(5 * time.Second)
	tick.ChangeInterval(2 * time.Second)
	time.Sleep(10 * time.Second)
	err = tick.Stop()
	assert.NoError(t, err)
}
