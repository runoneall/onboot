package cli

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "onboot",
	Short: "onboot 是一个简约的跨平台服务管理器",
}

func Init() *cobra.Command {
	initSubCmd()

	return Cmd
}
