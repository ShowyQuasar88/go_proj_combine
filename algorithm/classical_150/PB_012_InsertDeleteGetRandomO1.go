package classical_150

import "math/rand"

type RandomizedSet struct {
	location map[int]int
	values   []int
	length   int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		location: make(map[int]int),
		values:   make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exist := this.location[val]
	if exist {
		return false
	}
	if this.length == len(this.values) {
		this.values = append(this.values, val)
	} else {
		this.values[this.length] = val
	}
	this.location[val] = this.length
	this.length++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exist := this.location[val]
	if !exist {
		return false
	}
	if this.length != 1 && this.values[this.length-1] != val {
		this.values[idx] = this.values[this.length-1]
		this.location[this.values[idx]] = idx
	}
	delete(this.location, val)
	this.length--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(this.length)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
