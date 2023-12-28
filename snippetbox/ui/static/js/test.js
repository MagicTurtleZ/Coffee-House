const countGoods = { 
    "expresso": 0,
    "americano": 0,
    "latte": 0,
    "capuchino": 0,
    "chocolate_muffin": 0,
    "blueberry_muffin": 0,
    "apple_tart": 0
}

const choicePriceGoods = { 
    "expresso": 80,
    "americano": 90,
    "latte": 110,
    "capuchino": 120,
    "chocolate_muffin": 80,
    "blueberry_muffin": 90,
    "apple_tart": 100
}

var sum = 0;
const goodsElements = document.querySelectorAll('[name="goods"]');
const countElements = document.querySelectorAll(".item__cout");
const resultElem = document.querySelector(".sum");
const btnElem = document.querySelector(".btn");
const clientElem = document.querySelectorAll(".name_surname");


goodsElements.forEach(product => {
    product.addEventListener("change", function() {
        if (product.checked) {
            countGoods[product.dataset.goods] = 1;
            showValue();
            calcSum();
        } else {
            countGoods[product.dataset.goods] = 0;
            showValue();
            calcSum();
        }
    });
});

countElements.forEach(elem => {
    elem.addEventListener("change", function() {
        countGoods[elem.id] = parseInt(elem.value);
        showValue();
        calcSum();
    });
});

btnElem.addEventListener("click", function() {
    alert(`Заказчик: ${clientElem[0].value} ${clientElem[1].value}\nИтого: ${sum}`);
});

function showValue() {
    for (var i = 0; i < countElements.length; i++) {
        countElements[i].value = countGoods[countElements[i].id];
    }
    console.log(countGoods);
}

function calcSum() {
    sum = 0;
    for (var i = 0; i < countElements.length; i++) {
        sum += countGoods[countElements[i].id] * choicePriceGoods[countElements[i].id];
    }
    console.log(sum);
    resultElem.textContent = ` ${sum} р`
}