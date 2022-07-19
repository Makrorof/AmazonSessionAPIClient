package AmazonSessionAPIClient

type ServerInfo struct {
	ProxyCount                        int            `json:"proxyCount"`
	SessionCreatorProxyCount          map[string]int `json:"sessionCreatorProxyCount"`
	SessionCount                      map[string]int `json:"sessionCount"`
	UsableSessionCount                map[string]int `json:"usableSessionCount"`
	SessionsAreCreating               bool           `json:"sessionsAreCreating"`
	SearchingForSessionCreatorProxies bool           `json:"searchingForSessionCreatorProxies"`
	UseServerIP                       bool           `json:"useServerIP"`
}
