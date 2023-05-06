package sounds

import (
	"fmt"
	"goau-biat/config"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

type SoundFile string

const UltimateReloaded = "ultimate-reloaded"
const PreUltimateReloaded = "pre-ultimate-reloaded"
const CheckListLostSouls = "checklist-lost-souls"
const CheckListGoannas = "checklist-goannas"
const Paused = "paused"
const Resumed = "resumed"

func PlaySound(file SoundFile) error {
	path := fmt.Sprintf("/home/kelvin/workspace/goau-biat/sounds/%s/%s.mp3", config.Lang, file)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}

	return nil
}
