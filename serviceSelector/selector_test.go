package serviceSelector

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

type mock_checkLifeService struct {
	name string
}

func (s *mock_checkLifeService) GetName() string {
	return s.name
}

func (s *mock_checkLifeService) CheckLife() (time.Duration, error) { return time.Duration(100), nil }
func (s *mock_checkLifeService) IsInverted() bool                  { return true }
func (s *mock_checkLifeService) GetQueueTime() time.Duration       { return time.Duration(0) }

func TestSelector(t *testing.T) {
	t.Run("Concurrency Supporting", func(t *testing.T) {
		selector := NewSelector()
		wg := sync.WaitGroup{}

		for i := range 10 {
			wg.Add(1)
			go func(value int) {
				defer wg.Done()
				selector.Insert(&mock_checkLifeService{
					name: strconv.Itoa(value),
				})
			}(i)
		}

		wg.Wait()

		for range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				selector.NextItem()
			}()
		}

		wg.Wait()
	})
}
