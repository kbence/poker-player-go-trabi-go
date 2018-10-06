package main

type PreFlopDecider struct{}

func (p *PreFlopDecider) Next() Decision {
	return Decision{Decision: CALL}
}
