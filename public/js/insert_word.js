'use strict';

document.getElementById('add_synonym').addEventListener('click', (e)=>{
    let input = document.createElement('input');
    input.type = 'text';
    input.autocomplete = 'off';
    input.name = 'synonym[]';
    let synonym_input = document.getElementById('synonym_input');
    synonym_input.appendChild(input);
    synonym_input.appendChild(document.createElement('br'));
});

document.getElementById('add_antonym').addEventListener('click', ()=>{
    let input = document.createElement('input');
    input.type = 'text';
    input.autocomplete = 'off';
    input.name = 'antonym[]';
    let synonym_input = document.getElementById('antonym_input');
    synonym_input.appendChild(input);
    synonym_input.appendChild(document.createElement('br'));
});

document.getElementById('add_examples').addEventListener('click', ()=>{
    let input = document.createElement('input');
    input.type = 'text';
    input.autocomplete = 'off';
    input.name = 'examples[]';
    let synonym_input = document.getElementById('examples_input');
    synonym_input.appendChild(input);
    synonym_input.appendChild(document.createElement('br'));
});

async function postData(url, data = {}) {
    // Default options are marked with *
    const response = await fetch(url, {
      method: "POST", // *GET, POST, PUT, DELETE, etc.
      mode: "cors", // no-cors, *cors, same-origin
      cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
      credentials: "same-origin", // include, *same-origin, omit
      headers: {
        "Content-Type": "application/json",
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      redirect: "follow", // manual, *follow, error
      referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
      body: JSON.stringify(data), // body data type must match "Content-Type" header
    });
    return response.json(); // parses JSON response into native JavaScript objects
  }

document.getElementById('myForm').addEventListener('submit', (e)=>{
    e.preventDefault();
    let word = document.getElementById('word');
    if(/^\s*$/.test(word.value)){
        return;
    };
    let type = document.getElementById('type');
    let meaning = document.getElementById('meaning');
    let synonyms = document.querySelectorAll('#synonym_input > input');
    let antonyms = document.querySelectorAll('#antonym_input > input');
    let examples = document.querySelectorAll('#examples_input > input');
    let getValues = (v)=>{
        let r = [];
        for(let i=0;i<v.length;i++){
            r.push(v[i].value);
        };
        return r;
    };
    let input = {
        word: word.value,
        type: type.value,
        meaning: meaning.value,
        synonym: getValues(synonyms),
        antonym: getValues(antonyms),
        examples: getValues(examples)
    };
    word.value = '';
    type.value = '';
    meaning.value = '';
    let removeElems = (e)=>{
        let elms = document.querySelectorAll(e);
        for(let i=0;i<elms.length;i++){
            elms[i].remove();
        };
    };
    removeElems('#synonym_input > *');
    removeElems('#antonym_input > *');
    removeElems('#examples_input > *');
    let h3 = document.createElement('h3');
    h3.textContent = 'Please waite....';
    document.getElementById('status').appendChild(h3);
    postData("/set_data", input).then((data) => {
        h3.textContent = data.status;
        setTimeout(()=>{
            h3.remove();
        }, 4000);
    });
});