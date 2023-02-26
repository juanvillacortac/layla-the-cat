package data

type LevelSavedData struct {
	Unlocked bool `json:"unlocked"`
	Cleared  bool `json:"cleared"`
}

type SaveFileData struct {
	Levels          map[int]*LevelSavedData `json:"levels"`
	LastPlayedLevel int                     `json:"last_played"`
}

var Levels []string = []string{
	"grass",   // 0
	"another", // 1
	"grass",   // 2
	"another", // 3
	"grass",   // 4
	"another", // 5
	"grass",   // 6
	"another", // 7
}

var SavedData = SaveFileData{}

func DefaultData() SaveFileData {
	data := SaveFileData{}
	data.Levels = make(map[int]*LevelSavedData)
	for idx := range Levels {
		data.Levels[idx] = &LevelSavedData{
			Unlocked: false,
		}
	}
	data.Levels[0] = &LevelSavedData{
		Unlocked: true,
	}
	return data
}

func init() {
	SavedData = DefaultData()
}
