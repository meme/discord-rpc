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

type Client struct {
	conn net.Conn
}

func writeCommandToPipe(conn net.Conn, command string, arguments interface{}) (error) {
	err := writeToPipe(conn, 1, rpcCommand{
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

func (client *Client) Connect(clientId string) (error) {
	conn, err := getDiscordRpcPipe()

	if err != nil {
		return err
	}

	err = writeToPipe(conn, 0, rpcHandshake{Version: 1, ClientId: clientId})

	if err != nil {
		return err
	}

	var response map[string]interface{}
	err = readFromPipe(conn, &response)

	if err != nil {
		return err
	}

	if response["code"] != nil {
		return errors.New("failed to complete handshake")
	}

	client.conn = conn

	return nil
}

func (client *Client) SetActivity(state, details, largeImage, largeText string) (error) {
	return writeCommandToPipe(client.conn, rpcSetActivityDirective, setActivityCommand{
		ProcessId: os.Getpid(),
		Activity: &activityArgument{
			State:   state,
			Details: details,
			Assets: &assetArgument{
				LargeImage: largeImage,
				LargeText:  largeText,
			},
			Instance: true,
		},
	})
}

func (client *Client) ClearActivity() (error) {
	return writeCommandToPipe(client.conn, rpcSetActivityDirective, &setActivityCommand{
		ProcessId: os.Getpid(),
		Activity:  nil,
	})
}