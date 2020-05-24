package lxcfs

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var (
	//Whether to enable lxcfs
	IsLxcfsEnabled bool

	//the absolute path of lxcfs
	LxcfsHomeDir string

	//file list of /proc
	LxcfsProcFiles = []string{"uptime", "swaps", "stat", "diskstats", "meminfo", "cpuinfo"}
)

// CheckLxcfsMount check if the the mount point of lxcfs exists
func CheckLxcfsMount() error {
	isMount := false
	f, err := os.Open("/proc/1/mountinfo")
	if err != nil {
		return fmt.Errorf("Check lxcfs mounts failed: %v", err)
	}
	fr := bufio.NewReader(f)
	for {
		line, err := fr.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("Check lxcfs mounts failed: %v", err)
		}

		if bytes.Contains(line, []byte(LxcfsHomeDir)) {
			isMount = true
			break
		}
	}
	if !isMount {
		return fmt.Errorf("%s is not a mount point, please run \" lxcfs %s \" before Pouchd", LxcfsHomeDir, LxcfsHomeDir)
	}
	return nil
}
