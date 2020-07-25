package cmd

import (
	"fmt"

	"../framework"
)

func AddCommand(ctx framework.Context) {
	if len(ctx.Args) == 0 {
		ctx.Reply("Add command usage: `:add <song>`\nValid inputs: `youtube url`, `soundcloud url`, " +
			"`youtube id`, `soundcloud id`")
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply("Not in a voice channel! To make the bot join one, use `:join`.")
		return
	}
	fmt.Println(ctx.Args[0])
	switch ctx.Args[0] {
	case "pop":
		ctx.Args[0] = "https://www.youtube.com/playlist?list=PLMC9KNkIncKtPzgY-5rmhvj7fax8fdxoj"
	case "rap":
		ctx.Args[0] = "https://www.youtube.com/playlist?list=PLw-VjHDlEOgsIgak3vJ7mrcy-OscZ6OAs"
	case "edm":
		ctx.Args[0] = "https://www.youtube.com/playlist?list=PLw-VjHDlEOgs658kAHR_LAaILBXb-s6Q5"
	case "triphop":
		ctx.Args[0] = "https://www.youtube.com/playlist?list=PLYSmQ0A-61NNWwuLj_C3c7Lm-SkGXbVtz"
	case "rock":
		ctx.Args[0] = "https://www.youtube.com/playlist?list=PL3oW2tjiIxvQWubWyTI8PF80gV6Kpk6Rr"
	}
	msg := ctx.Reply("Adding songs to queue...")
	for _, arg := range ctx.Args {
		t, inp, err := ctx.Youtube.Get(arg)
		if err != nil {
			ctx.Reply("An error occured!")
			fmt.Println("error getting input, %s", err)
			return
		}
		switch t {
		case framework.ErrorType:
			ctx.Reply("An error occured!")
			fmt.Println("error type", t)
			return
		case framework.VIDEO_TYPE:
			{
				video, err := ctx.Youtube.Video(*inp)
				if err != nil {
					ctx.Reply("An error occured!")
					fmt.Println("error getting video1,", err)
					return
				}
				song := framework.NewSong(video.Media, video.Title, arg)
				sess.Queue.Add(*song)
				ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, "Added `"+song.Title+"` to the song queue."+" Use `music play` to start playing the songs! To see the song queue, use `music queue`.")
				break
			}
		case framework.PLAYLIST_TYPE:
			{
				videos, err := ctx.Youtube.Playlist(*inp)
				if err != nil {
					ctx.Reply("An error occured!")
					fmt.Println("error getting playlist,", err)
					return
				}
				for _, v := range *videos {
					id := v.Id
					_, i, err := ctx.Youtube.Get(id)
					if err != nil {
						ctx.Reply("An error occured!")
						fmt.Println("error getting video2,", err)
						continue
					}
					video, err := ctx.Youtube.Video(*i)
					if err != nil {
						ctx.Reply("An error occured!")
						fmt.Println("error getting video3,", err)
						return
					}
					song := framework.NewSong(video.Media, video.Title, arg)
					sess.Queue.Add(*song)
				}
				if err == nil {
					ctx.Reply("Finished adding songs to the playlist. Use `:play` to start playing the songs! \n" + "To see the song queue, use `:queue`.")
				}
				break
			}
		}
	}
}

// 	msg := ctx.Reply("Adding songs to queue...")
// 	for _, arg := range ctx.Args {
// 		t, inp, err := ctx.Youtube.Get(arg)

// 		if err != nil {
// 			ctx.Reply("An error occured!")
// 			fmt.Println("error getting input, %s", err)
// 			return
// 		}

// 		switch t {
// 		case framework.ErrorType:
// 			ctx.Reply("An error occured!")
// 			fmt.Println("error type", t)
// 			return
// 		case framework.VIDEO_TYPE:
// 			{
// 				video, err := ctx.Youtube.Video(*inp)
// 				if err != nil {
// 					ctx.Reply("An error occured!")
// 					fmt.Println("error getting video1,", err)
// 					return
// 				}
// 				song := framework.NewSong(video.Media, video.Title, arg)
// 				sess.Queue.Add(*song)
// 				ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, "Added `"+song.Title+"` to the song queue."+
// 					" Use `:play` to start playing the songs! To see the song queue, use `:queue`.")
// 				break
// 			}
// 		case framework.PLAYLIST_TYPE:
// 			{
// 				videos, err := ctx.Youtube.Playlist(*inp)
// 				if err != nil {
// 					ctx.Reply("An error occured!")
// 					fmt.Println("error getting playlist,", err)
// 					return
// 				}
// 				for _, v := range *videos {
// 					id := v.Id
// 					_, i, err := ctx.Youtube.Get(id)
// 					if err != nil {
// 						ctx.Reply("An error occured!")
// 						fmt.Println("error getting video2,", err)
// 						continue
// 					}
// 					video, err := ctx.Youtube.Video(*i)
// 					if err != nil {
// 						ctx.Reply("An error occured!")
// 						fmt.Println("error getting video3,", err)
// 						return
// 					}
// 					song := framework.NewSong(video.Media, video.Title, arg)
// 					sess.Queue.Add(*song)
// 				}
// 				ctx.Reply("Finished adding songs to the playlist. Use `:play` to start playing the songs! " +
// 					"To see the song queue, use `:queue`.")
// 				break
// 			}
// 		}
// 	}
// }
