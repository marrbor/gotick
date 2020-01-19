package gotick

import "time"

const (
	// message of callCh
	StopMessage = "stop"

	// message of ctrlCh
	StopConfirmMessage = "stopped"
)

type (
	Tick struct {
		ticker   *time.Ticker
		interval time.Duration
		callback func()
		callCh   chan string // call control message
		respCh   chan string // response for control message
	}
)

func (t *Tick) run() {
	for {
		select {
		case <-t.ticker.C:
			t.callback()
		case msg := <-t.callCh:
			if msg == StopMessage {
				t.ticker.Stop()
				t.respCh <- StopConfirmMessage
			}
		}
	}
}

func (t *Tick) Start() {
	t.ticker = time.NewTicker(t.interval)
	go t.run()
}

func (t *Tick) Stop() {
	t.callCh <- StopMessage
	<-t.respCh // waiting for stop confirm message
}

func NewTick(interval time.Duration, f func()) *Tick {
	return &Tick{
		ticker:   nil,
		interval: interval,
		callback: f,
		callCh:   make(chan string),
		respCh:   make(chan string),
	}
}
