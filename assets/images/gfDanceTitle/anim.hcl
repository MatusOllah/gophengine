controller "gfDanceTitle" {
	animation "danceLeft" {
		frames = fromIndices(PATH, "gfDance", [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14])
		frame_duration = DUR_24FPS
	}

	animation "danceRight" {
		frames = fromIndices(PATH, "gfDance", [15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29])
		frame_duration = DUR_24FPS
	}
}
