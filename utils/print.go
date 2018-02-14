package utils

import (
	"fmt"

	"github.com/chrislusf/gleam/util"
)

// PrintRow prints a row with tab seperated columns
func PrintRow(row *util.Row) error {
	fmt.Printf("\n\n%v\t", row.K[0])
	for _, v := range row.V {
		fmt.Printf("%v\t", v)
	}
	return nil
}
