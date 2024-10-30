package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "short desc",
	Long:  "long desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		fmt.Println(
			// 原生获取
			cmd.Flags().Lookup("viper").Value,
			cmd.PersistentFlags().Lookup("author").Value,
			cmd.PersistentFlags().Lookup("config").Value,
			cmd.PersistentFlags().Lookup("license").Value,
			cmd.Flags().Lookup("source").Value)
		// viper绑定后从配置文件中获取
		fmt.Println(
			viper.GetString("author"),
			viper.GetString("license"))
		fmt.Println("root cmd run end")
	},
	TraverseChildren: true, // 是否将本地命令传递到下级结构中
}

var (
	cfgFile     string
	userLicense string
)

func init() {
	cobra.OnInitialize(initConfig)
	// 按名称接收命令行参数
	rootCmd.PersistentFlags().Bool("viper", true, "")
	// 指定flag缩写
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "")
	// 通过参数传递赋值到变量
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	// 通过参数传递赋值到变量并设定缩写
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "")
	// 添加本地标志
	rootCmd.Flags().StringP("source", "s", "", "")
	// viper 绑定配置
	err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	if err != nil {
		fmt.Println(err)
	}
	err = viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	if err != nil {
		fmt.Println(err)
	}
	viper.SetDefault("author", "default author")
	viper.SetDefault("license", "MIT")
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("use config file:", viper.ConfigFileUsed())
}
