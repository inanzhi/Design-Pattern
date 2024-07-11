package facade_pattern

import "fmt"

// AudioPlayer 音频播放器
type AudioPlayer struct{}

func (a *AudioPlayer) PlayAudio() {
	fmt.Println("Playing audio")
}

// VideoPlayer 视频播放器
type VideoPlayer struct{}

func (v *VideoPlayer) PlayVideo() {
	fmt.Println("Playing video")
}

// SubtitleDisplayer 字幕显示器
type SubtitleDisplayer struct{}

func (s *SubtitleDisplayer) DisplaySubtitle() {
	fmt.Println("Displaying subtitle")
}
