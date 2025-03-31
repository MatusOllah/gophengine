controller "titleEnter" {
	default_anim = "idle"

	animation "idle" {
		frames = fromPrefix(PATH, "Press Enter to Begin")
		frame_duration = DUR_24FPS
	}

	animation "press" {
		frames = fromPrefix(PATH, "ENTER PRESSED")
		frame_duration = DUR_24FPS
	}
}
