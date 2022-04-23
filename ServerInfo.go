package AmazonSessionAPIClient

type ServerInfo struct {
	ProxyCount                        int  `json:"proxyCount"`
	SessionCreatorProxyCount          int  `json:"sessionCreatorProxyCount"`
	SessionCount                      int  `json:"sessionCount"`
	SessionsAreCreating               bool `json:"sessionsAreCreating"`
	SearchingForSessionCreatorProxies bool `json:"searchingForSessionCreatorProxies"`
}
