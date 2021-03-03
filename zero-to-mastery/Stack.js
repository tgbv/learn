/*
*	holds one node from the stack
*/
class Node {
	constructor(value){
		this.value = value
		this.next = null
	}
}

/*
*	the actual stack management class
*/
class Stack 
{
	/*
	*	constructor
	*/
	constructor(value = null){
		this.top = null
		this.bottom = null
		this.length = 0

		if(value) this.push(value)
	}

	/*
	*	pushes element to stack
	*/
	push(value){
		// first we check to see if the stack is empty
		if(this.isEmpty())
		{
			// if so create the element
			let node = new Node(value)

			// assign ref element to the bottom
			this.bottom = node

			// reference the bottom to the top because it's only one element
			this.top = this.bottom
		}
		else 
		{
			// otherwise create the element
			let node = new Node(value)

			// chain the element with the top node
			node.next = this.top

			// update the top
			this.top = node
		}

		// update length
		this.length++

		return this
	}

	/*
	*	shows the last pushed element
	*/
	peek(){
		return this.top
	}

	/*
	*	removes the last pushed element
	*/
	pop(){
		if(this.isEmpty()) return this

		this.top = this.top.next

		this.length--

		return this
	}

	/*
	*	checks if stack is empty
	*/
	isEmpty(){
		return this.length === 0
	}
}


/*
*	the stack array class
*/
class StackArray {
	constructor(){
		this.stack = []
	}

	// push element
	push(value){
		this.stack.push(value)

		return this
	}

	// see the last pushed element
	peek(){
		if(this.isEmpty()) return null
		else return this.stack[this.stack.length-1]
	}

	// delete the last element
	pop(){
		if(this.isEmpty()) return this
		else {
			this.stack.splice(this.stack.length-1)
			return this
		}
	}

	// reverse
	reverse(){
		this.stack = this.stack.reverse()
		return this
	}

	// checks if stack is empty
	isEmpty(){
		return this.stack.length === 0
	}
}


const stack = new Stack("test")

stack.push("Hello").push("World").push("!").pop()


const stackArray = new StackArray

stackArray.push("Hello").push("World").push("!")

console.log(stackArray)