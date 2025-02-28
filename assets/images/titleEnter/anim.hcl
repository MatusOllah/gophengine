AnimController "titleEnter" {
    DefaultAnim = "idle"

    Animation "idle" {
        Frames = fromPrefix(PATH, "Press Enter to Begin")
        FrameDuration = DUR_24FPS
    }

    Animation "press" {
        Frames = fromPrefix(PATH, "ENTER PRESSED")
        FrameDuration = DUR_24FPS
    }
}
