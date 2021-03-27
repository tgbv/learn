const num = [33, 22, 11, 55, 32, 76, 324, 5]


const sort = function(arr) {
    let tmp;
    let smallest;
    let smallestIndex;

    for(let i=0; i<arr.length; i++){
        smallest = arr[i]
        smallestIndex = i
        
        for(let j=i; j<arr.length; j++)
        {
            if(smallest > arr[j]){
                smallest = arr[j]
                smallestIndex = j
            }
        }

        tmp = arr[i]
        arr[i] = smallest
        arr[smallestIndex] = tmp
    }

    return arr
}

console.log(sort(num))