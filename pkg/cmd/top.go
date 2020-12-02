package cmd

import (
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/playlist"
	"github.com/mrydengren/elvis/pkg/spinner"
	"log"
)

func Top(artist string, limit int) {
	spinner.Start("Fetching top tracks.")

	client := lastfm.NewClient()

	options := &lastfm.TopTracksOptions{
		Limit: &limit,
	}

	toptracks, err := client.ArtistTopTracks(artist, options)
	if err != nil {
		spinner.Fail()
		log.Fatal(err)
	}

	tracklist := playlist.FromTopTracks(toptracks)

	spinner.Succeed()

	playlist.Create(tracklist)
}
