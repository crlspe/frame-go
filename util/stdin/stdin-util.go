package stdin

import (
	"bufio"
	"fmt"
	"os"
)

func GetStdin() (string, error) {
	var pipedInput = ""
	var stat, err = os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			pipedInput += scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			os.Exit(1)
		}
	}

	return pipedInput, nil
}
