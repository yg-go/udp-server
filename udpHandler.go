package main

type RFSUDPHandler struct {
	dispatcherChan chan<- Payload
}

func (h *RFSUDPHandler) OnMessage(payload []byte) {
	h.dispatcherChan <- payload
}
