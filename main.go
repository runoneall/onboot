package main

import (
	"fmt"
	"onboot/cli"
	"onboot/cmd_init"
	"onboot/cmd_self_install"
	"onboot/cmd_service"
	"os"
)

func main() {
	cli.InstallWithSystemd.Run = cmd_self_install.ProcessUseSystemd
	cli.InstallWithRcUpdate.Run = cmd_self_install.ProcessUseRcUpdate
	cli.InstallWithLaunchd.Run = cmd_self_install.ProcessUseLaunchd
	cli.InstallWithScExe.Run = cmd_self_install.ProcessUseScExe
	cli.InstallWithCronTab.Run = cmd_self_install.ProcessUseCronTab

	cli.InitCmd.Run = cmd_init.CallCmdInit
	cli.ServiceCmd.Run = cmd_service.CallCmdService

	if err := cli.Init().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
