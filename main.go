package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
)

type LimitData struct {
	Econom   string
	Comfort  string
	Business string
}
type Cabinet string
type User struct {
	ID      string
	LimitID string
}

type LimitsData struct {
	data map[Cabinet]LimitData
	mx   *sync.RWMutex
}
type Service struct {
	limitsData LimitsData
}

func New() *Service {
	return &Service{
		limitsData: LimitsData{
			data: map[Cabinet]LimitData{},
			mx:   &sync.RWMutex{},
		},
	}
}

func (s *Service) getActualLimitsData(c Cabinet) (LimitData, error) {
	return LimitData{
		Econom:   uuid.NewString(),
		Comfort:  uuid.NewString(),
		Business: uuid.NewString(),
	}, nil
}

func (s *Service) RefillLimitsData(cabinets []Cabinet) {
	wg := &sync.WaitGroup{}

	for _, c := range cabinets {
		wg.Add(1)

		go func(c Cabinet) {
			defer wg.Done()

			actualLimitData, err := s.getActualLimitsData(c)
			if err != nil {
				log.Printf("got some error: %v", err)
			}
			fmt.Println(actualLimitData)

			s.limitsData.mx.Lock()
			s.limitsData.data[c] = actualLimitData
			s.limitsData.mx.Unlock()
		}(c)
	}

	wg.Wait()
}

func main() {
	cabinets := []Cabinet{
		"cab01", "cab02", "cab03",
	}

	srvc := New()

	fmt.Println(srvc.limitsData.data)
	srvc.RefillLimitsData(cabinets)
	fmt.Println(srvc.limitsData.data)
}
