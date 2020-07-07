package leetcode

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter    *Iterator
	cache   int
	isCache bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:    iter,
		cache:   0,
		isCache: false,
	}
}
func (this *PeekingIterator) hasNext() bool {

	return this.isCache || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if !this.isCache {
		return this.iter.next()
	}
	res := this.cache
	this.isCache = false
	return res
}

func (this *PeekingIterator) peek() int {
	if !this.isCache {
		this.cache = this.iter.next()
		this.isCache = true
	}

	return this.cache
}
