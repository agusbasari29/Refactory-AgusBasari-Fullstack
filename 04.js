// Create row of numbers
let numbers = [];
for (let i=0; i<1001; i++) {
    numbers[i] = i;
}

// grouping the numbers by even and odd
let even = [];
let odd = [];
numbers.forEach(e => {
    if(e % 2 == 0){
        even.push(e);
    } else {
        odd.push(e);
    }
});

//grouping the nummbers multiplies by 5
let multipliesByFive = [];
numbers.forEach(e => {
    if(e % 5 == 0) {
        multipliesByFive.push(e);
    }
});


const isPrime = (n) => {
    for (let j = 2; j < n; j++) {
        if(n % j === 0) return false;
        return n > 1;
    }
}

// grouping by prime numbers
let primeNumbers = [];
numbers.forEach(e => {
    const check = isPrime(e);
    if (check) {
        primeNumbers.push(e);
    }
});

// grouping by prime numbers less than 100
let primeNumbersLessThanOneHundred = [];
for (let k=0; k < 100; k++) {
    const check = isPrime(numbers[k]);
    if (check) {
        primeNumbersLessThanOneHundred.push(numbers[k]);
    }
}

let results = {};
results.numbers = numbers;
results.even = even;
results.odd = odd;
results.primeNumbers = primeNumbers;
results.primeNumbersLessThanOneHundred = primeNumbersLessThanOneHundred;
console.log(results);