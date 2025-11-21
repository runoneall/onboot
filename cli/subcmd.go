package cli

import "github.com/spf13/cobra"

var SelfInstallCmd = &cobra.Command{
	Use:   "self-install",
	Short: "安装 onboot 服务端到目标平台",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var InstallWithSystemd = &cobra.Command{
	Use:   "systemd",
	Short: "使用 systemd 为 linux 系统安装",
}

var InstallWithRcUpdate = &cobra.Command{
	Use:   "rc-update",
	Short: "使用 rc-update 为 linux 系统安装",
}

var InstallWithLaunchd = &cobra.Command{
	Use:   "launchd",
	Short: "使用 launchd 为 macos 系统安装",
}

var InstallWithScExe = &cobra.Command{
	Use:   "sc.exe",
	Short: "使用 sc.exe 为 windows 系统安装",
}

var InstallWithCronTab = &cobra.Command{
	Use:   "crontab",
	Short: "使用 crontab 为其他系统安装",
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化 onboot 服务端及加载用户服务",
}

var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "用户服务管理",
}

func initSubCmd() {
	SelfInstallCmd.AddCommand(
		InstallWithSystemd,
		InstallWithRcUpdate,
		InstallWithLaunchd,
		InstallWithScExe,
		InstallWithCronTab,
	)

	Cmd.AddCommand(
		SelfInstallCmd,
		InitCmd,
		ServiceCmd,
	)
}
