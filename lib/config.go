package gohelper

type Config struct {
	Name					string			`json:"name"`
	Version					int				`json:"version"`
	MyObject				SubObject		`json:"myobject"`
}

type SubObject struct {
	Enabled					bool			`json:"enabled"`
	ID						float64			`json:"id"`
}
