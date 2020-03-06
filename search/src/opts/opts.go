package opts

type Opts struct {
	Base   string `short:"b" long:"base" description:"The root folder from where to look after files" required:"false" default:"./"`
	Max    int    `short:"m" long:"max" description:"Maximum amount of matches" required:"false" default:"100"`
	Search string ""
}
