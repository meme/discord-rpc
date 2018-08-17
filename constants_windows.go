// +build windows

package discord_rpc

import (
	"github.com/Microsoft/go-winio"
	"time"
	"net"
)

func getDiscordRpcPipe() (net.Conn, error) {
	var timeout, _ = time.ParseDuration("5m")
	return winio.DialPipe("\\\\?\\pipe\\discord-ipc-0", &timeout)
}