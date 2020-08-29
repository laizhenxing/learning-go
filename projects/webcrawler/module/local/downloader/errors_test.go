package downloader

import (
	"testing"
	"webcrawler/errors"
)

func TestErrorGenError(t *testing.T)  {
	simpleErrMsg := "testing error"
	expectedErrType := errors.ERROR_TYPE_DOWNLOAD
	err := genError(simpleErrMsg)
	ce, ok := err.(errors.CrawlerError)
	if !ok {
		t.Fatalf("Inconsistent error type: expected: %T, actual: %T",
			errors.NewCrawlerError("", ""), err)
	}
	if ce.Type() != expectedErrType {
		t.Fatalf("Inconsistent error type stirng: expected: %q, actual: %q",
			expectedErrType, ce.Type())

	}
	// crawler error: [downloader error] testing error
	expectedErrMsg := "crawler error: [downloader error] " + simpleErrMsg
	if ce.Error() != expectedErrMsg {
		t.Fatalf("Inconsistent error message: expectd: %q, actual: %q",
			expectedErrMsg, ce.Error())
	}
}
