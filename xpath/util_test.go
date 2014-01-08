package xpath

import "testing"
import "github.com/zhigangc/gokogiri/help"

func CheckXmlMemoryLeaks(t *testing.T) {
	help.LibxmlCleanUpParser()
	if !help.LibxmlCheckMemoryLeak() {
		t.Errorf("Memory leaks: %d!!!", help.LibxmlGetMemoryAllocation())
		help.LibxmlReportMemoryLeak()
	}
}
