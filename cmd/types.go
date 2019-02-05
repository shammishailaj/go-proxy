package cmd

var (
	port        int
	proxy       string
	setAll      bool   = false
	password    string = ""
	cntlmFile   string = "/usr/local/etc/cntlm.conf"
	bashProfile string = "~/.bash_profile"
)

type execCommand struct {
	cmd  string
	args []string
}
