controller "logoBumpin" {
	default_anim = "bump"

	animation "bump" {
		frames = fromPrefix(PATH, "logo bumpin")
		frame_duration = DUR_24FPS
	}
}
