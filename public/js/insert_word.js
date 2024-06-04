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
