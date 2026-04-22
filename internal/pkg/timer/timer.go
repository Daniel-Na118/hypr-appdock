package timer

import (
	"sync"
	"time"
)

type Timer struct {
	timer     *time.Timer
	active    bool
	startTime time.Time
	mu        sync.Mutex
	handler   func()        // ?ек??ий handler (?ол?ко дл? ак?ивного ?айме?а)
	duration  time.Duration // ?ек??а? дли?ел?но??? (?ол?ко дл? ак?ивного ?айме?а)
}

func New() *Timer {
	return &Timer{}
}

func (t *Timer) Run(ms int, handler func()) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.stopUnsafe() // о??анавливаем п?ед?д??ий ?айме?, е?ли б?л

	t.active = true
	t.startTime = time.Now()
	t.handler = handler
	t.duration = time.Duration(ms) * time.Millisecond

	t.timer = time.AfterFunc(t.duration, func() {
		t.mu.Lock()
		defer t.mu.Unlock()

		if t.active {
			t.active = false
			t.handler()
		}
	})
}

func (t *Timer) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.stopUnsafe()
}

func (t *Timer) stopUnsafe() {
	if !t.active {
		return
	}

	t.timer.Stop()
	t.active = false
	// ??и?аем в?еменн?е данн?е
	t.handler = nil
	t.duration = 0
}

func (t *Timer) ExecIf(check func(elapsedMs int) bool) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.active {
		return false
	}

	elapsed := time.Since(t.startTime)
	if check(int(elapsed.Milliseconds())) {
		t.triggerUnsafe()
		return true
	}
	return false
}

func (t *Timer) ExecNow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.active {
		return false
	}

	t.triggerUnsafe()
	return true
}

func (t *Timer) IsRunning() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.active
}

func (t *Timer) triggerUnsafe() {
	if !t.active || t.handler == nil {
		return
	}

	t.timer.Stop()
	t.active = false
	handler := t.handler
	// ??и?аем в?еменн?е данн?е пе?ед в?зовом handler
	t.handler = nil
	t.duration = 0

	handler()
}

