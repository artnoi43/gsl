package concurrent

import (
	"errors"
	"strings"
	"sync"
)

// WaitAndCorrectErrors waits on wg and close errChan once the code moved past wg.Wait().
// It then collect the errors into one big error message joined by ','.
func WaitAndCollectErrors(wg *sync.WaitGroup, errChan chan error) error {
	go func() {
		wg.Wait()
		close(errChan)
	}()
	if errChan != nil {
		var errs []string
		for err := range errChan {
			errs = append(errs, err.Error())
		}
		if len(errs) > 0 {
			return errors.New(strings.Join(errs, ","))
		}
	}
	return nil
}
