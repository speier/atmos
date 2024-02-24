package cmd

import (
	"github.com/spf13/cobra"
)

// describeCmd describes configuration for stacks and components
var describeCmd = &cobra.Command{
	Use:                "describe",
	Short:              "Execute 'describe' commands",
	Long:               `This command shows configuration for CLI, stacks and components`,
	FParseErrWhitelist: struct{ UnknownFlags bool }{UnknownFlags: false},
}

func init() {
	describeCmd.PersistentFlags().String("jsonpath", "", "JSONPath query to filter and query the command output. https://goessner.net/articles/JsonPath")
	describeCmd.PersistentFlags().String("jmespath", "", "JMESPath query to filter and query the command output. https://jmespath.org  https://jmespath.site/main")

	RootCmd.AddCommand(describeCmd)
}
