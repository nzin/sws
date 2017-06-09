package sws

import (
	"time"
)


//
// The TimerEvent struct is an opaque object, you mainly
// use func TimerAddEvent() and StopRepeat() functions
//
// TimerEvent are used to trigger an callback at a later time.
// It can be one-shot or repeatable.
//
type TimerEvent struct {
	triggertime time.Time
	repeat      time.Duration // if >=0
	next        *TimerEvent
	trigger     func()
}

var timerlist *TimerEvent



//
// Main entry point to create a new timer (and place it into the event queue)
//
// If repeat>0, this event is repeatable (until stopped. See StopRepeat())
//
func TimerAddEvent(triggertime time.Time, repeat time.Duration, trigger func()) *TimerEvent {
	te := &TimerEvent{
		triggertime: triggertime,
		repeat:      repeat,
		trigger:     trigger,
		next:        nil,
	}
	placeEvent(te)
	return te
}

func placeEvent(te *TimerEvent) {
	if timerlist == nil {
		timerlist = te
		return
	} else {
		if (timerlist.triggertime.After(te.triggertime)) {
			te.next = timerlist
			timerlist = te
			return
		}
		e := timerlist
		for (e.next != nil) {
			if (e.next.triggertime.After(te.triggertime)) {
				break
			}
			e = e.next
		}
		te.next = e.next
		e.next = te
	}
}



//
// For now, the events are not running into their own thread, so
// we must pool regularly, if an event has to be trigger.
// You don't need to call this function normally. It is done
// in the main loop ( PoolEvent(bool) )
//
func TriggerEvents() {
	now := time.Now()
	for (timerlist != nil && timerlist.triggertime.Before(now)) {
		t := timerlist
		timerlist = timerlist.next
		t.trigger()
		if (t.repeat > 0) {
			t.triggertime = t.triggertime.Add(t.repeat)
			placeEvent(t)
		}
	}
}



//
// When you need to stop a repeatable event, call this function
//
func (te *TimerEvent) StopRepeat() bool {
	if (te == nil || timerlist == nil) {
		return false
	}
	if timerlist == te {
		timerlist = timerlist.next
		return true
	}
	e := timerlist
	for (e.next != nil) {
		if e.next == te {
			e.next = te.next
			return true
		}
		e = e.next
	}
	return false
}
