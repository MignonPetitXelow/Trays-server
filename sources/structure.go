package main

type Object struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Hash string `json:"hash"`
}

type Files struct {
	Script  []Object `json:"script"`
	Config  []Object `json:"config"`
	Mods    []Mod    `json:"mods"`
	Version []Object `json:"version"`
}

type Mod struct {
	File    Object `json:"file"`
	Version string `json:"version"`
	Icon    string `json:"icon"`
}

type Pack struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	ApiVersion string `json:"apiversion"`
	Path       string `json:"path"`
	Files      Files  `json:"files"`
	Hash       string `json:"hash"`
	Icon       string `json:"icon"`
	Background string `json:"background"`
	Color      string `json:"color"`
}
