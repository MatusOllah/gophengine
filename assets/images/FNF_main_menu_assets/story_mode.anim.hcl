AnimController {
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