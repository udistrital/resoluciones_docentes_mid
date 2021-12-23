package models

type RpDocente struct {
	CdpRpDocente struct {
		CdpRp []struct {
			Cdp      string `json:"cdp"`
			Vigencia string `json:"vigencia"`
			Rp       string `json:"rp"`
		} `json:"cdp_rp"`
	} `json:"cdp_rp_docente"`
}
