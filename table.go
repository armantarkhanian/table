package table

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Print(tableData []map[string]string) {
	str := ""
	for i := 0; i < len(tableData); i++ {
		for key := range tableData[i] {
			if strings.Contains(str, key) {
				continue
			} else {
				str += key + "&"
			}
		}
	}
	collumns := strings.Split(str, "&")
	collumns = collumns[:len(collumns)-1]

	for i := 0; i < len(tableData); i++ {
		for j := 0; j < len(collumns); j++ {
			collumn := strings.TrimSpace(collumns[j])
			if utf8.RuneCountInString(tableData[i][collumn]) > utf8.RuneCountInString(collumns[j]) {
				collumns[j] += strings.Repeat(" ", utf8.RuneCountInString(tableData[i][collumn])-utf8.RuneCountInString(collumns[j]))
			}
		}
	}

	var line string
	for i := 0; i < len(collumns); i++ {
		if i+1 == len(collumns) {
			line += fmt.Sprintf("+%s+", strings.Repeat("-", utf8.RuneCountInString(collumns[i])+2))
		} else {
			line += fmt.Sprintf("+%s", strings.Repeat("-", utf8.RuneCountInString(collumns[i])+2))
		}
	}

	fmt.Println(line)
	for i := 0; i < len(collumns); i++ {
		if i+1 == len(collumns) {
			fmt.Printf("| %s |\n", collumns[i])
		} else {
			fmt.Printf("| %s ", collumns[i])
		}
	}
	fmt.Println(line)

	for i := 0; i < len(tableData); i++ {
		tr := ""
		for j := 0; j < len(collumns); j++ {
			collumn := strings.TrimSpace(collumns[j])
			if utf8.RuneCountInString(tableData[i][collumn]) < utf8.RuneCountInString(collumns[j]) {
				tableData[i][collumn] += strings.Repeat(" ", utf8.RuneCountInString(collumns[j])-utf8.RuneCountInString(tableData[i][collumn]))
			}
			if j+1 == len(collumns) {
				tr += fmt.Sprintf("| %s |\n", tableData[i][collumn])
			} else {
				tr += fmt.Sprintf("| %s ", tableData[i][collumn])
			}
		}
		fmt.Print(tr)
		fmt.Println(line)
	}
}
