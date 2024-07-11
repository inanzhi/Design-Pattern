package facade_pattern

// MediaFacade 门面类，简化音视频子系统的使用
type MediaFacade struct {
	audioPlayer       *AudioPlayer
	videoPlayer       *VideoPlayer
	subtitleDisplayer *SubtitleDisplayer
}

func NewMediaFacade() *MediaFacade {
	return &MediaFacade{
		audioPlayer:       &AudioPlayer{},
		videoPlayer:       &VideoPlayer{},
		subtitleDisplayer: &SubtitleDisplayer{},
	}
}

func (m *MediaFacade) Play() {
	m.videoPlayer.PlayVideo()
	m.audioPlayer.PlayAudio()
	m.subtitleDisplayer.DisplaySubtitle()
}
