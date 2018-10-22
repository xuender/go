package cmd

import (
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/cobra"
)

var valueCmd = &cobra.Command{
	Use:   "value",
	Short: "反射测试",
	Long:  `反射输出测试.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("信息")
		arr := []int{1, 2, 3, 4}
		value := reflect.ValueOf(arr)
		t := reflect.TypeOf(arr)
		fmt.Println(t)
		fmt.Println(value.Type())

		ret := reflect.New(t).Elem()
		ret.Set(reflect.MakeSlice(t, 0, value.Len()))
		for i := 0; i < value.Len()-1; i++ {
			ret.SetLen(i + 1)
			ret.Index(i).Set(value.Index(i))
		}
		log.Println(ret)
		log.Println(reflect.DeepEqual(1, 1))
	},
}

func init() {
	rootCmd.AddCommand(valueCmd)
}
