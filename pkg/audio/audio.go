package audio

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"layla/pkg/assets"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	suffixes = []string{".mp3", ".ogg", ".wav"}

	audioContext *audio.Context

	currentBgm *audio.Player = nil
	bgmPlayers               = map[string]*audio.Player{}
	seBytes                  = map[string][]byte{}

	mute = false

	bgmKeys []string
	seKeys  []string
)

func bgm(slice map[string]*audio.Player) []string {
	keys := make([]string, 0, len(slice))
	for key := range slice {
		keys = append(keys, key)
	}
	return keys
}

func IsMute() bool {
	return mute
}

func SetMute(enabled bool) {
	mute = enabled
}

func init() {
	bgm, _ := assets.AudioFS.ReadDir("audio/bgm")
	for _, entry := range bgm {
		if !entry.IsDir() {
			name := entry.Name()
			bgmKeys = append(bgmKeys, name)
		}
	}
	sfx, _ := assets.AudioFS.ReadDir("audio/sfx")
	for _, entry := range sfx {
		if !entry.IsDir() {
			name := entry.Name()
			seKeys = append(seKeys, name)
		}
	}
}

func init() {
	const sampleRate = 44100
	audioContext = audio.NewContext(sampleRate)
}

func Load() error {
	soundDirs := []string{
		"bgm",
		"sfx",
	}
	for _, dir := range soundDirs {
		filenames := make([]string, 0)
		switch dir {
		case "bgm":
			filenames = bgmKeys
		case "sfx":
			filenames = seKeys
		}
		for _, n := range filenames {
			b, err := assets.AudioFS.ReadFile("audio/" + dir + "/" + n)
			if err != nil {
				return err
			}
			var stream io.ReadSeeker
			var length int64
			switch {
			case strings.HasSuffix(n, ".mp3"):
				stream, err = mp3.Decode(audioContext, bytes.NewReader(b))
				if err != nil {
					return err
				}
				length = stream.(*mp3.Stream).Length()
			case strings.HasSuffix(n, ".ogg"):
				stream, err = vorbis.Decode(audioContext, bytes.NewReader(b))
				if err != nil {
					return err
				}
				length = stream.(*vorbis.Stream).Length()
			case strings.HasSuffix(n, ".wav"):
				stream, err = wav.Decode(audioContext, bytes.NewReader(b))
				if err != nil {
					return err
				}
				length = stream.(*wav.Stream).Length()
			default:
				panic("invalid file name")
			}
			switch dir {
			case "bgm":
				stream = audio.NewInfiniteLoop(stream, length)
			case "sfx":
				// stream = stream
			}
			p, err := audioContext.NewPlayer(stream)
			if err != nil {
				return err
			}
			switch dir {
			case "bgm":
				bgmPlayers[n] = p
			case "sfx":
				b, err := ioutil.ReadAll(stream)
				if err != nil {
					return err
				}
				seBytes[n] = b
			}
		}
	}

	return nil
}

func Finalize() error {
	var soundPlayers map[string]*audio.Player
	soundPlayers = bgmPlayers
	for _, p := range soundPlayers {
		if err := p.Close(); err != nil {
			return err
		}
	}
	return nil
}

func SetBGMVolume(volume float64) {
	if mute {
		return
	}
	for _, p := range bgmPlayers {
		if !p.IsPlaying() {
			continue
		}
		p.SetVolume(volume)
		return
	}
}

func PauseBGM() {
	if mute {
		return
	}
	for _, p := range bgmPlayers {
		p.Pause()
	}
}

func StopBGM() {
	PauseBGM()
	for _, p := range bgmPlayers {
		p.Pause()
		p.Rewind()
	}
}

func ResumeBGM(bgm string) {
	if mute {
		return
	}
	PauseBGM()
	p := bgmPlayers[bgm]
	if p == nil {
		err := fmt.Errorf(`BGM "%s" doesn't found on memory`, bgm)
		panic(err)
	}
	p.SetVolume(1)
	p.Play()
}

func PlayBGM(bgm string) error {
	if mute {
		return nil
	}
	PauseBGM()
	p := bgmPlayers[bgm]
	if p == nil {
		err := fmt.Errorf(`BGM "%s" doesn't found on memory`, bgm)
		panic(err)
	}
	p.SetVolume(1)
	if err := p.Rewind(); err != nil {
		return err
	}
	p.Play()
	return nil
}

func PlaySE(se string) {
	if mute {
		return
	}
	p := seBytes[se]
	if p == nil {
		err := fmt.Errorf(`Sound Effect "%s" doesn't found on memory`, se)
		panic(err)
	}
	sePlayer := audio.NewPlayerFromBytes(audioContext, seBytes[se])
	// sePlayer is never GCed as long as it plays.
	sePlayer.Play()
}
