package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

type Task struct {
	id      int64
	content string
	created time.Time
	status  string
	due     string
}

func assignId() int64 {
	_, err := os.Stat("tasks.csv")

	if os.IsNotExist(err) {
		return 1
	} else {
		taskList := readCsv()
		if len(taskList) == 0 {
			return 1
		}
		lastTask := taskList[len(taskList)-1]
		id, _ := strconv.ParseInt(lastTask[0], 10, 32)
		return id + 1
	}
}

func addTask(t string, d string) {
	task := Task{
		id:      0,
		content: t,
		created: time.Now().UTC(),
		status:  "Incomplete",
		due:     d,
	}

	task.id = assignId()

	taskDet := [][]string{
		{
			strconv.FormatInt(task.id, 10),
			task.content,
			task.created.Format(time.DateTime),
			task.status,
			task.due,
		},
	}

	writeCsv(taskDet)
}

func showTasksCompact() {
	writer := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', 0)
	tasks := readCsv()

	if assignId() == 1 {
		fmt.Println("No created tasks. Please create a new task to display")
		os.Exit(1)
	}

	defer writer.Flush()

	fmt.Fprintln(writer, "ID\tTask\tCreated\tDue")
	count := 0
	for _, task := range tasks {
		created, _ := time.Parse(time.DateTime, task[2])
		due := timediff.TimeDiff(created)
		if strings.EqualFold(task[3], "Incomplete") {
			fmt.Fprintf(writer, "%v\t%v\t%v\t%v\n", task[0], task[1], due, task[4])
			count += 1
		}
	}

	if count == 0 {
		fmt.Fprintln(writer, "---\t---\t---\t---")
		fmt.Println("Looks like all your tasks are completed! Have some rest\n")
	}
}

func showTasksExtended() {
	writer := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', 0)
	tasks := readCsv()

	if assignId() == 1 {
		fmt.Println("No created tasks. Please create a new task to display")
		os.Exit(1)
	}

	defer writer.Flush()

	fmt.Fprintln(writer, "ID\tTask\tCreated\tStatus\tDue")
	for _, task := range tasks {
		created, _ := time.Parse(time.DateTime, task[2])
		due := timediff.TimeDiff(created)
		fmt.Fprintf(writer, "%v\t%v\t%v\t%v\t%v\n", task[0], task[1], due, task[3], task[4])
	}
}

func markComplete(id int64) {
	tasks := readCsv()
	t := [][]string{}

	for _, task := range tasks {
		v, _ := strconv.ParseInt(task[0], 10, 64)
		if v == id {
			task[3] = "Completed"
		}
		t = append(t, task)
	}

	err := os.Remove("tasks.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Task successfully marked as completed")
	writeCsv(t)
}

func delTask(id int64) {
	tasks := readCsv()
	t := [][]string{}

	for _, task := range tasks {
		v, _ := strconv.ParseInt(task[0], 10, 64)
		if v != id {
			t = append(t, task)
		}
	}

	err := os.Remove("tasks.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Task successfully deleted")
	writeCsv(t)
}
