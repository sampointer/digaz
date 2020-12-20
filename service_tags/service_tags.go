package service_tags

type ServiceTags struct {
	ChangeNumber int64    `json:"changeNumber"`
	Cloud        string   `json:"cloud"`
	Values       []Ranges `json:"values"`
}

//New is a constructor for ServiceTags
func New(r io.Reader) (*ServiceTags, error) {
	var serviceTags ServiceTags
	return &serviceTags, nil
}
