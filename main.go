package main

import (
	"fmt"
	"strings"

	"./cmd"
	"./framework"
	"github.com/bwmarrin/discordgo"
)

var (
	conf       *framework.Config
	CmdHandler *framework.CommandHandler
	Sessions   *framework.SessionManager
	youtube    *framework.Youtube
	botID      string
	PREFIX     string
)

func init() {
	conf = framework.LoadConfig("config.json")
	PREFIX = conf.Prefix

}

func main() {
	CmdHandler = framework.NewCommandHandler()
	registerCommands()
	Sessions = framework.NewSessionManager()
	youtube = &framework.Youtube{Conf: conf}
	discord, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}
	if conf.UseSharding {
		discord.ShardID = conf.ShardId
		discord.ShardCount = conf.ShardCount
	}
	usr, err := discord.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details,", err)
		return
	}
	botID = usr.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateStatus(0, conf.DefaultStatus)
		guilds := discord.State.Guilds
		fmt.Println("Ready with", len(guilds), "guilds.")
	})
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}
	fmt.Println("Started")
	<-make(chan struct{})
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		return
	}
	content := message.Content
	fmt.Println(content)
	// fmt.Println(len(content) <= len(PREFIX))
	if len(content) <= len(PREFIX) {
		return
	}
	// fmt.Println(content[:len(PREFIX)] != PREFIX)
	if content[:len(PREFIX)] != PREFIX {
		return
	}
	content = content[len(PREFIX):]
	if len(content) < 1 {
		return
	}
	args := strings.Fields(content)
	name := strings.ToLower(args[0])
	command, found := CmdHandler.Get(name)
	if !found {
		return
	}
	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel,", err)
		return
	}
	guild, err := discord.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return
	}
	ctx := framework.NewContext(discord, guild, channel, user, message, conf, CmdHandler, Sessions, youtube)
	ctx.Args = args[1:]
	c := *command
	c(*ctx)
}

func registerCommands() {
	CmdHandler.Register("tulung", cmd.HelpCommand, "Help command bro!")
	CmdHandler.Register("h", cmd.HelpCommand, "Help command bro!")
	CmdHandler.Register("admin", cmd.AdminCommand, "")
	CmdHandler.Register("join", cmd.JoinCommand, "Join voice channel syntax join <arg>")
	CmdHandler.Register("j", cmd.JoinCommand, "Join voice channel syntax join <arg>")
	CmdHandler.Register("leave", cmd.LeaveCommand, "Cabut dari voice channel!")
	CmdHandler.Register("l", cmd.LeaveCommand, "Cabut dari voice channel!")
	CmdHandler.Register("play", cmd.PlayCommand, "Ngeplay queue bro!")
	CmdHandler.Register("p", cmd.PlayCommand, "Ngeplay queue bro!")
	CmdHandler.Register("stop", cmd.StopCommand, "Rem not blong!")
	CmdHandler.Register("st", cmd.StopCommand, "Rem not blong!")
	CmdHandler.Register("add", cmd.AddCommand, "Add a song to the queue `add <youtube-link>")
	CmdHandler.Register("a", cmd.AddCommand, "Add a song to the queue `add <youtube-link>")
	CmdHandler.Register("skip", cmd.SkipCommand, "Skip")
	CmdHandler.Register("s", cmd.SkipCommand, "Skip")
	CmdHandler.Register("queue", cmd.QueueCommand, "Print queue lah...")
	CmdHandler.Register("q", cmd.QueueCommand, "Print queue lah...")
	CmdHandler.Register("clear", cmd.ClearCommand, "Bersihin queue lah...")
	CmdHandler.Register("c", cmd.ClearCommand, "Bersihin queue lah...")
	CmdHandler.Register("curr", cmd.CurrentCommand, "Judul current song lah...")
	CmdHandler.Register("youtube", cmd.YoutubeCommand, "Jangan pake dulu brok sabar")
	CmdHandler.Register("yt", cmd.YoutubeCommand, "Jangan pake dulu brok sabar")
	CmdHandler.Register("shuf", cmd.ShuffleCommand, "Shuffle queue lah...")
	CmdHandler.Register("pauseq", cmd.PauseCommand, "Pause lagu lah...")
	CmdHandler.Register("pq", cmd.PauseCommand, "Pause lagu lah...")
	// CmdHandler.Register("gadd", cmd.AddGenreCommand, "Request by Moses Kevin")
}
