package timefmt

import (
	"fmt"
	"testing"
)

var fmtData = []string{
	"2006-01-02",
	"2006-01",
	"2006-01-02 15:04:05",
	"2006-01-02 15",
	"2006-01-02 15:04",
}

func TestGetUnixTimeByFmt(t *testing.T) {
	for _, str := range fmtData {
		todayBegin := GetUnixTimeByFmt(str)
		fmt.Printf("%s => %d\n", str, todayBegin)
	}
}
