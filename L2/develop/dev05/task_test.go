package main

import (
	"log"
	"os/exec"
	"testing"
)

type testArgs struct {
	flags   []string
	pattern string
	file    string
}

func TestGrep(t *testing.T) {
	cases := []testArgs{
		{
			flags:   []string{"-v"},
			pattern: "given",
			file:    "test.txt",
		},
		{
			flags:   []string{"-A", "1", "-B", "2"},
			pattern: "line",
			file:    "test.txt",
		},
		{
			flags:   []string{"-C", "3", "-i"},
			pattern: "line",
			file:    "test.txt",
		},
		{
			flags:   []string{"-c", "-v"},
			pattern: "line",
			file:    "test.txt",
		},
		{
			flags:   []string{"-n"},
			pattern: "line",
			file:    "test.txt",
		},
		{
			flags:   []string{"-A", "1", "-B", "2", "-F"},
			pattern: "line that matches a pattern.",
			file:    "test.txt",
		},
	}
	for _, testCase := range cases {
		command := []string{"run", "task.go"}
		command = append(command, testCase.flags...)
		command = append(command, testCase.pattern, testCase.file)

		out, err := exec.Command("go", command...).CombinedOutput()
		if err != nil {
			log.Println(err)
		}
		bash := append([]string{}, testCase.flags...)
		bash = append(bash, testCase.pattern, testCase.file)
		realOut, err := exec.Command("grep", bash...).CombinedOutput()
		if err != nil {
			log.Println(err)
		}
		if string(out) != string(realOut) {
			t.Errorf("Got: %s \n Should: %s", out, realOut)
		}
	}
}
