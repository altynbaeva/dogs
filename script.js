// отрисовка листинга
function drawListing() {
    let xhrListDogs = new XMLHttpRequest();
    xhrListDogs.open("GET", "http://localhost:8080/api/v1/dogs");
    xhrListDogs.send();
    xhrListDogs.onerror = function() {
        alert('Запрос собак не удался: не удалось отправить запрос');
    };
    xhrListDogs.onload = function() {
        if (xhrListDogs.status === 200) {
            let elements = JSON.parse(xhrListDogs.responseText);
            
            // отрисовка строк
            let list = document.querySelector('table');

            for (let i = 0; i < elements.length; i++) {
                let newRow = document.createElement('tr');
                newRow.classList.add('dogs-list-body-row');

                for (let j = 0; j < 8; j++) {
                    let newCell = document.createElement('td');

                    if (j === 0) {
                        newCell.textContent = elements[i].id
                    }

                    if (j === 1) {
                        newCell.textContent = elements[i].name
                    }

                    if (j === 2) {
                        newCell.textContent = elements[i].gender
                    }

                    if (j === 3) {
                        newCell.textContent = elements[i].age
                    }

                    if (j === 4) {
                        if (elements[i].isVaccinated === true) {
                            newCell.textContent = 'Привит'
                        }   else {
                            newCell.textContent = 'Не привит'
                        }
                    }

                    if (j === 5) {
                        newCell.textContent = elements[i].createdAt
                    }

                    if (j === 6) {
                        if (elements[i].updatedAt === null) {
                            newCell.textContent = 'Не заполнено'
                        } else {
                            newCell.textContent = elements[i].updatedAt
                        }
                    }

                    if (j === 7) {
                        let newButton = document.createElement('button');
                        newButton.classList.add('more-button');
                        newButton.onclick = function() {
                            let xhrDogById = new XMLHttpRequest();
                            xhrDogById.open('GET', 'http://localhost:8080/api/v1/dogs/'+ elements[i].id);
                            xhrDogById.send();
                            xhrDogById.onerror = function() {
                                alert('Запрос собаки не удался: не удалось отправить запрос');
                            }

                            xhrDogById.onload = function() {
                                if (xhrDogById.status === 200) {
                                    // парсим ответ
                                    let dogCard = JSON.parse(xhrDogById.responseText);
                                
                                    // проставляем значения в карточку
                                    // TODO писать не в ячейку, а в вложенный элемент p?
                                    document.getElementById("dog-card-photo").src = dogCard.photoURL;
                                    document.getElementById("id-cell").textContent = dogCard.id;
                                    document.getElementById("name-cell").textContent = dogCard.name;
                                    document.getElementById("gender-cell").textContent = dogCard.gender;
                                    document.getElementById("age-cell").textContent = dogCard.age;
                                    if (dogCard.isVaccinated === true) {
                                        document.getElementById("is-vaccinated-cell").textContent = 'Привит'
                                    } else {
                                        document.getElementById("is-vaccinated-cell").textContent = 'Не привит'
                                    }
                                
                                    document.getElementById("created-at-cell").textContent = dogCard.createdAt;
                                    if (dogCard.updatedAt === null) {
                                        document.getElementById("updated-at-cell").textContent = 'Не заполнено'
                                    } else {
                                        document.getElementById("updated-at-cell").textContent = dogCard.updatedAt
                                    }
                                    // TODO в столбик
                                    document.getElementById("favorite-food-cell").textContent = dogCard.personality.favoriteFood.join(','); 

                                    // отображаем карточку
                                    document.getElementById("dogs-listing").hidden=true;
                                    document.getElementById("dog-by-id").hidden=false;
                                } else {
                                    alert("Запрос собаки не удался: http code != 200"); 
                                }
                            }
                        }
                        newCell.appendChild(newButton);
                        let newPar = document.createElement('p');
                        newPar.textContent = 'Подробнее...'
                        newButton.appendChild(newPar);

                    }

                    newRow.appendChild(newCell);
                }

                list.appendChild(newRow);
            }
        } else {
            alert("Запрос собак не удался: http code != 200");
        }
    }
};

drawListing();

let returnButton = document.getElementById("dog-card-return-button");
returnButton.onclick = function() {
    document.getElementById("dog-by-id").hidden=true;
    document.getElementById("dogs-listing").hidden=false;
};
 
let deleteButton = document.getElementById("dog-card-delete-button");
deleteButton.onclick = function() {
    let xhrDeleteDog = new XMLHttpRequest();
    xhrDeleteDog.open('DELETE', 'http://localhost:8080/api/v1/dogs/' + document.getElementById("id-cell").textContent);
    xhrDeleteDog.send();
    xhrDeleteDog.onerror = function() {
        alert('Удаление не удалось: не удалось отправить запрос');
    }

    xhrDeleteDog.onload = function() {
        if (xhrDeleteDog.status === 200) {
            let elementsToRemove = document.querySelectorAll('.dogs-list-body-row');
            for (let i = 0; i < elementsToRemove.length; i++) {
                elementsToRemove[i].remove();
            }
            
            drawListing();

            document.getElementById('dog-by-id').hidden=true;
            document.getElementById('dogs-listing').hidden=false;
            

        } else {
            alert("Удаление не удалось: http code != 200");
        }
    } 
};

let createButton = document.getElementById("create-dog-button");
createButton.onclick = function() {
    document.getElementById('dogs-listing').hidden=true;
    document.getElementById('send-new-dog-form').hidden=false;
};

let sendButton = document.getElementById("send-button");
sendButton.onclick = function() {
    const form = document.getElementById("create-dog-form");
    form.addEventListener('submit', function(event) {
    event.preventDefault();
    const formData = new FormData(form);
    });
    
    let xhrSendDog = new XMLHttpRequest();
    xhrSendDog.open('POST', 'http://localhost:8080/api/v1/dogs');
    xhrSendDog.send()
};