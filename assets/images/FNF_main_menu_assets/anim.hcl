AnimController "story mode" {
	DefaultAnim = "idle"

	Animation "idle" {
		Frames = fromPrefix(PATH, "story mode basic")
		FrameDuration = DUR_24FPS
	}

	Animation "selected" {
		Frames = fromPrefix(PATH, "story mode white")
		FrameDuration = DUR_24FPS
	}
}

AnimController "freeplay" {
	DefaultAnim = "idle"

	Animation "idle" {
		Frames = fromPrefix(PATH, "freeplay basic")
		FrameDuration = DUR_24FPS
	}

	Animation "selected" {
		Frames = fromPrefix(PATH, "freeplay white")
		FrameDuration = DUR_24FPS
	}
}

AnimController "donate" {
	DefaultAnim = "idle"

	Animation "idle" {
		Frames = fromPrefix(PATH, "donate basic")
		FrameDuration = DUR_24FPS
	}

	Animation "selected" {
		Frames = fromPrefix(PATH, "donate white")
		FrameDuration = DUR_24FPS
	}
}

AnimController "options" {
	DefaultAnim = "idle"

	Animation "idle" {
		Frames = fromPrefix(PATH, "options basic")
		FrameDuration = DUR_24FPS
	}

	Animation "selected" {
		Frames = fromPrefix(PATH, "options white")
		FrameDuration = DUR_24FPS
	}
}
