AnimController {
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