package util

type OcppTopic struct {
	host string
	action string
}

func NewOcppTopic(host string, action string) *OcppTopic {
	return &OcppTopic{host: host,action: action}
}

func (r *OcppTopic) String() string {
	return r.host + ":" + r.action
}
