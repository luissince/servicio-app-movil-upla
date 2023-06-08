package model

type DatosFicha struct {
	Codigo             string `json:"cod,omitempty"`
	ApellidoPaterno    string `json:"apellidoPaterno,omitempty"`
	ApellidoMaterno    string `json:"apellidoMaterno,omitempty"`
	Nombres            string `json:"nombres,omitempty"`
	FechaNacimiento    string `json:"fechaNacimiento,omitempty"`
	Telefono           string `json:"telefono,omitempty"`
	Celular            string `json:"celular,omitempty"`
	Mail               string `json:"mail,omitempty"`
	FechaIngreso       string `json:"fechaIngreso,omitempty"`
	Facultad           string `json:"facultad,omitempty"`
	Carrera            string `json:"carrera,omitempty"`
	ModalidadAcademico string `json:"modalidadAcademico,omitempty"`
	Direccion          string `json:"direccion,omitempty"`
	EstadoCivil        string `json:"estadoCivil,omitempty"`
	Dni                string `json:"dni,omitempty"`
	Edad               string `json:"edad,omitempty"`
}
