package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "add",
	Short: "short init",
	Long:  "long init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init cmd run begin")
		fmt.Println(
			// 原生获取
			cmd.Flags().Lookup("viper").Value,
			cmd.Flags().Lookup("author").Value,
			cmd.Flags().Lookup("config").Value,
			cmd.Flags().Lookup("license").Value,
			cmd.Parent().Flags().Lookup("source"))
		// viper绑定后从配置文件中获取
		fmt.Println(
			viper.GetString("author"),
			viper.GetString("license"))
		fmt.Println("init cmd run end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
