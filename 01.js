const CSVToJSON = require('csvtojson');

CSVToJSON().fromFile("./source.csv").then(source => {
    const sorter = source.sort(function (a, b) {
        return a.PRICE - b.PRICE;
    });
    const result = sorter.map(function(data){
        const priceFormatter = data.PRICE.toString().split('').reverse().join(''),
            separation = priceFormatter.match(/\d{1,3}/g),
            price = 'Rp'+separation.join('.').split('').reverse().join('');
            return {name: data.NAME, price: price, category: data.CATEGORY};
    });
    console.log(result);
});