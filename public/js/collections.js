'use strict';

async function getCollections() {
    const response = await fetch('/words');
    const collections = response.json();
    return collections;
}

function createCollectionElement(element) {
    let root = document.createElement('div');
    //root.id = `id_${element._id}`;
    let div1 = document.createElement('div');
    div1.innerHTML = `${element.word} (${element.type}) - ${element.meaning}`;
    root.appendChild(div1);
    if (element.synonym){
        let synonym = document.createElement('div');
        synonym.innerHTML = 'Synonym: '+element.synonym.toString();
        root.appendChild(synonym);
    };
    if (element.antonym){
        let antonym = document.createElement('div');
        antonym.innerHTML = 'Antonym: '+element.antonym.toString();
        root.appendChild(antonym);
    }
    if(element.examples){
        let examples = document.createElement('ol');
        for(let i=0;i<element.examples.length;i++){
           let li = document.createElement('li');
           li.innerHTML = element.examples[i];
           examples.appendChild(li);
        }
        root.appendChild(examples);
    }
    let delete_btn = document.createElement('button');
    delete_btn.textContent = 'Delete';
    delete_btn.addEventListener('click', async ()=>{
        let cnf = confirm(`Do you want to delete -\n\nWord: ${element.word}\nType: ${element.type}`);
        if (!cnf) {
            return;
        };
        let response = await fetch(`/words/id/${element._id}`, {method: 'DELETE'});
        let result = {}
        try {
            result = await response.json();
        } catch(err){
            alert('Failed to remove');
            return;
        }
        if ((result._id==element._id)&&(result.status=='Deleted')) {
            root.remove();
            return;
        };
        alert('Failed to remove');
    })
    root.appendChild(delete_btn);
    root.appendChild(document.createElement('hr'));
    return root;
}

!async function(){
    let collections = await getCollections();
    let colls = document.getElementById('colls');
    for(let i=0; i<collections.length;i++){
        colls.appendChild(createCollectionElement(collections[i]));
    }
}();
