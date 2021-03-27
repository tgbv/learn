const num = [33, 22, 11, 55, 32, 76, 324, 5]


const sort = function(arr) {

    let tmp
    
    for(let i=0; i<arr.length; i++)
        for(let j=0; j<arr.length; j++)
            if(arr[j+1])
                if(arr[j] > arr[j+1])
                {
                    tmp = arr[j+1]
                    arr[j+1] = arr[j]
                    arr[j] = tmp
                }

    return arr
}

console.log(sort(num) )