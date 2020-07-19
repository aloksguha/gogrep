package gogrep

import (
	"fmt"
	"testing"
)

func TestGrepForInvalidFilePath(t *testing.T) {
	fPath := ""
	timeout := 60
	q := "Hi"
	now := 10

	grep := NewSearch(fPath, timeout, q, now)
	_, err := grep.search()
	if err == nil {
		t.Errorf("should retun error for invalid file path")
	}
}

func TestGrepForInvalidReportObjectLength(t *testing.T) {
	fPath := "../test_files/text.file"
	timeout := 60
	q := "Hi"
	now := 10

	grep := NewSearch(fPath,timeout,q, now)
	reports, err := grep.search()
	if err != nil {
		t.Errorf("should retun error for invalid file path")
	}
	fmt.Println(reports)
	if len(reports) != now {
			t.Errorf("Report should be retured for all workers FOUND : %d, SHOULD BE : %d", len(reports), now)
	}


}

func TestGrepReportObjectCreation(t *testing.T) {
	fPath := "../test_files/text.file"
	timeout := 60
	q := "Hi"
	now := 10

	grep := NewSearch(fPath,timeout,q, now)
	reports, err := grep.search()
	if err != nil {
		t.Errorf("should retun error for invalid file path")
	}

	for index:=0; index < len(reports); index++ {
		report := reports[index]
		if report.Status != STATUS(SUCEESS) {
			if report.ByteCnt != 0 {
				t.Errorf("ByteCnt should be 0 if report status is not %s", SUCEESS)
			}
			if report.Elapsed != 0 {
				t.Errorf("Elapsed time should be 0 if report status is not %s", SUCEESS)
			}
			if report.Remaining != 0 {
				t.Errorf("Remaining time should be 0 if report status is not %s", SUCEESS)
			}
		}

		if report.Status == STATUS(SUCEESS) {
			if report.ByteCnt == 0 {
				t.Errorf("ByteCnt should be 0 if report status is  %s", SUCEESS)
			}
			if report.Elapsed == 0 {
				t.Errorf("Elapsed time should be 0 if report status is  %s", SUCEESS)
			}
			if report.Remaining == 0 {
				t.Errorf("Remaining time should be 0 if report status is  %s", SUCEESS)
			}
		}
	}
}