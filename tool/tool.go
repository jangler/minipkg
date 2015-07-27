package tool

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// Description is the description that is printed in the main usage message. If
// set, the text should begin and end with a newline.
var Description string

// Command is a subcommand of the main tool.
type Command struct {
	Name        string
	Summary     string // one-line description for command list
	Usage       string // metavariables displayed in usage message
	Description string // block description for usage message
	Function    func([]string)
	FlagSet     *flag.FlagSet
	MinArgs     int
	MaxArgs     int
	HasOpts     bool
}

// Commands is the list of subcommands available to the program.
var Commands = make(map[string]*Command)

// UsageFunc constructs a FlagSet.Usage function from a Command.
func UsageFunc(cmd *Command) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "Usage: %s %s %s\n",
			os.Args[0], cmd.Name, cmd.Usage)
		fmt.Fprint(os.Stderr, cmd.Description)
		if cmd.HasOpts {
			fmt.Fprintln(os.Stderr, "\nOptions:")
			cmd.FlagSet.PrintDefaults()
		}
	}
}

func readStdin(args []string, max int) []string {
	r := bufio.NewReader(os.Stdin)
	for max < 0 || len(args) < max {
		line, err := r.ReadString('\n')
		if len(line) > 0 {
			args = append(args, string(line[:len(line)-1]))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return args
}

func getArgs(cmd string, min, max int, args []string) []string {
	if max >= 0 && len(args) > max {
		fmt.Fprintf(os.Stderr, "%s: %s: too many arguments\n", os.Args[0], cmd)
		os.Exit(2)
	}

	if len(args) < min {
		args = readStdin(args, max)
	}
	if len(args) < min {
		fmt.Fprintf(os.Stderr, "%s: %s: missing argument\n", os.Args[0],
			cmd)
		os.Exit(2)
	}

	return args
}

func printCommands() {
	maxlen := 0
	for _, cmd := range Commands {
		if len(cmd.Name) > maxlen {
			maxlen = len(cmd.Name)
		}
	}
	space := "                "

	for _, cmd := range Commands {
		fmt.Fprintf(os.Stderr, "  %s%s  %s\n", cmd.Name,
			space[:maxlen-len(cmd.Name)], cmd.Summary)
	}
}

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <cmd> [<arg>]...\n", os.Args[0])
		fmt.Fprint(os.Stderr, Description)
		fmt.Fprintf(os.Stderr, `
If not enough command-line arguments are specified for a command,
remaining arguments are read from standard input. For help regarding a
specific command, see '%s <cmd> -h'.
`, os.Args[0])
		fmt.Fprint(os.Stderr, "\nCommands:\n")
		printCommands()
		os.Exit(2)
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
	}
}

// Main runs the tool. Commands should be initialized before calling this
// function.
func Main() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", os.Args[0]))

	parseFlags()

	if cmd := Commands[flag.Arg(0)]; cmd != nil {
		log.SetPrefix(fmt.Sprintf("%s: %s: ", os.Args[0], cmd.Name))
		if err := cmd.FlagSet.Parse(flag.Args()[1:]); err == flag.ErrHelp {
			cmd.FlagSet.Usage()
			os.Exit(2)
		} else if err != nil {
			log.Fatal(err)
		}
		args := getArgs(cmd.Name, cmd.MinArgs, cmd.MaxArgs, cmd.FlagSet.Args())
		cmd.Function(args)
	} else {
		fmt.Fprintf(os.Stderr, "%s: no such command: %s\n", os.Args[0],
			flag.Arg(0))
		os.Exit(2)
	}
}
