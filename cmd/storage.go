package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func writeCsv(taskDet [][]string) {
	file, err := os.OpenFile("tasks.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(taskDet)
	writer.Flush()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Saved in file")
}

func readCsv() [][]string {
	file, err := os.OpenFile("tasks.csv", os.O_RDONLY, 0222)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No stored tasks found. Please create a new task to get started")
		}
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	taskDet, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(taskDet)

	return taskDet
}
