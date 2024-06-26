package serviceSelector

import (
	"context"
	"errors"
	"lifeChecker/checkers"
	"lifeChecker/database"
	"lifeChecker/tests/mocks"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type ctxKey string

type testCase struct {
	service checkers.LifeChecker
	result  database.TimeSerie
}

type dbMock struct {
	testCases map[string]testCase
}

func (db *dbMock) WriteTimeSerie(t database.TimeSerie) error {
	return nil
}
func (db *dbMock) WriteTimeSerieFromChannel(c context.Context, ch <-chan database.TimeSerie) error {
	var t *testing.T = c.Value(ctxKey("testing")).(*testing.T)
loop:
	for {
		select {
		case i := <-ch:
			res := db.testCases[i.Name].result
			assert.Equal(t, res.Alive, i.Alive)
		case <-c.Done():
			break loop
		}
	}

	return c.Err()
}

func (db *dbMock) Connect() error {
	return nil
}

func (db *dbMock) CloseConnection() error {
	return nil
}

func TestSelector(t *testing.T) {
	t.Run("Concurrency Supporting", func(t *testing.T) {
		selector := NewSelector()
		wg := sync.WaitGroup{}
		for i := range 10 {
			wg.Add(1)
			go func(value int) {
				defer wg.Done()
				selector.Insert(&mocks.Mock_checkLifeService{
					Name: strconv.Itoa(value),
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

func TestRunningSelector(t *testing.T) {
	var testCases = map[string]testCase{
		"Successful Check Life Process": {
			service: &mocks.Mock_checkLifeService{
				Name:     "Successful Check Life Process",
				Inverted: false,
				Err:      nil,
			},
			result: database.TimeSerie{
				Name:        "Successful Check Life Process",
				RequestTime: 100,
				Alive:       true,
			},
		},
		"Successful Inverted Check Life Process": {
			service: &mocks.Mock_checkLifeService{
				Name:     "Successful Inverted Check Life Process",
				Inverted: true,
				Err:      errors.New("Something went wrong"),
			},
			result: database.TimeSerie{
				Name:        "Successful Inverted Check Life Process",
				RequestTime: 100,
				Alive:       true,
			},
		},
		"Unsuccessful Life Check": {
			service: &mocks.Mock_checkLifeService{
				Name:     "Unsuccessful Life Check",
				Inverted: false,
				Err:      errors.New("Unsuccessful Life Check"),
			},
			result: database.TimeSerie{
				Name:        "Unsuccessful Life Check",
				RequestTime: 100,
				Alive:       false,
			},
		},
	}

	db := dbMock{
		testCases: testCases,
	}

	s := NewSelector()

	for _, v := range testCases {
		s.Insert(v.service)
	}

	base := context.WithValue(context.Background(), ctxKey("testing"), t)
	ctx, cancel := context.WithCancel(base)

	go s.RunChecking(ctx, &db)

	time.Sleep(6 * time.Second)

	cancel()
}
