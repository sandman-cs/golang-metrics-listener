package main

// Configuration File Opjects
type configuration struct {
	ServerName     string
	AppName        string
	AppVer         string
	Broker         string
	BrokerUser     string
	BrokerPwd      string
	BrokerExchange string
	BrokerVhost    string
	ChannelSize    int
	ChannelCount   int
	SrvPort        string
}

type metricPayload struct {
	Bu      string `json:"bu"`
	Env     string `json:"env"`
	HostEnv string `json:"hostEnv"`
	App     string `json:"app"`
	Version string `json:"version"`
	Route   string `json:"route"`
	Token   string `json:"token"`
	Points  []struct {
		Name      string        `json:"name"`
		Value     int           `json:"value"`
		Timestamp int           `json:"timestamp"`
		Tags      []interface{} `json:"tags"`
	} `json:"points"`
}
