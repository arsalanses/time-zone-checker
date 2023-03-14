/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package check

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// tzdateCmd represents the tzdate command
var tzdateCmd = &cobra.Command{
	Use:   "tzdate",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("tzdate called")

		out, err := exec.Command("bash", "-c", "dpkg -s tzdata | grep -i version").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The date is %s\n", out)
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	CheckCmd.AddCommand(tzdateCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tzdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tzdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
