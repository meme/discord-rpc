package discord_rpc

type RpcCommand struct {
	Command   string      `json:"cmd"`
	Arguments interface{} `json:"args"`
	Nonce     string      `json:"nonce"`
}

type TimestampArgument struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type AssetArgument struct {
	LargeImage string  `json:"large_image"`
	LargeText  string  `json:"large_text"`
	SmallImage *string `json:"small_image,omitempty"`
	SmallText  *string `json:"small_text,omitempty"`
}

type PartyArgument struct {
	Id   string `json:"id"`
	Size []int  `json:"size"`
}

type SecretsArgument struct {
	Join     string `json:"join"`
	Spectate string `json:"spectate"`
	Match    string `json:"match"`
}

type ActivityArgument struct {
	State      string             `json:"state"`
	Details    string             `json:"details"`
	Timestamps *TimestampArgument `json:"timestamps,omitempty"`
	Assets     *AssetArgument     `json:"assets,omitempty"`
	Party      *PartyArgument     `json:"party,omitempty"`
	Secrets    *SecretsArgument   `json:"secrets,omitempty"`
	Instance   bool               `json:"instance"`
}

type SetActivityCommand struct {
	ProcessId int              `json:"pid"`
	Activity  *ActivityArgument `json:"activity"`
}

type RpcHandshake struct {
	Version  int32  `json:"v"`
	ClientId string `json:"client_id"`
}

type RpcResponse struct {
	Command string                 `json:"cmd"`
	Data    map[string]interface{} `json:"data"`
	Event   string                 `json:"evt"`
	Nonce   string                 `json:"nonce"`
}
