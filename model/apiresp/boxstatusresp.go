package apiresp

type BoxStatusResp struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"errorMsg"`
	State    int    `json:"state"`
}
