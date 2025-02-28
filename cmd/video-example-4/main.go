package main

import (
	"test/pkg"
	"test/pkg2"
)

func main() {
	videoSource := []pkg2.VideoSource{
		// I woke up and chose the violence
		{Source: "https://youtube.com/clip/UgkxSaK_spNaPIT3ug3iOtXn_fjQruG_veBS?si=13r_MRSBRFmHwGK7"},

		// CHROMIUM BROO
		{Source: "https://youtube.com/clip/Ugkx1pYKaM52emqwVeDOZoi_NEdh5kyTcTZM?si=VNF2yAuN6sc4O6Ho", Interval: "00:00:0.3-inf"},
	}

	videoCompiler := pkg.NewVideoCompiler(pkg.VideoCompilerConfig{
		CacheFolder:     "./videos-cache-high",
		CacheFileType:   "mp4",
		MaxVideoQuality: "1080",
	})

	ctx := pkg2.RenderCtx{
		CacheFolder:            "./videos-cache-high",
		CacheFileType:          "mp4",
		DownloadQualityOptions: "",
	}

	videoResources, err := pkg2.InitializeVideoResources(ctx, videoSource...)
	if err != nil {
		panic(err)
	}

	pkg2.MergeVideos(videoResources...)

	err := videoCompiler.AddVideoSources(videoSource...)
	if err != nil {
		panic(err)
	}

	_, err = videoCompiler.MergeVideos("finalVideo.mp4")
	if err != nil {
		panic(err)
	}

}
