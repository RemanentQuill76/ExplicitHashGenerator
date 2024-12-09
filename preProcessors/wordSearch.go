package preprocessors
//wordSearch searches for valid word substrings within each word, and adds a
//space that the mask processor will create permutations within and around.

import (
	"strings"
	"fmt"
	"time"
	"strconv"
	"bufio"
	"os"
	jsoniter "github.com/json-iterator/go"
)

func searchForWords(input string) string {
	actions := map[string]bool{}
	var arr []string
	for i := 0; i < len(input); i++ {
		if strings.Contains(input, "L") {
		}
	}
	return arr
}

