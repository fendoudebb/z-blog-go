package configs

import "time"

type App struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Search   Search   `yaml:"search"`
	Pprof    Pprof    `yaml:"pprof"`
	Website  Website  `yaml:"website"`
}

type Server struct {
	Host            string
	Port            int
	GracefulTimeout time.Duration
}

type Database struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type Search struct {
	Tsconfig  string
	Size      int
	Maxlength int
}

type Pprof struct {
	Host string
	Port int
}

type Website struct {
	Name        string
	Author      string
	Keywords    string
	Description string
	GitHub      string
	Css         map[string]string
	Js          map[string][]string
	Promotion   Promotion
	Comment     Comment
	Icp         ICP
	Meta        []WebsiteMeta
	Analysis    []WebsiteAnalysis
}

type Promotion struct {
	Enabled bool
	Text    string
	Image   string
}

type Comment struct {
	Enabled bool
	Type    string
	Script  string
}

type ICP struct {
	Id   string
	Link string
}

type WebsiteMeta struct {
	Name    string
	Content string
}

type WebsiteAnalysis struct {
	Name   string
	Value  string
	Enable bool
}
