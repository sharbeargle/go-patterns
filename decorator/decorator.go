package main

import (
	"fmt"
	"os"
)

type LoggerAPI interface {
	Write(string) (string, error)
}

type STDOUTLogger struct {
	L LoggerAPI
}

func (s *STDOUTLogger) Write(str string) (string, error) {
	if s.L != nil {
		str, _ = s.L.Write(str)
	}
	newStr := "STDOUT: " + str
	return newStr, nil
}

type FileLogger struct {
	L LoggerAPI
}

func (f *FileLogger) Write(str string) (string, error) {
	if f.L != nil {
		str, _ = f.L.Write(str)
	}
	newStr := "FILE: " + str
	return newStr, nil
}

type SysLogger struct {
	L LoggerAPI
}

func (s *SysLogger) Write(str string) (string, error) {
	if s.L != nil {
		str, _ = s.L.Write(str)
	}
	newStr := "SYSLOG: " + str
	return newStr, nil
}

type Decorator func(string)

func DecoratorFunc(d Decorator) Decorator {
	return func(str string) {
		fmt.Println("Before decorator run on", str)
		d(str)
		fmt.Println("After decorator run on", str)
	}
}

func main() {
	// Decorator interfaces
	l := &SysLogger{&FileLogger{&STDOUTLogger{}}}
	str := "This is a really important message!"
	result, err := l.Write(str)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(result)

	// Decorator functions
	var f Decorator = func(str string) {
		fmt.Printf("DECORATED[%s]\n", str)
	}
	n := DecoratorFunc(f)
	n("Tree")
}
