package lang

import (
	"fmt"
	"log"
	"regexp"
)

func SanitizePuncuation(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(s, "")
	fmt.Printf("A string of %s becomes %s \n", s, processedString)
	return processedString
}
