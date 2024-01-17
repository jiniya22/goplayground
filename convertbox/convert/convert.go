package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	tableName := "user"
	filepath := "../resources/data.csv"
	columNames, rows := parseCsv(filepath)
	sql := createSql(tableName, columNames, rows)
	fmt.Println(sql)
}

func parseCsv(filepath string) ([]string, [][]string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// reader := csv.NewReader(bufio.NewReader(file))
	reader := csv.NewReader(file)

	var rows [][]string
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		rows = append(rows, row)
	}
	return rows[0], rows[1:]
}

func createSql(tableName string, columnNames []string, rows [][]string) string {
	var builder strings.Builder
	prefix := insertStatementPrefix(tableName, columnNames)

	for _, row := range rows {
		builder.WriteString(prefix)
		for j, col := range row {
			if j == 0 {
				builder.WriteString("(")
			}
			builder.WriteString(getData(col))
			if j < len(row)-1 {
				builder.WriteString(",")
			} else {
				builder.WriteString(")")
			}
		}
		builder.WriteString(";\n")
	}

	return builder.String()
}

func insertStatementPrefix(tableName string, columnNames []string) string {
	return fmt.Sprintf("INSERT INTO `%s` (%v) values", tableName, strings.Join(columnNames, ","))
}

func getData(s string) string {
	if IsInt(s) || IsFloat(s) {
		return fmt.Sprintf("%s", s)
	} else if s == "NULL" {
		return "NULL"
	}
	return fmt.Sprintf("\"%s\"", s)
}

func IsInt(s string) bool {
	return regexp.MustCompile("^\\d+$").MatchString(s)
}

func IsFloat(s string) bool {
	return regexp.MustCompile("^\\d*[.]\\d*$").MatchString(s)
}
