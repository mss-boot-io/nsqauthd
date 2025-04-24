package dto

type AuthReq struct {
	Secret   string `query:"secret" form:"secret"`
	RemoteIP string `query:"remote_ip" form:"remote_ip"`
	TLS      bool   `query:"tls" form:"tls"`
}

type AuthResp struct {
	Message        string         `json:"message,omitempty"`
	Identity       string         `json:"identity,omitempty"`
	TTL            int            `json:"ttl,omitempty"`
	Authorizations []*AuthAccount `json:"authorizations,omitempty"`
}

type AuthAccount struct {
	Channels    []string `json:"channels"`
	Topic       string   `json:"topic"`
	Permissions []string `json:"permissions"`
}
