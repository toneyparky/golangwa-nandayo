package main

import (
	"fmt"
	"github.com/jaeyeom/gogo/task"
	"sort"
	"time"
)

func ExampleTaskSorting() {
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

	// Output:
	// [[v] Breakfast 2016-08-16 15:43:00 +0000 UTC [ ] Wash Dish 2011-08-16 15:43:00 +0000 UTC [ ] Nap 2013-08-16 15:43:00 +0000 UTC [ ] Laundry 2012-08-16 15:43:00 +0000 UTC]
	// =======================
	// [[ ] Laundry 2012-08-16 15:43:00 +0000 UTC [ ] Nap 2013-08-16 15:43:00 +0000 UTC [ ] Wash Dish 2011-08-16 15:43:00 +0000 UTC [v] Breakfast 2016-08-16 15:43:00 +0000 UTC]
	// =======================
	// [[v] Breakfast 2016-08-16 15:43:00 +0000 UTC [ ] Wash Dish 2011-08-16 15:43:00 +0000 UTC [ ] Nap 2013-08-16 15:43:00 +0000 UTC [ ] Laundry 2012-08-16 15:43:00 +0000 UTC]
	// =======================
	// [[v] Breakfast 2016-08-16 15:43:00 +0000 UTC [ ] Nap 2013-08-16 15:43:00 +0000 UTC [ ] Laundry 2012-08-16 15:43:00 +0000 UTC [ ] Wash Dish 2011-08-16 15:43:00 +0000 UTC]
}
