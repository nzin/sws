package sws

import (
    "time"
)


type TimerEvent struct {
    triggertime time.Time
    repeat      time.Duration // if >=0
    next        *TimerEvent
    trigger     func()
}

var timerlist *TimerEvent



func TimerAddEvent(triggertime time.Time, repeat time.Duration, trigger func()) *TimerEvent {
    te := &TimerEvent{
        triggertime: triggertime,
        repeat:      repeat,
        trigger:     trigger,
        next:        nil,
    }
    if timerlist==nil {
        timerlist=te
        return te
    } else {
        if (timerlist.triggertime.After(te.triggertime)) {
            te.next=timerlist
            timerlist=te
            return te
        }
        e:=timerlist
        for (e.next!=nil) {
            if (e.next.triggertime.After(te.triggertime)) {
                break
            }
        }
        te.next=e.next
        e.next=te
    }
    return te
}



func TriggerEvents() {
    now:=time.Now()
    for (timerlist!=nil && timerlist.triggertime.Before(now)) {
        t:=timerlist
        timerlist=timerlist.next
        t.trigger()
        if (t.repeat>0) {
            TimerAddEvent(t.triggertime.Add(t.repeat),t.repeat,t.trigger)
        }
    }
}



func StopRepeat(te *TimerEvent) bool {
    if timerlist==te {
        timerlist=timerlist.next
        return true
    }
    e:=timerlist
    for (e.next!=nil) {
        if e.next==te {
            e.next=te.next
            return true
        }
        e=e.next
    }
    return false
}
