package sounds

import (
	"fmt"
	"os"
	"sync"
	"time"

	"goau-biat/config"

	"github.com/ebitengine/oto/v3"
	mp3 "github.com/hajimehoshi/go-mp3"
)

type SoundFile string

const UltimateReloaded SoundFile = "ultimate-reloaded"
const PreUltimateReloaded SoundFile = "pre-ultimate-reloaded"
const CheckListLostSouls SoundFile = "checklist-lost-souls"
const CheckListGoannas SoundFile = "checklist-goannas"
const Paused SoundFile = "paused"
const Resumed SoundFile = "resumed"

var (
	audioCtx  *oto.Context
	audioOnce sync.Once
)

func ensureContext(sampleRate int) {
	audioOnce.Do(func() {
		c, ready, err := oto.NewContext(&oto.NewContextOptions{
			SampleRate:   sampleRate,
			ChannelCount: 2,
			Format:       oto.FormatSignedInt16LE,
		})
		if err != nil {
			return
		}
		<-ready
		audioCtx = c
	})
}

func PlaySound(file SoundFile) error {
	path := fmt.Sprintf("C:\\Users\\kelvi\\workspace\\goau-biat\\sounds\\%s\\%s.mp3", config.Lang, file)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	ensureContext(d.SampleRate())

	if audioCtx == nil {
		return fmt.Errorf("audio context not initialized")
	}

	p := audioCtx.NewPlayer(d)

	p.Play()

	for p.IsPlaying() {
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}
