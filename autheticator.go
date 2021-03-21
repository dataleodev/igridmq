package igridmq

import "context"

type Op string

const (
	PUBISH    Op = "publish"
	SUBSCRIBE Op = "subscribe"
	CONNECT   Op = "connect"
)

type AuthNZParams struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	IP       string   `json:"ip"`
	Topics   []string `json:"topics"`
	Op       string   `json:"op"`
	Payload  []byte   `json:"payload"`
}

type AuthNZ interface {

	//Pass is AuthNZ (Authentication and Authorization) client that either pass the message
	//or block it
	//username - name of the client
	Pass(ctx context.Context, clientId string, username string, password string,
		ip string, topics []string, op string, payload []byte) (err error)
}
