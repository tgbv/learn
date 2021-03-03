class Node {
	constructor(value){
		this.value = value
		this.left = null
		this.right = null
	}
}


class BTree {
	constructor(){
		this.root = null
	}

	insert(value)
	{
		// create the node 
		let n = new Node(value)

		// first we check if root is null
		if(this.root === null)
			this.root = n

		// otherwise
		else {
			// iterate tree
			// assign value to left or right
			// the lower value goes to left, while the bigger/equal to the right
			let currentNode = this.root

			// infinite iteration
			// I want to traverse the tree
			while(true)
			{
				if(value < currentNode.value)
				{
					// if there is a left node, loop IT instead
					if(currentNode.left)
						currentNode = currentNode.left

					// if there is no left node, allocate new value to it
					else
					{
						currentNode.left = n
						return this			
					}

				}
				else
				{
					// see upper comment
					if(currentNode.right)
						currentNode = currentNode.right
					else {
						currentNode.right = n
						return this				
					}

				}
			}
			
		}

		return this
	}

	lookup(value){
		// check if root exists
		if(this.root)
		{
			let currentNode = this.root

			// iterate until current node contains the value (or not)
			while(true)
			{
				if(value < currentNode.value)
				{
					if(currentNode.left)
						currentNode = currentNode.left
					else return false
				}
				else if( value > currentNode.value)
				{
					if(currentNode.right)
						currentNode = currentNode.right 
					else return false
				}
				else if(currentNode.value === value)
					return currentNode
			}
		}

		else return false;
	}

	remove(value){
		// check if root exists
		if(this.root)
		{
			// 
		}
		else  return this
	}

	traverseTree(){}
}


let tree = new BTree

tree.insert(10).insert(8).insert(20).insert(11).insert(21)

console.log(tree.lookup(20))