AnimController {
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