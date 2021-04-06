const text_1 = "Mammals",
    text_2 = "Bruiser build";

const countWords = (words) => {
    let chars = {}, 
        minimize = words.toLowerCase();
        filters = minimize.split("").filter(function(str) {
            return /\S/.test(str);
        });
        filters.forEach(e => {
            if (chars[e]) {
                chars[e] = chars[e]+"*";
            } else {
                chars[e] = "*";
            }
        });
        return chars;
};

console.log(JSON.stringify(countWords(text_1)));
console.log(JSON.stringify(countWords(text_2)));