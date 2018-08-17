package sleepy

import (
	"testing"
	"time"

)

 func TestTimeSleep(t *testing.T){
	start := time.Now()
	sleep(1)
	elapsed := time.Since(start).Seconds()

	if elapsed < 1 && elapsed > 0.9999 {
		t.Error("Sleep fail")
	}
 }


