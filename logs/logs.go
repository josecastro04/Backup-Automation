package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func ShowLogs(flags []string) {
	if len(flags) == 0 {
		flags = append(flags, "logs")
	}

	for _, flag := range flags {
		path := "../Logs/" + flag + ".log"
		data, err := os.ReadFile(path)
		if err != nil {
			SaveLogs(logrus.Error, fmt.Sprintf("Unable to read file: %s", path))
		}

		fmt.Printf("Logs (%s):\n%s", path, string(data))
	}
}

func SaveLogs(functionType func(args ...interface{}), string2 string) {
	functionType(string2)
}
