package model

type Profile struct {
	Uid        string
	Nickname   string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Xinzuo     string
	House      string
	Car        string
}

type Citylist struct {
	Url  string
	Name string
}

type City struct {
	Url  string
	Name string
}
