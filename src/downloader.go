package md_spotify_dl

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"github.com/zmb3/spotify/v2"
	"io"
	gourl "net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Downloader is a function to download files
func Downloader(url string, track spotify.FullTrack) {

	videonameTag := fmt.Sprintf("%s.mp4", track.Name)
	nameTag := fmt.Sprintf("%s.mp3", track.Name)

	u, err := gourl.ParseRequestURI(url)
	if err != nil {
		fmt.Println("=> An error occured while trying to parse url")
		fmt.Println(err.Error())
		fmt.Println(url)
		os.Exit(1)
	}

	watchId := strings.Split(u.String(), "v=")[1]

	videoID := watchId
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		fmt.Println("=> An error occured while trying to download")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	fmt.Println("=> Select format: ")
	for index, format := range formats {
		fmt.Println("=> Format:", index, " - Audi Quality: ", format.AudioQuality)
	}
	formatNumber := 0
	fmt.Print("Enter Your Format Number: ")
	fmt.Scan(&formatNumber)

	fmt.Println("=> Start Downloading ", videonameTag, " ...")
	stream, _, err := client.GetStream(video, &formats[formatNumber])
	if err != nil {
		fmt.Println("=> An error occured while trying to download")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	file, err := os.Create(videonameTag)
	if err != nil {
		fmt.Println("=> An error occured while trying to download")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Println("=> An error occured while trying to download")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("=> ", videonameTag, "downloaded successfully")
	fmt.Println("=> Extract audio from video")

	currentDirectory, _ := os.Getwd()
	input := strings.TrimSpace(filepath.Join(currentDirectory, videonameTag))

	output := strings.TrimSpace(filepath.Join(currentDirectory, nameTag))

	cmd, err := exec.Command("ffmpeg", "-y", "-i", input, output).CombinedOutput()
	if err != nil {
		fmt.Println("=> An error occured while trying to extract audio from video")
		fmt.Println(err.Error())
		fmt.Println(string(cmd))
		fmt.Println(err)
		os.Exit(1)
	}

	//utils.TagFileWithSpotifyMetadata(nameTag, track)

}
