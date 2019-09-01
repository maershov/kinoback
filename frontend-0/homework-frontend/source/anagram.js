'use strict';


let anagram = function(inputArr){

  if (!Array.isArray(inputArr)) { //проверка того, что входной массиы вообще массив
    return [];
  }

  if (!inputArr.length){
    return [];
  }

  inputArr = inputArr.sort();


  let anagramBuffer = {};
  //ключ - упорядоченное отсортированное слово, хранится само слово по ключу

  inputArr.forEach(word => {


    let sorted = word.toLowerCase().split('').sort().join('');

    if (!anagramBuffer[sorted]) {
      anagramBuffer[sorted] = [];
    }

    anagramBuffer[sorted].push(word);
  });


  let outputArr = [];
  for (const key in anagramBuffer) {
    if (anagramBuffer[key].length > 1) {
      outputArr.push(anagramBuffer[key]);
    }
  }

  return outputArr;
};
