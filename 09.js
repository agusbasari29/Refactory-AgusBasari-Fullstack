const arrayList = [3, 12, 4, 5, 8, 9]
const sortMethod =(items)=>{
	items.sort(function(a, b) {
        return a - b;
    });
    console.log(items);
};

sortMethod(arrayList)