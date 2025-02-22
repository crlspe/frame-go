package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/crlspe/frame-go/color"
	"github.com/spf13/pflag"
)

type ActionCommand interface {
	Exec(Flags)
}

type command struct {
	cmd           *pflag.FlagSet
	actionCommand ActionCommand
	description   string
}

type climanager struct {
	appName  string
	version  string
	rootSet  pflag.FlagSet
	commands map[string]*command
	flags    map[string]any
}

func NewCliManager(appName, version string) climanager {
	var rootSet = pflag.NewFlagSet("root", pflag.ExitOnError)

	var flags = make(map[string]any)
	flags["help"] = rootSet.BoolP("help", "h", false, "Shows help")
	flags["version"] = rootSet.BoolP("version", "v", false, "Shows version")

	return climanager{
		appName:  appName,
		version:  version,
		rootSet:  *rootSet,
		commands: make(map[string]*command),
		flags:    flags,
	}
}

func (c *climanager) AddCommand(commandName string, description string, actionCommand ActionCommand) {
	var commandFlagSet = pflag.NewFlagSet(commandName, pflag.ExitOnError)
	c.commands[commandName] = &command{
		cmd:           commandFlagSet,
		actionCommand: actionCommand,
		description:   description,
	}
}

func (c *climanager) AddCommandFlagBool(commandName string, name string, shortname string, defaultValue bool, description string) {
	c.flags[commandName+"/"+name] = c.commands[commandName].cmd.BoolP(name, shortname, defaultValue, description)
}

func (c *climanager) AddFlagBool(name string, shortname string, defaultValue bool, description string) {
	c.flags[name] = pflag.BoolP(name, shortname, defaultValue, description)
}

func (c climanager) registerCommands() {
	for _, command := range c.commands {
		c.rootSet.AddFlagSet(command.cmd)
	}
}

func (c climanager) parse() ActionCommand {
	c.registerCommands()
	err := c.rootSet.Parse(os.Args)
	if err != nil {
		fmt.Println("parse root")
		c.printHelp()
		return nil
	}

	if c.getFlagBool("help") {
		c.printHelp()
		return nil
	}
	if c.getFlagBool("version") {
		c.printVersion()
		return nil
	}

	if len(os.Args[1:]) < 1 {
		fmt.Println("no command check")
		c.printHelp()
	}

	var commandName = os.Args[1]
	if _, ok := c.commands[commandName]; !ok {
		fmt.Println("command does not exists")
		c.printHelp()
		return nil
	}

	c.commands[commandName].cmd.Parse(os.Args[2:])

	if c.getFlagBool(commandName + "/help") {
		c.commands[commandName].cmd.PrintDefaults()
		os.Exit(0)
	}

	return c.commands[commandName].actionCommand
}

func (c *climanager) RunCli() {
	var cmd = c.parse()
	if cmd != nil {
		cmd.Exec(Flags{
			values: &c.flags,
		})
	}
}

func (c climanager) getFlagBool(flagName string) bool {
	if val, ok := c.flags[flagName]; ok {
		return val.(bool)
	}
	return false
}
func (c climanager) getFlag(flagName string) any {
	if val, ok := c.flags[flagName]; ok {
		return val
	}
	return nil
}

func (c climanager) printHelp() {
	c.printVersion()
	fmt.Println("Usage:", color.Blue("snp"), color.Yellow("<command>"), color.Green("<args>"))
	fmt.Println("Available commands:")
	for cmdname, set := range c.commands {
		fmt.Println(color.Yellow(cmdname) + "\t" + set.description)
		set.cmd.PrintDefaults()
	}
	os.Exit(0)
}

func (c climanager) printVersion() {
	fmt.Println(color.BrightBlue(strings.ToUpper(c.appName) + " Version " + c.version))
}
