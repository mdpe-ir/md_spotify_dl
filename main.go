package main

import (
	"context"
	"fmt"
	"github.com/inancgumus/screen"
	mdspotifydl "github.com/mdpe-ir/md_spotify_dl/src"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func main() {
	var trackID string
	var playlistID string
	var albumID string
	var spotifyURL string

	var rootCmd = &cobra.Command{
		Use:     mdspotifydl.AppUse,
		Version: mdspotifydl.AppVersion,
		Short:   mdspotifydl.AppShortDescription,
		Long:    mdspotifydl.AppLongDescription,
		Run: func(cmd *cobra.Command, args []string) {

			screen.Clear()

			ctx := context.Background()

			if len(args) == 0 {
				_ = cmd.Help()
				fmt.Println("")
				os.Exit(0)
			}

			isSpotifyUrl := true
			spotifyURL = args[0]

			if len(spotifyURL) == 0 {
				fmt.Println("=> Spotify URL Or Trak Name required.")
				_ = cmd.Help()
				return
			}

			splitURL := strings.Split(spotifyURL, "/")

			if len(splitURL) < 2 {
				isSpotifyUrl = false
				//fmt.Println("=> Please enter the url copied from the spotify client.")
				//os.Exit(1)
			}

			if isSpotifyUrl {
				spotifyID := splitURL[len(splitURL)-1]
				if strings.Contains(spotifyID, "?") {
					spotifyID = strings.Split(spotifyID, "?")[0]
				}

				if strings.Contains(spotifyURL, "album") {
					albumID = spotifyID
					mdspotifydl.DownloadAlbum(ctx, albumID)
				} else if strings.Contains(spotifyURL, "playlist") {
					playlistID = spotifyID
					mdspotifydl.DownloadPlaylist(ctx, playlistID)
				} else if strings.Contains(spotifyURL, "track") {
					trackID = spotifyID
					mdspotifydl.DownloadSong(ctx, trackID)
				} else {
					fmt.Println("=> Only Spotify Album/Playlist/Track URL's are supported.")
					_ = cmd.Help()
				}
			}

			if !isSpotifyUrl {
				fmt.Println("=> Start searching in spotify for " + spotifyURL + " ...")
				os.Exit(1)
			}
		},
	}

	rootCmd.SetUsageTemplate(fmt.Sprintf("%s [spotify_url or track name] \n", mdspotifydl.AppUse))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
