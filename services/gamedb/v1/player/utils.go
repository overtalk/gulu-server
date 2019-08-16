package player

func (p *player) addUpdate(key string, value interface{}) {
	p.update[key] = value
}
