package main

import (
	"fmt"
	"github.com/jaeyeom/gogo/task"
)

func ExampleMarkDone() {
	tasks := task.Task{
		Title:    "Laundry",
		Status:   task.TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []task.Task{{
			Title:    "Wash",
			Status:   task.TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []task.Task{
				{"Put", task.DONE, nil, 2, nil},
				{"Detergent", task.TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   task.TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   task.TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}

	tasks.MarkDone()
	fmt.Println(task.IncludeSubTasks(tasks).String())

	// Output:
	// [v] Laundry <nil>
	//   [v] Wash <nil>
	//     [v] Put <nil>
	//     [v] Detergent <nil>
	//   [v] Dry <nil>
	//   [v] Fold <nil>

}
