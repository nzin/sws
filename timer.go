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
    placeEvent(te)
    return te
}



func placeEvent(te *TimerEvent) {
    if timerlist==nil {
        timerlist=te
        return
    } else {
        if (timerlist.triggertime.After(te.triggertime)) {
            te.next=timerlist
            timerlist=te
            return
        }
        e:=timerlist
        for (e.next!=nil) {
            if (e.next.triggertime.After(te.triggertime)) {
                break
            }
            e=e.next
        }
        te.next=e.next
        e.next=te
    }
}



func TriggerEvents() {
    now:=time.Now()
    for (timerlist!=nil && timerlist.triggertime.Before(now)) {
        t:=timerlist
        timerlist=timerlist.next
        t.trigger()
        if (t.repeat>0) {
            t.triggertime=t.triggertime.Add(t.repeat)
            placeEvent(t)
        }
    }
}



func StopRepeat(te *TimerEvent) bool {
    if (te==nil || timerlist==nil) { return false }
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
