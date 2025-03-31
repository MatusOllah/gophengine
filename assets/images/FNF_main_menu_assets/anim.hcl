controller "story mode" {
	default_anim = "idle"

	animation "idle" {
		frames = fromPrefix(PATH, "story mode basic")
		frame_duration = DUR_24FPS
	}

	animation "selected" {
		frames = fromPrefix(PATH, "story mode white")
		frame_duration = DUR_24FPS
	}
}

controller "freeplay" {
	default_anim = "idle"

	animation "idle" {
		frames = fromPrefix(PATH, "freeplay basic")
		frame_duration = DUR_24FPS
	}

	animation "selected" {
		frames = fromPrefix(PATH, "freeplay white")
		frame_duration = DUR_24FPS
	}
}

controller "donate" {
	default_anim = "idle"

	animation "idle" {
		frames = fromPrefix(PATH, "donate basic")
		frame_duration = DUR_24FPS
	}

	animation "selected" {
		frames = fromPrefix(PATH, "donate white")
		frame_duration = DUR_24FPS
	}
}

controller "options" {
	default_anim = "idle"

	animation "idle" {
		frames = fromPrefix(PATH, "options basic")
		frame_duration = DUR_24FPS
	}

	animation "selected" {
		frames = fromPrefix(PATH, "options white")
		frame_duration = DUR_24FPS
	}
}
