package basefunc

//ChartYaml chart下的Chart.yaml
type ChartYaml struct {
	Name        string                 `json:"name"`
	Home        string                 `json:"home"`
	Version     string                 `json:"version"`
	AppVersion  string                 `json:"appVersion"`
	APIVersion  string                 `json:"apiVersion"`
	Description string                 `json:"description"`
	Sources     []string               `json:"sources"`
	Maintainers []ChartYamlMaintainers `json:"maintainers"`
	Email       string                 `json:"  email"`
}

//ChartYamlMaintainers ChartYaml-Maintainers
type ChartYamlMaintainers struct {
	Name string `json:"name"`
}
