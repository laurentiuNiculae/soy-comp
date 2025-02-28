package main

import (
	"test/pkg"
)

func main() {
	videoSource := []pkg.VideoSource{
		// I woke up and chose the violence
		{URL: "https://youtube.com/clip/UgkxSaK_spNaPIT3ug3iOtXn_fjQruG_veBS?si=13r_MRSBRFmHwGK7"},

		// CHROMIUM BROO
		{URL: "https://youtube.com/clip/Ugkx1pYKaM52emqwVeDOZoi_NEdh5kyTcTZM?si=VNF2yAuN6sc4O6Ho", Interval: "00:00:0.3-inf"},

		// // CHROMIUM BROO
		// {URL: "./videos-cache/Ugkx1pYKaM52emqwVeDOZoi_NEdh5kyTcTZM_00-00-00_inf.mp4", Interval: "00:00:0.5-inf"},

		// it's a pretty cool idea
		{URL: "https://youtube.com/clip/UgkxoOpVfYh_xO_293qfU5okeevHwDxu-Sxz?si=te49aa-5tXSgZ8rJ"},

		// FREACKING CACHE
		{URL: "https://youtube.com/clip/UgkxOoj_TdSpwkQI7B1Wtwf5JUMVuXwQnAY5?si=ptyOtjQaNX2r7ACi"},

		// Web deved in my ass
		{URL: "https://youtube.com/clip/UgkxVhGJLt_aEN4yLGtKAvi5EdIanHA_usTm?si=SEqfGX_s61EEZVV4"},

		// Ooga booga dev
		{URL: "https://www.youtube.com/clip/UgkxrKlpvWmM8VXrFiDfXgYXbSewx-r0wPUF", Interval: "00:00:00-00:00:6"},

		// Ooga booga dev
		{URL: "https://youtube.com/clip/Ugkx1kaFJj5BRgTSS3nluS0ufyf9BvPVVanQ?si=MTPPQxnMKMXO-Luf"},

		// Feel reckt now chat
		{URL: "https://youtube.com/clip/UgkxMopu-4Fp9vjNCzEdwlUelT3sROR62-pM?si=qCLcjc7XIF7MedLb"},

		// Do you feel Rekt chat?
		{URL: "https://youtube.com/clip/UgkxhXFINlHoz0yCRKLVgv_D5DkCVT5EWj1u?si=tS8cVzVgWgvTozRE"},

		// WHY DOESN"T IT WORK WHO MADE THIS LANGUAGE (linux creators did)
		{URL: "https://youtube.com/clip/UgkxZwLywmZRfrcGmAI8kWMOmV3R9oV23lyL?si=7z99lqY6GfyZUCsX"},

		// bbc penetration testing
		{URL: "https://www.twitch.tv/tsoding/clip/FinePerfectSpindleCurseLit-SHUdEcxbgn6QOWOw", Interval: "00:00:00-00:00:08"},

		// It feels so nice
		{URL: "https://youtube.com/clip/UgkxY0GjOlrBcA904yDruRRj6DtWkzhicNei?si=0CGJ_CbTI8CwisRZ"},

		// Cliteral
		{URL: "https://youtube.com/clip/Ugkxdbyqx5zUrk8ycq3bvxzAK9XxAA4cjGtl?si=Rhe1jHzzK3Gk4VA-"},

		// DX
		{URL: "https://youtube.com/clip/UgkxBoGRnOL9FQ62yrM6mN8gDqsB3G4Rjv_8?si=5dTvFSdBxYspsqaf"},

		// dont touch that mate
		{URL: "https://youtube.com/clip/UgkxuAAQOVtNJ6wd9nHT1q_kS2zkJnCrnOvL?si=o8UX4cdiUNZw0CUG"},

		// What's the point of Po...
		{URL: "https://youtube.com/clip/UgkxeS200OXvc9HcNoCiKVA_SiAl0gbXSNmp"},

		// Femboy when?t
		{URL: "https://youtube.com/clip/UgkxKSOFyBAZjcAifFT6i0d9jsQtU4fbeGPe?si=rmrLcGkiEIbFNdB9"},

		// suffer for RUST
		{URL: "https://youtube.com/clip/Ugkx5gyXHRoeNBCFAxcITsDygUsp7HcSo5ek?si=nu4hLzhAsjilb9RA"},

		// Rust, you are being unethical
		{URL: "https://youtube.com/clip/UgkxWy1-DxpZElAr2E9SeVbkyYBI42xSZ5MI?si=x0-5T-OjY38ACmmD"},

		// old config windows kekw
		{URL: "https://youtube.com/clip/UgkxIRL8T73OpSw2a_cCyUskXx-6RfNM_E1f?si=NazpNKH67DDh1mMz"},

		// Fuckage
		{URL: "https://youtube.com/clip/UgkxJBAp2MTW7kDfhVDDTRy8VB8ZxHKa8I4L?si=UKpouJB8jHU2_beO"},

		// We have a SUCK IT
		{URL: "https://www.twitch.tv/tsoding/clip/PoliteInspiringSoymilkSMOrc-g2rB8-_fiwBPsZx3"},

		// Raw Dog the Posix API like urmom
		{URL: "https://www.twitch.tv/tsoding/clip/VenomousHedonisticFlamingoBlargNaut-Pvfn8OejRMKw7lX6"},

		// tsoding sad :')
		{URL: "https://www.twitch.tv/tsoding/clip/FamousBlazingWasabiPoooound-c9YGTTHT3VpWQWQu", Interval: "00:00:00-00:00:48"},

		// FUTURE DETROYED IN 1 LINE OF CODE
		{URL: "https://www.twitch.tv/tsoding/clip/OptimisticBumblingSmoothiePeteZaroll-KRrjSWgiAga6NQsk"},

		// Can you rust do that? I DON"T FUCKING THINK SO
		{URL: "https://www.twitch.tv/tsoding/clip/AnimatedThankfulNoodleFUNgineer-fXFBX1_duH7kQOQW"},

		// Rust is forcing you to write more readable code
		{URL: "https://www.twitch.tv/tsoding/clip/RacyExuberantDeerSmoocherZ-kJi4i8mPnJxT9wu7"},

		// Best video yet
		{URL: "https://www.youtube.com/clip/Ugkx-YS341MAc8u2ZW5yJCsATJ5W9qpPXNsQ"},

		// // Shader is working shared is twerking
		// {URL: "https://youtube.com/clip/UgkxoA7v4eTxTgO5P4i-NaEw4SHQNIJ0MKvG?si=27nXd6D7LW2LyATr", Interval: "00:00:00-00:00:04"},

		// Thank you for watching
		{URL: "https://www.twitch.tv/tsoding/clip/MagnificentFuriousSquidNomNom-04PJ2wAGlvhWHt3u", Interval: "00:00:0.2-inf"},
	}

	videoCompiler := pkg.NewVideoCompiler(pkg.VideoCompilerConfig{
		CacheFolder:     "./videos-cache-high",
		CacheFileType:   "mp4",
		MaxVideoQuality: "1080",
	})

	err := videoCompiler.AddVideoSources(videoSource...)
	if err != nil {
		panic(err)
	}

	_, err = videoCompiler.MergeVideos("finalVideo.mp4")
	if err != nil {
		panic(err)
	}

}
