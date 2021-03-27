const num = [1,3, 2, 443, 43, 11000, 34, 21, 0]

const sort = function(arr) {
    let len = arr.length

    for(let i=0; i<len; i++)
    {
        if(i === 0) continue

        if(arr[i] < arr[i-1])
            for(let j=0; j<i; j++)
                if( arr[i] > arr[j]) continue
                else {
                    arr.splice(j, 0, arr[i]) // it pushes the element
                    arr.splice(i+1, 1)      // ..therefore the number we want to delete will be an index further, so we splice i+1

                    break;
                }
    }

    return arr
}

console.log(sort(num))