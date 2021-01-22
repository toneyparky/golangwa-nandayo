package main

import (
	"fmt"
	"github.com/jaeyeom/gogo/task"
	"sort"
	"time"
)

type PrioritySorting []task.Task

func (p PrioritySorting) Len() int {
	return len(p)
}

func (p PrioritySorting) Less(i, j int) bool {
	return p[i].Priority < p[j].Priority
}

func (p PrioritySorting) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type DeadLineSorting []task.Task

func (d DeadLineSorting) Len() int {
	return len(d)
}

func (d DeadLineSorting) Less(i, j int) bool {
	//return d[i].Deadline.Sub(d[j].Deadline.Time) > 0 이래도 동작한다.
	return d[i].Deadline.Time.Sub(d[j].Deadline.Time) > 0
}

func (d DeadLineSorting) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	tasks := []task.Task{
		task.Task{
			Title:    "Breakfast",
			Status:   task.DONE,
			Deadline: task.NewDeadline(time.Date(2016, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 0,
			SubTasks: nil,
		},
		task.Task{
			Title:    "Laundry",
			Status:   task.TODO,
			Deadline: task.NewDeadline(time.Date(2012, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 4,
			SubTasks: nil,
		},
		task.Task{
			Title:    "Wash Dish",
			Status:   task.TODO,
			Deadline: task.NewDeadline(time.Date(2011, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 1,
			SubTasks: nil,
		},
		task.Task{
			Title:    "Nap",
			Status:   task.TODO,
			Deadline: task.NewDeadline(time.Date(2013, time.August, 16, 15, 43, 0, 0, time.UTC)),
			Priority: 3,
			SubTasks: nil,
		},
	}

	priorityTasks := PrioritySorting(tasks)

	sort.Sort(priorityTasks)
	fmt.Println(priorityTasks)

	fmt.Println("=======================")

	// 검색하다 발견한 sort의 방향을 바꾸는 방법
	sort.Sort(sort.Reverse(priorityTasks))
	fmt.Println(priorityTasks)

	fmt.Println("=======================")

	// 검색하다 발견한 Go 1.8에 추가된 방법.
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority < tasks[j].Priority
	})
	fmt.Println(tasks)

	fmt.Println("=======================")

	deadLineTasks := DeadLineSorting(tasks)

	sort.Sort(deadLineTasks)
	fmt.Println(deadLineTasks)
}
