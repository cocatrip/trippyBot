package cmd

import (
	"bytes"
	"fmt"
	"sort"

	"../framework"
)

func HelpCommand(ctx framework.Context) {
	cmds := ctx.CmdHandler.GetCmds()
	buffer := bytes.NewBufferString("Commands:\n")
	fmt.Println(cmds)

	keys := make([]string, 0, len(cmds))
	for k := range cmds {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, cmdName := range keys {
		if len(cmdName) == 1 {
			continue
		}
		switch cmdName {
		case "admin":
			continue
		case "q":
			continue
		}
		fmt.Println(cmdName, cmds[cmdName])
		msg := fmt.Sprintf("`\t %s%s - %s`\n", ctx.Conf.Prefix, cmdName, cmds[cmdName].GetHelp())
		buffer.WriteString(msg)
	}
	str := buffer.String()
	ctx.Reply(str[:len(str)-2])
}
