package model

type DatosFicha struct {
	Cod    string `json:"cod,omitempty"`
	ApPat  string `json:"ap_pat,omitempty"`
	ApMat  string `json:"ap_mat,omitempty"`
	Nom    string `json:"nom,omitempty"`
	FNac   string `json:"f_nac,omitempty"`
	Tel    string `json:"tel,omitempty"`
	Cel    string `json:"cel,omitempty"`
	Mail   string `json:"mail,omitempty"`
	FIng   string `json:"f_ing,omitempty"`
	Fac    string `json:"fac,omitempty"`
	Carr   string `json:"carr,omitempty"`
	ModAcd string `json:"mod_acd,omitempty"`
	Dir    string `json:"dir,omitempty"`
	ECiv   string `json:"e_civ,omitempty"`
	Dni    string `json:"dni,omitempty"`
	Edad   string `json:"edad,omitempty"`
}
