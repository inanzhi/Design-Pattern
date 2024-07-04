package mediator_pattern

import "sync"

// 车站管理员
type StationManager struct {
	//平台可停靠
	isPlatformFree bool
	lock           *sync.Mutex
	//火车队列
	trainQueue []train
}

func NewStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}
func (s *StationManager) canLand(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}
func (s *StationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.StartArrival()
	}
}
