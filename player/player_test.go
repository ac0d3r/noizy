package player

import (
	"log"
	"testing"
	"time"
)

func TestPlay(t *testing.T) {
	p := New()
	p.Init()

	p.Play("Rain Drops")

	<-time.After(30 * time.Second)
	log.Println("播放完成")

}
