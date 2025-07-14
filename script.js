let elements = [
  {
    "id": 1,
    "name": "Шарик",
    "gender": "M",
    "age": 1,
    "isVaccinated": true,
    "createdAt": "2024-12-12T12:47:40+03:00",
    "updatedAt": null
  },
  {
    "id": 2,
    "name": "Бобик",
    "gender": "TG",
    "age": 228,
    "isVaccinated": false,
    "createdAt": "2024-12-12T12:47:40+03:00",
    "updatedAt": '2222222'
  }
];

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