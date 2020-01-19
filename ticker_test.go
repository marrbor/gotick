package gotick_test

import (
	"fmt"
	"github.com/marrbor/gotick"
	"os"
	"testing"
	"time"
)

var tick *gotick.Tick

func callback() {
	path := fmt.Sprintf("%s/%s.log", os.Getenv("HOME"), time.Now().Format("2006_01_02_15_04_05.log"))
	f, _ := os.Create(path)
	if f != nil {
		_ = f.Close()
	}
}

func TestNewTick(t *testing.T) {
	tick = gotick.NewTick(1*time.Second, callback)
}

func TestTick_Start(t *testing.T) {
	tick.Start()
}

func TestTick_Stop(t *testing.T) {
	time.Sleep(1 * time.Minute)
	tick.Stop()
}
