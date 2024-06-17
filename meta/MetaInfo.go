package meta

type MetaInfo struct {
	GitBranch    string `json:"branch"`
	GitHash      string `json:"hash"`
	GitAuthor    string `json:"author"`
	BuildDate    string `json:"date"`
	BuildDocker  string `json:"docker"`
	BuildVersion string `json:"version"`
	AppId        string `json:"appid"`
}

var GitBranch string
var GitHash string
var GitAuthor string
var BuildDate string
var BuildDocker string
var BuildVersion string
