package cmd_init

import (
	"fmt"
	"log"
	"net"
	"onboot/globalvar"
	"os"
	"time"

	"github.com/runoneall/pgoipc/ipcdial"
	"github.com/runoneall/pgoipc/ipcserver"
	"github.com/spf13/cobra"
)

var logger = log.New(
	os.Stdout,
	fmt.Sprintf("[%s] (%s) ", "BACKEND", time.Now().Format("2006-01-02 15:04:05")),
	0,
)

var ipcname = globalvar.ONBOOT_BACKEND_IPCNAME

func CallCmdInit(cmd *cobra.Command, args []string) {
	if conn, err := ipcdial.Dial(ipcname); err == nil {
		logger.Fatalln("backend already running:", conn.RemoteAddr())
	}

	fmt.Println("onboot backend serv at ipcname:", ipcname)

	ipcserver.Serv(ipcname, func(conn net.Conn) {
		logger.Println("new client connected:", conn.RemoteAddr())
	})
}
