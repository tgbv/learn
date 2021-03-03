class Node {
	constructor(value){
		this.value = value
		this.next = null
	}
}


class Queue {
	constructor(){
		this.first = null
		this.last = null
		this.length = 0
	}


	// push at the tail of queue
	enqueue(value)
	{
		// create new node
		let elem = new Node(value)

		// check if collection is empty
		if(this.isEmpty())
		{
			// assign references to it
			this.first = elem
			this.last = elem
		}
		else
		{
			// update ref of last element
			// ref will point to the new node
			this.last.next = elem

			// update ref of last
			this.last = elem
		}

		// increment length
		this.length++

		return this
	}

	// remove from front of queue
	dequeue(){
		// if empty, return this
 		if(this.isEmpty()) 
 			return this

 		// if one element, remove refs from both first and last
 		else if (this.length === 1)
 		{
 			this.first = null
 			this.last = null
 		}
 		else 
 		{
 			// change the ref of first to next
 			this.first = this.first.next
 		}

 		this.length--

 		return this
	}

	// see the first element of queue
	peek(){
		return this.isEmpty() ? null : this.first.value
	}

	// see if queue is empty
	isEmpty(){
		return this.length === 0
	}
}

// queue

const q = new Queue

q.enqueue("Alice").enqueue("Romeo").enqueue("Neataa")


console.log( q.dequeue().dequeue().peek() )