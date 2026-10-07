type Tuple struct {
	value string
	timestamp int
}
type TimeMap struct {
	dict map[string][]*Tuple
}

func Constructor() TimeMap {
	return TimeMap{dict: make(map[string][]*Tuple)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.dict[key] = append(this.dict[key], &Tuple{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, ok := this.dict[key]
	if !ok {
		return ""
	} 
	l, r := 0, len(values)-1
	for l<r{
		mid := l + (r-l)/2
		if values[mid].timestamp <= timestamp {
			l = mid+1
		} else {
			r = mid
		}
	}
	if values[l].timestamp <= timestamp {
		return values[l].value
	}

	if l > 0 {
		return values[l-1].value
	}

	return ""
}
