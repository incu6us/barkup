package barkup

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Location is an `Exporter` interface that backs up a local file or directory
type Location struct {
	Path string
}

// Export produces a specified location or directory, and creates a gzip compressed tarball archive.
func (x Location) Export() *ExportResult {
	result := &ExportResult{MIME: "application/x-tar"}

	dumpPath := fmt.Sprintf(`%v_%v`, x.Path, time.Now().Unix())

	result.Path = dumpPath + ".tar.gz"
	out, err := exec.Command(TarCmd, "-czf", result.Path, dumpPath).Output()
	if err != nil {
		result.Error = makeErr(err, string(out))
		return result
	}
	os.Remove(dumpPath)

	return result
}
