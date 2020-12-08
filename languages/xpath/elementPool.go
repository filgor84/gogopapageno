package xpath

type pool struct {
	pool        []interface{}
	cur         int
	constructor func() interface{}
}

func newPool(length int, constructor func() interface{}) *pool {
	p := pool{make([]interface{}, length), 0, constructor}

	for i := 0; i < length; i++ {
		p.pool[i] = constructor()
	}

	return &p
}

func (p *pool) Get() interface{} {
	if p.cur >= len(p.pool) {
		return p.constructor()
	}
	addr := p.pool[p.cur]
	p.cur++
	return addr
}

func (p *pool) Remainder() int {
	return len(p.pool) - p.cur
}
