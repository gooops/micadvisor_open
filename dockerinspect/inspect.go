package dockerinspect

type InspectConfig struct {
	Hostname string
	Labels   interface{}
}

type Inspect struct {
	Id      string
	Created string
	Name    string
	Config  *InspectConfig
}
