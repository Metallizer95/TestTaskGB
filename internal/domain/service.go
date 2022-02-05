package domain

func (h *Holders) AddValue(address string, value int) {
	_, ok := h.Holders[address]
	if !ok {
		h.Holders[address] = value
	} else {
		h.Holders[address] += value
	}
}

func (h *Holders) SubValue(address string, value int) {
	_, ok := h.Holders[address]
	if !ok {
		h.Holders[address] = value
	} else {
		h.Holders[address] -= value
	}
}
