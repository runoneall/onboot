package cli

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "onboot",
	Short: "onboot 是一个现代化的跨平台服务管理器",
}

func Init() *cobra.Command {
	initSubCmd()

	return Cmd
}
