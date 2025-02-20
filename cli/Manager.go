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
	cmd 				*pflag.FlagSet
	actionCommand 		ActionCommand
	description			string
}

type manager struct {
	appName			string
	version			string
	rootSet			pflag.FlagSet
	commands		map[string]*command
	flags 			map[string]any
}

func NewCliManager(appName, version string) manager {
	var rootSet = pflag.NewFlagSet("root", pflag.ExitOnError)

	var flags = make(map[string]any)
	flags["help"] = rootSet.BoolP("help", "h", false, "Shows help")
	flags["version"] = rootSet.BoolP("version", "v", false, "Shows version")

	return manager{
		appName: appName,
		version: version,
		rootSet: *rootSet,
		commands: make(map[string]*command),
		flags: flags,
	}
}

func (m *manager) AddCommand(commandName string, description string, actionCommand ActionCommand){
	var commandFlagSet = pflag.NewFlagSet(commandName, pflag.ExitOnError)
	m.commands[commandName] = &command{
        cmd:            commandFlagSet,
        actionCommand:  actionCommand,
        description:    description,
    }
}

func (m *manager) AddCommandFlagBool(commandName string, name string, shortname string, defaultValue bool, description string) {
	m.flags[commandName+"/"+name] = m.commands[commandName].cmd.BoolP(name,shortname,defaultValue,description)
}

func (m *manager) AddFlagBool(name string, shortname string, defaultValue bool, description string) {
	m.flags[name] = pflag.BoolP(name,shortname,defaultValue,description)
}

func (m manager) registerCommands() {
	for _ ,command := range m.commands {
		m.rootSet.AddFlagSet(command.cmd)
	}
}

func (m manager) parse() ActionCommand {
	m.registerCommands()
    err := m.rootSet.Parse(os.Args)
    if err != nil {
    	fmt.Println("parse root")
        m.printHelp()
        return nil
    }

    if *m.getFlag("help").(*bool) {
        m.printHelp()
        return nil
    }
    if *m.getFlag("version").(*bool) {
        m.printVersion()
        return nil
    }

    if len(os.Args[1:]) < 1 {
    	fmt.Println("no command check")
        m.printHelp()
    }

    var commandName = os.Args[1]
    if _, ok := m.commands[commandName]; !ok {
    	fmt.Println("command does not exists")
        m.printHelp()
        return nil
    }

    m.commands[commandName].cmd.Parse(os.Args[2:])

    if *m.getFlag(commandName+"/help").(*bool) {
        m.commands[commandName].cmd.PrintDefaults()
        os.Exit(0)
    }

    return m.commands[commandName].actionCommand
}

func (m *manager) RunCli() {
	var cmd = m.parse()
	if cmd != nil {
		cmd.Exec(Flags{
			values: &m.flags,
		})
	}
}


func (m manager) getFlag(flagName string) any {
    if val, ok := m.flags[flagName]; ok {
        return val
    }
    return nil
}

func (m manager) printHelp() {
	m.printVersion()
	fmt.Println("Usage:", color.Blue("snp"), color.Yellow("<command>"), color.Green("<args>"))
	fmt.Println("Available commands:")
	for cmdname , set := range m.commands {
		fmt.Println(color.Yellow(cmdname) + "\t" + set.description)
		set.cmd.PrintDefaults()
	}
	os.Exit(0)
}

func (m manager) printVersion() {
	fmt.Println(color.BrightBlue(strings.ToUpper(m.appName) + " Version " + m.version))
}
