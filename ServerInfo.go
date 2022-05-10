package AmazonSessionAPIClient

type ServerInfo struct {
	ProxyCount                        int            `json:"proxyCount"`
	SessionCreatorProxyCount          map[string]int `json:"sessionCreatorProxyCount"`
	SessionCount                      int            `json:"sessionCount"`
	UsableSessionCount                int            `json:"usableSessionCount"`
	SessionsAreCreating               bool           `json:"sessionsAreCreating"`
	SearchingForSessionCreatorProxies bool           `json:"searchingForSessionCreatorProxies"`
}
