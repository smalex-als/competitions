package test

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func RunTest(name string, t *testing.T) {
	cmdBuild := exec.Command("go", "build", name+".go")
	err := cmdBuild.Run()
	if err != nil {
		t.Errorf("build failed")
	}

	file, err := os.Open(name + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	start := time.Now()
	defer os.Remove(name)
	defer log.Printf("%s: %.05f s", name, float64(time.Since(start).Nanoseconds())/1000000000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "__IN__" {
			break
		}
	}
	i := 0
	for scanner.Scan() {
		var buf bytes.Buffer
		buf.WriteString(scanner.Text())
		buf.WriteString("\n")
		for scanner.Scan() {
			if scanner.Text() == "__OUT__" {
				break
			}
			buf.WriteString(scanner.Text())
			buf.WriteString("\n")
		}
		input := buf.String()
		buf.Reset()
		for scanner.Scan() {
			line := scanner.Text()
			if line == "__IN__" {
				break
			}
			buf.WriteString(scanner.Text())
			buf.WriteString("\n")
		}
		output := strings.TrimRight(buf.String(), "\n")
		cmd := exec.Command("./" + name)
		cmd.Stdin = strings.NewReader(input)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			t.Errorf("test number %v failed %q ", i+1, err)
		}
		res := strings.TrimRight(out.String(), "\n")
		if res != output {
			t.Fatalf("TEST %d\nOUTPUT\n%s\nINPUT\n%s"+
				"EXPECTED\n%s\n\n", i+1, res, input, output)
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
