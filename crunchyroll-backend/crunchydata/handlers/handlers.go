package myCustomHandlers

type VideoAnime struct {
	VideoTitle     string
	Year           int
	NumberOfVideos int
	ImgUrl         string
}

type ScrapedVideoAnimeTitle struct{}

// Create a function called GetVideoTitleCollection that returns a VideoAnime slice.
func GetVideoTitleCollection() []VideoAnime {
	animeCollection := []VideoAnime{
		{VideoTitle: "One Piece", Year: 1998, NumberOfVideos: 1, ImgUrl: "https://i.ytimg.com/vi/Q-_q-_q-_q/maxresdefault.jpg"},
		{VideoTitle: "Dragonball Super", Year: 1997, NumberOfVideos: 189, ImgUrl: "https://i.ytimg.com/vi/Q-_q-_q-_q/maxresdefault.jpg"},
		{VideoTitle: "Naruto", Year: 1996, NumberOfVideos: 622, ImgUrl: "https://i.ytimg.com/vi/Q-_q-_q-_q/maxresdefault.jpg"},
	}
	return animeCollection
}

// func GetAnimeByVideoByTitle(videoTitle string) {
// 	animeVideo := func() VideoAnime {
// 		return VideoAnime{}
// 	}
// 	return *animeVideo
// }

// Here are the following function for scraping.

func fetchScrapedAnimeVideoValueByTitle() {}
