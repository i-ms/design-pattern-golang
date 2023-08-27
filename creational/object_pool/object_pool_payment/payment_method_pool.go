package object_pool_payment

import "sync"

type PaymentMethodPool struct {
	mu        sync.Mutex
	available []PaymentMethod
	inUse     []PaymentMethod
}

// Acquire returns a payment method from the pool for use
func (p *PaymentMethodPool) Acquire() PaymentMethod {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.available) == 0 {
		return nil
	}

	var pm PaymentMethod
	pm, p.available = p.available[0], p.available[1:]
	p.inUse = append(p.inUse, pm)
	return pm
}

// Release returns a payment method back to the pool
func (p *PaymentMethodPool) Release(pm PaymentMethod) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.available = append(p.available, pm)
	for i, item := range p.inUse {
		if item == pm {
			p.inUse = append(p.inUse[:i], p.inUse[i+1:]...)
			break
		}
	}
}
