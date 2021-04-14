package setup

import "github.com/Foxcapades/Argonaut/v0"

func ParseCLI() (out CLIConfig) {
	cli.NewCommand().
		Flag(cli.SlFlag('c', "config", "Path to the queue config yaml file.").
			Arg(cli.NewArg().
				Required(true).
				Default("queues.yml").
				Bind(&out.ConfigFile))).
		Flag(cli.SlFlag('q', "queue", "URL of the Fireworq queue service").
			Arg(cli.NewArg().
				Required(true).
				Default("http://localhost").
				Bind(&out.QueueURL))).
		MustParse()

	return
}
