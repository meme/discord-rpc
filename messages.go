package discord_rpc

type rpcCommand struct {
	Command   string      `json:"cmd"`
	Arguments interface{} `json:"args"`
	Nonce     string      `json:"nonce"`
}

type timestampArgument struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type assetArgument struct {
	LargeImage string  `json:"large_image"`
	LargeText  string  `json:"large_text"`
	SmallImage *string `json:"small_image,omitempty"`
	SmallText  *string `json:"small_text,omitempty"`
}

type partyArgument struct {
	Id   string `json:"id"`
	Size []int  `json:"size"`
}

type secretsArgument struct {
	Join     string `json:"join"`
	Spectate string `json:"spectate"`
	Match    string `json:"match"`
}

type activityArgument struct {
	State      string             `json:"state"`
	Details    string             `json:"details"`
	Timestamps *timestampArgument `json:"timestamps,omitempty"`
	Assets     *assetArgument     `json:"assets,omitempty"`
	Party      *partyArgument     `json:"party,omitempty"`
	Secrets    *secretsArgument   `json:"secrets,omitempty"`
	Instance   bool               `json:"instance"`
}

type setActivityCommand struct {
	ProcessId int               `json:"pid"`
	Activity  *activityArgument `json:"activity"`
}

type rpcHandshake struct {
	Version  int32  `json:"v"`
	ClientId string `json:"client_id"`
}

type rpcResponse struct {
	Command string                 `json:"cmd"`
	Data    map[string]interface{} `json:"data"`
	Event   string                 `json:"evt"`
	Nonce   string                 `json:"nonce"`
}
