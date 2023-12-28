btnElement = document.querySelector(".footer__sub");
helloElement = document.querySelector(".hello__world");
emailResult = document.querySelector('[name="coffee__mail"]');

btnElement.addEventListener("click", function(){
    if (emailResult.value != "") {
        helloElement.textContent = `Проверочное письмо было отправлено на почту: ${emailResult.value}`;
    } else helloElement.textContent = "Пожалуйста, укажите почту и повторите попытку!"
});