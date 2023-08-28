package cmd

type Option struct {
	Name        string
	Default     string
	Description string
}

var DefaultPortOption = Option{
	Name:        "port",
	Default:     "8080",
	Description: "git docs server port",
}

var DefaultHostOption = Option{
	Name:        "host",
	Default:     "localhost",
	Description: "git docs server host",
}

var DefaultSourceOption = Option{
	Name:        "source",
	Default:     "docs",
	Description: "git docs direcotry",
}
