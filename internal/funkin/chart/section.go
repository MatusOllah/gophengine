package chart

type Section struct {
	SectionNotes   [][]interface{} `json:"sectionNotes"`
	LengthInSteps  int             `json:"lengthInSteps"`
	TypeOfSection  int             `json:"typeOfSection"`
	MustHitSection bool            `json:"mustHitSection"`
	Bpm            int             `json:"bpm"`
	ChangeBPM      bool            `json:"changeBPM"`
	AltAnim        bool            `json:"altAnim"`
}
