let xhr = new XMLHttpRequest();
xhr.open("GET", "http://localhost:8080/api/v1/dogs");
xhr.send();
xhr.onerror = function() {
    alert('Запрос собак не удался: не удалось отправить запрос');
};

xhr.onload = function() {
    if (xhr.status === 200) {
        let elements = JSON.parse(xhr.responseText);
        console.log(elements);
        
        // отрисовка строк
        let list = document.querySelector('table');

        for (let i = 0; i < elements.length; i++) {
            let newRow = document.createElement('tr');

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

