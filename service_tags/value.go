package service_tags

//Value represents a ServiceTag value
type Value struct {
	Name            string   `json:"name"`
	Id              string   `json:"id"`
	Properties      Property `json:"properties"`
	NetworkFeatures []string `json:"networkFeatures"`
}

//Property represents a Value property
type Property struct {
	ChangeNumber    int64    `json:"changeNumber"`
	Region          string   `json:"region"`
	RegionId        int64    `json:"regionId"`
	Platform        string   `json:"platform"`
	SystemService   string   `json:"systemService"`
	AddressPrefixes []string `json:"addressPrefixes"`
}
