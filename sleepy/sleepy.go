package sleepy

import (
	"time"
)


// toy sleep implementation with select
func sleep(seconds int){
	select {
	    case <- time.After(time.Duration(seconds) * time.Second):
		    return
	}
 }
