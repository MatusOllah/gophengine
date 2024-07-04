AnimController {
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