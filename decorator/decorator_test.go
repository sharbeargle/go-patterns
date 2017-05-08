package main

import (
	"testing"
)

func TestSTDOUTLogger_Write(t *testing.T) {
	s := &STDOUTLogger{}
	str := "Important Message!"
	expectedResult := "STDOUT: Important Message!"
	result, err := s.Write(str)
	if err != nil {
		t.Error(err)
	}
	if result != expectedResult {
		t.Errorf("Unexpected Result\nExpected: %s\nResult: %s\n", expectedResult, result)
	}
}

func TestFileLogger_Write(t *testing.T) {
	f := &FileLogger{}
	str := "Important Message!"
	expectedResult := "FILE: Important Message!"
	result, err := f.Write(str)
	if err != nil {
		t.Error(err)
	}
	if result != expectedResult {
		t.Errorf("Unexpected Result\nExpected: %s\nResult: %s\n", expectedResult, result)
	}
}

func TestSysLogger_Write(t *testing.T) {
	s := &SysLogger{}
	str := "Important Message!"
	expectedResult := "SYSLOG: Important Message!"
	result, err := s.Write(str)
	if err != nil {
		t.Error(err)
	}
	if result != expectedResult {
		t.Errorf("Unexpected Result\nExpected: %s\nResult: %s\n", expectedResult, result)
	}
}

func TestFullStack(t *testing.T) {
	l := &SysLogger{&FileLogger{&STDOUTLogger{}}}
	str := "Important Message!"
	expectedResult := "SYSLOG: FILE: STDOUT: Important Message!"
	result, err := l.Write(str)
	if err != nil {
		t.Error(err)
	}
	if result != expectedResult {
		t.Errorf("Unexpected Result\nExpected: %s\nResult: %s\n", expectedResult, result)
	}
}
