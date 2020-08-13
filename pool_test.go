package goruntinepool

import (
	"fmt"
	"testing"
)

func TestGOR(t *testing.T) {
	p := NewPool(13)

	for i := 0; i < 10; i++ {
		idx := i
		p.Run(func() error {
			fmt.Printf("job %d", idx)
			if idx == 8 {
				return fmt.Errorf("haha error")
			}
			return nil
		})
	}

	p.Wait()
}
