package output

import (
	"fmt"
	"github.com/iancoleman/orderedmap"
	"strconv"
	"strings"
)

//func Keys(m map[string]any) []string {
//	result := make([]string, 0, len(m))
//	for k := range m {
//		result = append(result, k)
//	}
//	return result
//}

func PrintTable(data []any, cols []string, displayHeader bool) {
	var colLength []int
	var table [][]string
	var colCnt int

	if data == nil || len(data) == 0 {
		_ = fmt.Errorf("print of nil")
		return
	}

	if cols == nil {
		keys, ok := data[0].(orderedmap.OrderedMap)
		if ok {
			cols = keys.Keys()
		} else {
			return
		}
	}
	colLength = make([]int, len(cols))
	colCnt = len(cols)
	if displayHeader {
		table = make([][]string, len(data)+1)
		table[0] = make([]string, colCnt)
		for i, col := range cols {
			colName := fmt.Sprintf("%v", col)
			colName = strings.Replace(colName, "_", " ", -1)
			table[0][i] = strings.ToUpper(colName)
			colLength[i] = len(col)
		}
	} else {
		table = make([][]string, len(data))
	}

	indexHeader := 0
	if displayHeader {
		indexHeader = 1
	}

	for i, row := range data {
		mapRow := row.(orderedmap.OrderedMap)
		table[i+indexHeader] = make([]string, colCnt)
		for j, col := range cols {
			value, _ := mapRow.Get(col)
			if value == nil {
				value = ""
			}
			table[i+indexHeader][j] = fmt.Sprintf("%v", value)
			colLen := len(table[i+indexHeader][j])
			if colLength[j] < colLen {
				colLength[j] = colLen
			}
		}
	}
	for _, row := range table {
		for j, col := range row {
			format := "%-" + strconv.Itoa(colLength[j]+3) + "s"
			fmt.Printf(format, col)
		}
		fmt.Println()
	}
}

func PrintKeyValue(data *orderedmap.OrderedMap) {
	length := 0
	keys := data.Keys()
	for _, key := range keys {
		if length < len(key) {
			length = len(key)
		}
	}

	for _, key := range keys {
		value, _ := data.Get(key)
		if value == nil {
			value = ""
		}
		key = key + ":"
		fmt.Printf("%-"+strconv.Itoa(length+1)+"s   %v\n", key, value)
	}
}
