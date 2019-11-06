package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var template string
var constesllationfile string
var mapsfile string
var outputPath string

var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "compose helm package",
	Long: `compose is for composing helm package based on mapfile and constesllationfile.
	abstrakt compose -t [helm] -f [constesllationfile] -m [mapsfile] -o [outputPath]`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args: " + strings.Join(args, " "))
		fmt.Println("template: " + template)
		fmt.Println("constesllationfile: " + constesllationfile)
		fmt.Println("mapsfile: " + mapsfile)
		fmt.Println("outputPath: " + outputPath)
	},
}

func init() {
	composeCmd.Flags().StringVarP(&constesllationfile, "constesllationfile", "f", "", "constesllationfile path")
	composeCmd.MarkFlagRequired("constesllationfile")
	composeCmd.Flags().StringVarP(&mapsfile, "mapsfile", "m", "", "mapsfile path")
	composeCmd.MarkFlagRequired("mapsfile")
	composeCmd.Flags().StringVarP(&outputPath, "outputPath", "o", "", "destination directory")
	composeCmd.MarkFlagRequired("outputPath")
	composeCmd.Flags().StringVarP(&template, "template", "t", "helm", "output template")
}
