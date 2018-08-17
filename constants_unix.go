// +build linux darwin

package discord_rpc

import (
    "net"
    "os"
    "path"
)

func GetDiscordRpcPipe() (net.Conn, error) {
	return net.Dial("unix", path.Join(os.Getenv("XDG_RUNTIME_DIR"), "discord-ipc-0"))
}
