package main

import (
	"fmt"
	"sort"
)

func main11() {
	startTimes := []int{461620201, 461620202, 461620203}
	backupDuration := []int{3, 1, 2}
	maxThreads := 2
	res := backupTimeEstimator(startTimes, backupDuration, maxThreads)
	for _, v := range res {
		fmt.Printf("res = %f\n", v)
	}
}

type Event struct {
	time     float64
	index    int
	evtType  int
	cpuIndex int
}

func (evt *Event) String() string {
	return fmt.Sprintf("time=%.2f, index=%d, evtType=%d", evt.time, evt.index, evt.evtType)
}

type EventByTime []*Event

func (a EventByTime) Len() int           { return len(a) }
func (a EventByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EventByTime) Less(i, j int) bool { return a[i].time < a[j].time }

func backupTimeEstimator(startTimes []int, backupDuration []int, maxThreads int) []float64 {
	cpu := make([]*Event, maxThreads)
	res := make([]float64, len(startTimes))
	events := make([]*Event, 0)
	for k, v := range startTimes {
		events = append(events, &Event{float64(v), k, 1, -1})
	}
	queue := make([]int, 0)
	cpuInUse := 0
	for len(events) > 0 {
		sort.Sort(EventByTime(events))
		evt := events[0]
		events = events[1:]
		t := evt.time
		if evt.evtType == 0 {
			res[evt.index] = evt.time
			if len(queue) > 0 {
				evt.index = queue[0]
				queue = queue[1:]
				evt.time = float64(t) + float64(backupDuration[evt.index]*cpuInUse)
				events = append(events, evt)
			} else {
				cpu[evt.cpuIndex] = nil
				updateWorkingTime(t, cpu, cpuInUse, cpuInUse-1)
				cpuInUse--
			}
		} else {
			if cpuInUse < maxThreads {
				updateWorkingTime(t, cpu, cpuInUse, cpuInUse+1)
				cpuInUse++
				for j := 0; j < len(cpu); j++ {
					if cpu[j] == nil {
						cpu[j] = evt
						// assign the job
						evt.time = float64(t) + float64(backupDuration[evt.index]*cpuInUse)
						evt.evtType = 0
						evt.cpuIndex = j
						break
					}
				}
				events = append(events, evt)
			} else {
				queue = append(queue, evt.index)
			}
		}
	}
	return res
}

func updateWorkingTime(t float64, cpu []*Event, oldCpu, newCpu int) {
	for j := 0; j < len(cpu); j++ {
		if cpu[j] != nil {
			evt := cpu[j]
			rest := float64(evt.time-t) / float64(oldCpu)
			evt.time = t + rest*float64(newCpu)
		}
	}
}
