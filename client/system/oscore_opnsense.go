//go:build (linux && !android) || freebsd

package system

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func readOsCoreFile() (osName string, osVer string) {
	file, err := os.Open("/usr/local/opnsense/version/core")
	if err != nil {
		log.Warnf("failed to open file /usr/local/opnsense/version/core: %s", err)
		return "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "product_name") {
			osName = strings.TrimSpace(regexp.MustCompile("\"|,").ReplaceAllString(strings.Split(line, ":")[1], ""))
			continue
		}
		if strings.Contains(line, "product_version") {
			osVer = strings.TrimSpace(regexp.MustCompile("\"|,").ReplaceAllString(strings.Split(line, ":")[1], ""))
			continue
		}

		if osName != "" && osVer != "" {
			break
		}
	}
	return
}
