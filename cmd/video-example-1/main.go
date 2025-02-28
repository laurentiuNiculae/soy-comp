package main

import (
	"log"
	"test/pkg"
)

func main() {
	log.SetFlags(log.Llongfile)

	videoSource := []pkg.VideoSource{
		// tsoding dance
		{URL: "https://www.twitch.tv/videos/2331596519", Interval: "00:00:00-00:00:05.1"},

		// ooga booga software developer
		{URL: "https://youtube.com/clip/UgkxL8ZYx8HKI4NissPiScboPm6d03NI-QME", Interval: ""},

		// beatbox
		{URL: "https://www.twitch.tv/tsoding/clip/TalentedRichGerbilNotATK-z2Rx_N5IvvAuVTX2", Interval: ""},

		// // Exiting your wife
		// {URL: "https://www.twitch.tv/tsoding/clip/GiftedManlyTortoisePJSugar-6WZzYNdaGttYbjFw", Interval: ""},

		// // song followed by ?
		// {URL: "https://www.twitch.tv/tsoding/clip/CleverAbstruseEggplantTBTacoLeft-deCsaUwm52wrIZuo", Interval: ""},

		// // tsodin vomit
		// {URL: "https://www.twitch.tv/tsoding/clip/NastyLivelyClipsdadSmoocherZ-Rj7pTLHhXEajvwlH", Interval: ""},

		// // Shreck scares zozin
		// {URL: "https://www.twitch.tv/tsoding/clip/SavageSuperGrassRaccAttack-F7Y15rRJMT6odDTV", Interval: ""},

		// // Fuck you you bamboozeled me
		// {URL: "https://www.twitch.tv/tsoding/clip/LivelyRoundApeKappaRoss-5cjxdNUsTjvi1j4x", Interval: ""},

		// // I can smell a piece of shaisu
		// {URL: "https://www.twitch.tv/tsoding/clip/AmusedBlueBunnyPupper-WoNMXC--rEeHJXz2", Interval: ""},

		// // bitbox lesson
		// {URL: "https://www.twitch.tv/tsoding/clip/TemperedGlamorousSpiderMikeHogu-u0r0y1y_lKQ_BDAP", Interval: ""},

		// // can your vim do that
		// {URL: "https://www.twitch.tv/tsoding/clip/BillowingBreakableToadGingerPower-eYreTYToWPRBta90", Interval: ""},

		// // i'm not against violence
		// {URL: "https://www.twitch.tv/tsoding/clip/TastyFrozenMilkDancingBanana-4pyJFGT0Ws3fbIPP", Interval: ""},

		// // OPERATING SYSTEM POG
		// {URL: "https://www.twitch.tv/tsoding/clip/HorribleTangentialSquirrelJonCarnage-XXmoAa0ATizVe8VG", Interval: ""},

		// // Zenzin
		// {URL: "https://www.twitch.tv/tsoding/clip/HappyObliviousMeerkatFreakinStinkin-v9p2tOYAVY_pWb38", Interval: "00:00:00-00:00:26.7"},

		// // Tsoding goes into shreks ass
		// {URL: "https://www.twitch.tv/tsoding/clip/EagerBrightBunnyPermaSmug-aXzc4xtrAngUXVis", Interval: ""},

		// // next big thing
		// {URL: "https://www.twitch.tv/tsoding/clip/GorgeousPoorReindeerFunRun-kOu4U-5SwVegm_CO", Interval: "00:00:00-00:00:15.4"},

		// // zoomer zoin
		// // {URL: "https://www.twitch.tv/tsoding/clip/HorriblePrettyTildeTheThing-1PSW9RG6q2k_EdiW", Interval: ""},

		// // tsoding dance
		// {URL: "https://www.twitch.tv/tsoding/clip/AverageHyperGrasshopperUnSane-CfrRSdeYPucEzTaF", Interval: "00:00:00-00:00:11"},

		// // blazingly fast weee
		// // {URL: "https://www.twitch.tv/tsoding/clip/InnocentDifficultPassionfruitMingLee-JIbKL528PvzIvRCC", Interval: ""},

		// // I tried to extract the essence of cock
		// {URL: "https://www.twitch.tv/tsoding/clip/InquisitiveToughAlpacaTwitchRaid-yQJAC-nG_njsLvTI", Interval: "00:00:00-00:00:6.88"},

		// // thank you for the suck
		// {URL: "https://www.twitch.tv/tsoding/clip/AuspiciousWonderfulPeanutTheTarFu-pVhV-m6kH0iZRKJe", Interval: "00:00:00-00:00:7.5"},

		// // Intimate with the OS
		// {URL: "https://www.twitch.tv/tsoding/clip/ExquisiteInspiringGaurBudBlast-0gvFyYgAa8Ghb9-G", Interval: ""},

		// // sex doesn't work
		// {URL: "https://www.twitch.tv/tsoding/clip/JollyCourteousGrasshopperRaccAttack-CDwyCA7MfgUjqV6h", Interval: ""},

		// // do the JVM
		// {URL: "https://www.twitch.tv/tsoding/clip/BreakableReliableRadicchioYee-M6yBGp58Jmc8RZCs", Interval: ""},

		// // JEJEJEJEJEBAITED
		// {URL: "https://www.twitch.tv/tsoding/clip/DifficultEndearingKaleBrainSlug-Mzn3vlQqEJ_poqFt", Interval: ""},

		// // tsoding crazy vscode
		// {URL: "https://www.twitch.tv/tsoding/clip/ViscousKnottyGrassBudStar-W5_r3eldQPPWljI8", Interval: "00:00:12-00:00:38.5"},

		// // bitbox from reading
		// {URL: "https://www.twitch.tv/tsoding/clip/CrazyReliableNoodleDxCat-z-TvQULPaxhmhlnq", Interval: "00:00:8.3-00:00:20"},

		// // bitbox while writing
		// {URL: "https://www.twitch.tv/tsoding/clip/RealEnthusiasticKathyKlappa-Z1moUb-WNBT7Vtyf", Interval: "00:00:0-00:00:20"},

		// // just rewrite this shit in C
		// {URL: "https://www.twitch.tv/tsoding/clip/ShyIgnorantTruffleYouWHY-KO3HeuxBUoEIrqAa", Interval: ""},

		// // bitbox while writing
		// {URL: "https://www.twitch.tv/tsoding/clip/SillyMoralHorseradishUncleNox-s6p_QUriXQAhvpOR", Interval: "00:00:9.1-inf"},
	}

	videoCompiler := pkg.NewVideoCompiler(pkg.VideoCompilerConfig{
		CacheFolder:     "../../cached-videos",
		CacheFileType:   "mp4",
		MaxVideoQuality: "360",
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
