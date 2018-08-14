package discord_rpc

import (
	"net"
	"encoding/json"
	"encoding/binary"
	"bytes"
	"time"
	"errors"
	"os"
)

const rpcSetActivityDirective = "SET_ACTIVITY"
const rpcErrorEvent = "ERROR"

type DiscordRpcClient struct {
	Conn net.Conn
}

func writeCommandToPipe(conn net.Conn, command string, arguments interface{}) (error) {
	err := writeToPipe(conn, 1, RpcCommand{
		Command:   command,
		Arguments: arguments,
		Nonce:     string(time.Now().Unix()),
	})

	if err != nil {
		return err
	}

	var response map[string]interface{}
	err = readFromPipe(conn, &response)

	if err != nil {
		return err
	}

	if response["evt"] == rpcErrorEvent {
		return errors.New(response["data"].(map[string]interface{})["message"].(string))
	}

	return nil
}

func writeToPipe(conn net.Conn, op uint32, payload interface{}) (error) {
	b, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	out := new(bytes.Buffer)
	binary.Write(out, binary.LittleEndian, op)
	binary.Write(out, binary.LittleEndian, uint32(len(b)))
	out.Write(b)

	_, err = out.WriteTo(conn)

	return err
}

func readFromPipe(conn net.Conn, v interface{}) (error) {
	var op, length uint32
	err := binary.Read(conn, binary.LittleEndian, &op)

	if err != nil {
		return err
	}

	err = binary.Read(conn, binary.LittleEndian, &length)

	if err != nil {
		return err
	}

	b := make([]byte, length)

	err = binary.Read(conn, binary.LittleEndian, &b)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func CreateRpcClient(clientId string) (*DiscordRpcClient, error) {
	conn, err := GetDiscordRpcPipe()

	if err != nil {
		return nil, err
	}

	err = writeToPipe(conn, 0, RpcHandshake{Version: 1, ClientId: clientId})

	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = readFromPipe(conn, &response)

	if err != nil {
		return nil, err
	}

	if response["code"] != nil {
		return nil, errors.New("failed to complete handshake")
	}

	return &DiscordRpcClient{Conn: conn}, nil
}

func (client *DiscordRpcClient) SetActivity(state, details string, largeImage, largeText string) (error) {
	return writeCommandToPipe(client.Conn, rpcSetActivityDirective, SetActivityCommand{
		ProcessId: os.Getpid(),
		Activity: &ActivityArgument{
			State:   state,
			Details: details,
			Assets: &AssetArgument{
				LargeImage: largeImage,
				LargeText:  largeText,
			},
			Instance: true,
		},
	})
}

func (client *DiscordRpcClient) ClearActivity() (error) {
	return writeCommandToPipe(client.Conn, rpcSetActivityDirective, &SetActivityCommand{
		ProcessId: os.Getpid(),
		Activity:  nil,
	})
}