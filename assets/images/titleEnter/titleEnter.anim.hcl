AnimController {
    DefaultAnim = "idle"

    Animation "idle" {
        Frames = fromPrefix(PATH, "Press Enter to Begin")
        FrameDuration = "41ms"
    }

    Animation "press" {
        Frames = fromPrefix(PATH, "ENTER PRESSED")
        FrameDuration = "41ms"
    }
}