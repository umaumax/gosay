package gosay

import (
	"fmt"
	"os/exec"
)

type Cmd struct {
	Voice string
	Rate  float64
}

func NewCmd() *Cmd {
	return &Cmd{}
}
func NewCmdV(voice string) *Cmd {
	return &Cmd{
		Voice: voice,
	}
}
func NewCmdVR(voice string, rate float64) *Cmd {
	return &Cmd{
		Voice: voice,
		Rate:  rate,
	}
}

func (c *Cmd) SetVoice(voice string) {
	c.Voice = voice
}
func (c *Cmd) SetRate(rate float64) {
	c.Rate = rate
}

func (c Cmd) Do(message string) (err error) {
	args := make([]string, 0)
	if c.Voice != "" {
		args = append(args, "-v", c.Voice)
	}
	if c.Rate > 0.0 {
		args = append(args, "-r", fmt.Sprint(c.Rate))
	}
	args = append(args, message)
	_, err = exec.Command("say", args...).CombinedOutput()
	fmt.Println(err)
	return
}

func Do(message string) (err error) {
	_, err = exec.Command("say", message).CombinedOutput()
	return
}

func Song() error {
	return NewCmdV("Good").Do(`oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo
oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo
oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo
oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo
oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo`)
}
