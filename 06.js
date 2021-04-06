const list_letters_first = ["c","d","e","g","h"];
const list_letters_second = ["X","Z"];

const findMissingLetter = (str => {
    let alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
    let startPoint = alphabet.indexOf(str[0]);
    for (let i = 0; i < str.length; i++) {
        if(str[i] !== alphabet[startPoint + i]) {
            return alphabet[startPoint + i];
        }
    }
    return undefined;
});

console.log("list_letters_first = ", findMissingLetter(list_letters_first));
console.log("list_letters_second = ", findMissingLetter(list_letters_second));