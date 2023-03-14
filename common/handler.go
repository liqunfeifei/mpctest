package common

func HandlerLoop(r Round) {
	for {
		if !r.ReceivedAll() {
			continue
		}
		r.Finalize()
	}
}
