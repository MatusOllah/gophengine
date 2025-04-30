controller "tutorial" {
	default_anim = "select"

	animation "select" {
		frames = fromPrefix(PATH, "tutorial selected")
		frame_duration = DUR_24FPS
	}
}

controller "week1" {
	default_anim = "select"

	animation "select" {
		frames = fromPrefix(PATH, "WEEK1 select")
		frame_duration = DUR_24FPS
	}
}

controller "week2" {
	default_anim = "select"

	animation "select" {
		frames = fromPrefix(PATH, "week2 select")
		frame_duration = DUR_24FPS
	}
}

controller "week4" {
	default_anim = "select"

	animation "select" {
		frames = fromPrefix(PATH, "Week 4 press")
		frame_duration = DUR_24FPS
	}
}

controller "week6" {
	default_anim = "select"

	animation "select" {
		frames = fromPrefix(PATH, "Week 6")
		frame_duration = DUR_24FPS
	}
}

controller "left_arrow" {
	default_anim = "idle"

	animation "idle" {
		frames = fromPrefix(PATH, "arrow left")
		frame_duration = DUR_24FPS
	}

	animation "press" {
		frames = fromPrefix(PATH, "arrow push left")
		frame_duration = DUR_24FPS
	}
}

controller "difficulty" {
	default_anim = "normal"

	animation "easy" {
		frames = fromPrefix(PATH, "EASY")
		frame_duration = DUR_24FPS
	}

	animation "normal" {
		frames = fromPrefix(PATH, "NORMAL")
		frame_duration = DUR_24FPS
	}

	animation "hard" {
		frames = fromPrefix(PATH, "HARD")
		frame_duration = DUR_24FPS
	}
}

controller "right_arrow" {
	default_anim = "idle"

	animation "idle" {
		frames = fromPrefix(PATH, "arrow right")
		frame_duration = DUR_24FPS
	}

	animation "press" {
		frames = fromPrefix(PATH, "arrow push right")
		frame_duration = DUR_24FPS
	}
}
