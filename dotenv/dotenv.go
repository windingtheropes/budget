package dotenv

import (
	"os"
	"strings"
)

func Init() {
	buf, err := os.ReadFile(".env")
	
	if err != nil {
		return
	}

	lines := strings.Split(string(buf), "\n")
	for i := 0; i < len(lines); i++ {
		kv := strings.Split(strings.Trim(lines[i], "\n\r "), "=")
		if len(kv) == 2 {
			os.Setenv(kv[0], kv[1])
		}
	}
}
