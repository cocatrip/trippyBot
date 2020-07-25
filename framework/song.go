package framework

import (
	"os/exec"
	"strconv"
)

// Song struct
type Song struct {
	Media    string
	Title    string
	Duration *string
	ID       string
}

// Ffmpeg global function
func (song Song) Ffmpeg() *exec.Cmd {
	return exec.Command("ffmpeg", "-i", song.Media, "-f", "s16le", "-ar", strconv.Itoa(FrameRate), "-ac",
		strconv.Itoa(Channel), "pipe:1")
}

// NewSong global function
func NewSong(media, title, id string) *Song {
	song := new(Song)
	song.Media = media
	song.Title = title
	song.ID = id
	return song
}
