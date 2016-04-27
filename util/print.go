package util

import (
	"encoding/json"
	"fmt"
	//"github.com/fatih/structs"
	//"github.com/olekukonko/tablewriter"
	//"os"
)

// Prints a struct as Json, each value on a new line.
// TODO: Add check for input type
func PrintJson(input ...interface{}) {
	b, _ := json.MarshalIndent(input, "", " ")
	fmt.Println(string(b))
}

func PrintTable(input ...interface{}) {
	fmt.Printf("%T", input)
    //fmt.Printf("%+v\n", input)
	//m := structs.Map(*input)
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Key", "Value"})
    //table.Append(input.ID, input.Name)

	//for k, v := range m {
	//	fmt.Println(k, v)
	//	//table.Append(k)
	//}

	//table.Render()
}
