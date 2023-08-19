document.querySelector('.butt').onclick = function (event) {
  event.preventDefault();
  let inputs = document.querySelectorAll(".inp > input");
  let data = {};
  for (i = 0; i < inputs.length; i++) {
    data[inputs[i].name] = inputs[i].value;
  }

  let xhr = new XMLHttpRequest();
  xhr.open("POST", "/auth/signIn");
  xhr.onload = function (e) {
    console.log(e)
  };

  xhr.send(JSON.stringify(data));

  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE) {
      if (xhr.status === 200) {
        window.location.replace('/api/main');
      } else {
        console.log('Ошибка: ' + xhr.status);
      }
    }
  };
};