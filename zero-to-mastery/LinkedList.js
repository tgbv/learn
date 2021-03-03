class LinkedList {
  // instantiates the linkedlist
  // creates basic vars
  constructor(value){
    this.head = {
      value: value,
      next: null,
    }

    this.tail = this.head

    this.length = 1
  }

  // appends new value to the list
  append(value){
    let obj = {
      value: value,
      next: null
    }

    this.tail.next = obj
    this.tail = obj

    //update the length
    this.length++

    // 
    return this
  }

  // prepends a value to this linked list
  prepend(value){
    // create new node 
    let newNode = {
      value: value,
      next: this.head
    }

    // modify reference of head to the new node
    this.head = newNode
    
    // modify length
    this.length++

    // return this because it's useful
    return this
  } 

  // inserts a node at an index
  insert(index, value){
    // check index if it's within boundaries
    if(index < 0 || index > this.length-1)
      throw "Index out of boundaries"


    // if index is 0, prepend
    if(index === 0) return this.prepend(value)

    // if index is this.length-1, append
    else if(index === this.length-1) return this.append(value)

    // otherwise attempt to find the node correlated with the index
    else {
      // holds the looked object
      let lookedObject = null

      // iterate the index 
      for(let i=0; i<index; i++)
        if(i === 0) lookedObject = this.head
        else lookedObject = lookedObject.next

      // insert the new node
      let newObject = {
        value: value,
        next: lookedObject.next
      }
      lookedObject.next = newObject

      // update length
      this.length++

      return this
    }

  }

  // deletes a node from an index
  delete(index){
    // check if index is within boundaries
    if(index < 0 || index > this.length-1)
      throw "Index out of boundaries"

    // if index === 0, delete the head
    if(index === 0) return this.deleteHead()

    else if (index === this.length-1) return this.deleteTail()

    else {
      // temporary holds the pre-index node
      let beforeIndexNode = null;

      // iterate and find node index 
      for(let i=0; i<index; i++)
        if(i === 0) beforeIndexNode = this.head
        else beforeIndexNode = beforeIndexNode.next

      // unlink the index node
      beforeIndexNode.next = beforeIndexNode.next.next
    }

    // update length
    this.length--

    return this
  }

  deleteHead(){
    if(this.head.next)
      this.head = this.head.next
    else 
    {
      this.head = null
      this.tail = null
    }

    this.length--
    return this
  }

  deleteTail(){
    // get tail
    let tmpNode = null

    for(let i=0; i<this.length-1; i++)
      if(i===0) tmpNode = this.head
      else tmpNode = tmpNode.next

    // delete
    this.tail = tmpNode
    tmpNode.next = null

    this.length--

    return this
  }

  /*
  * return a reverse of the linked list
  */
  reversedList() {
    // create new linked list
    const newLinkedList = new LinkedList(this.head.value)

    // temp node
    let tempNode = this.head

    // iterate current linked list and append the remaining values
    for(let i=1; i<this.length; i++)
    {
      tempNode = tempNode.next
      newLinkedList.prepend(tempNode.value)
    }

    // returns the new linked list
    return newLinkedList
  }

  /*
  *	reverse a linked list
  */
  reverseList(){
  	if(!this.head.next)
  		return this;

  	let first = this.head
  	let second = first.next

 	this.tail = first
 	this.tail.next = null

  	while(second)
  	{
  		let third = second.next
  		second.next = first

  		first = second
  		second = third
  	}

  	//this.head.next = null
  	this.head = first

  	return this
  }

  /*
  * get node at index
  */
  getNodeAtIndex(index){
    // check if index is within boundaries
    if( index < 0 || index > this.length-1)
      throw "Out of boundaries"

    // holds nodeIndex
    let nodeIndex = null

    // iterate index 
    for(let i=0; i<=index; i++)
      if(i===0) nodeIndex = this.head
      else nodeIndex = nodeIndex.next

    // returns the node
    return nodeIndex
  }
}


const List = new LinkedList("A")
List.append("B").append("C")


console.log(List.reverseList())

