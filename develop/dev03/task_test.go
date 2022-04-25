package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type args struct {
	files []string
	flags []string
}

func TestTask(t *testing.T) {
	cases := []args{
		{
			files: []string{"test3.txt"},
			flags: []string{"-k", "1"},
		},
		{
			files: []string{"test3.txt"},
			flags: []string{"-k", "1", "-r"},
		},
		{
			files: []string{"test.txt", "test2.txt"},
			flags: []string{"-k", "1", "-n", "-r"},
		},
		{
			files: []string{"test.txt", "test2.txt"},
			flags: []string{"-k", "1", "-n", "-r", "-u"},
		},
		{
			files: []string{"test.txt"},
		},
		{
			files: []string{"test.txt", "test2.txt"},
		},
	}
	for _, testCase := range cases {
		command := []string{"run", "task.go"}
		command = append(command, testCase.flags...)
		command = append(command, testCase.files...)
		got, err := exec.Command("go", command...).CombinedOutput()
		if err != nil {
			log.Println("Starting test failed:", err)
		}
		out := strings.ReplaceAll(string(got), "\r", "")

		bash := make([]string, 0)
		bash = append(bash, testCase.flags...)
		bash = append(bash, testCase.files...)

		rOut, err := exec.Command("sort", bash...).CombinedOutput()
		if err != nil {
			log.Println("Starting test failed:")
		}
		realOut := strings.ReplaceAll(string(rOut), "\r", "")

		if realOut != out {
			t.Errorf("Should: %s\nGot: %s\n", realOut, out)
			os.Exit(1)
		}
	}
}
