// document.querySelector('.butt').onclick = function (event) {
//   event.preventDefault();
//   let inputs = document.querySelectorAll(".inp > input");
//   let data = {};

//   // Clear error blocks
//   let errorBlocks = document.querySelectorAll('.err');
//   errorBlocks.forEach(block => block.remove());

//   for (i = 0; i < inputs.length; i++) {
//     data[inputs[i].name] = inputs[i].value;
//   }

//   fetch('/auth/signIn', {
//     method: 'POST',
//     body: JSON.stringify(data),
//     headers: {
//       'Content-Type': 'application/json'
//     }
//   })
//     .then(response => response.json())
//     .then(data => {
//       if (data && data.error === "true") {
//         const errorBlock = document.createElement('div');
//         errorBlock.classList.add('err');
//         errorBlock.textContent = "Введите правильный логин или пароль";
//         document.querySelector('.title').insertAdjacentElement('afterend', errorBlock)

//       } else if (data && data.accessToken.length !== 0) {
//         localStorage.setItem('accessTokenken', data.accessToken);
//         const accessToken = localStorage.getItem('accessTokenken');

//         window.location.replace('/api/main');

//         const url = '/info-api/info-main-page';
//         const headers = {
//           'Authorization': `Bearer ${accessToken}`
//         };

//         fetch(url, {
//           method: 'GET',
//           headers: headers
//         })
//           .then(response => {
//             if (response.status === 200) {
//               return response.text();
//             } else if (response.status === 404) {
//               throw new Error('Страница не найдена');
//             } else {
//               throw new Error('Ошибка сервера');
//             }
//           })
//           .then(html => {
//             // Отобразить полученный HTML на странице
//             document.body.innerHTML = html;
//           })
//           .catch(error => {
//             console.error('Error:', error);
//           });
//       }
//     })
//     .catch(error => {
//       console.error('Error:', error);
//     });
// };

// document.querySelector('.butt').onclick = function (event) {
//   event.preventDefault();
//   let inputs = document.querySelectorAll(".inp > input");
//   let data = {};

//   let errorBlocks = document.querySelectorAll('.err');
//   errorBlocks.forEach(block => block.remove());

//   for (i = 0; i < inputs.length; i++) {
//     data[inputs[i].name] = inputs[i].value;
//   }

//   const xhr = new XMLHttpRequest();
//   xhr.open('POST', '/auth/signIn');
//   xhr.onload = function () {
//     if (data && data.error === "true") {
//       const errorBlock = document.createElement('div');
//       errorBlock.classList.add('err');
//       errorBlock.textContent = "Введите правильный логин или пароль";
//       document.querySelector('.title').insertAdjacentElement('afterend', errorBlock)

//     } else if (data && data.accessToken.length !== 0) {
//       localStorage.setItem('accessToken', data.accessToken);
//       const url = '/api/main';
//       window.location.replace(url);
//     }
//   };

// };

document.querySelector('.butt').onclick = function (event) {
  event.preventDefault();
  let inputs = document.querySelectorAll(".inp > input");
  let data = {};

  // Clear error blocks
  let errorBlocks = document.querySelectorAll('.err');
  errorBlocks.forEach(block => block.remove());

  for (i = 0; i < inputs.length; i++) {
    data[inputs[i].name] = inputs[i].value;
  }

  fetch('/auth/signIn', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(response => response.json())
    .then(data => {
      if (data && data.error === "true") {
        const errorBlock = document.createElement('div');
        errorBlock.classList.add('err');
        errorBlock.textContent = "Введите правильный логин или пароль";
        document.querySelector('.title').insertAdjacentElement('afterend', errorBlock)

      } else if (data && data.accessToken.length !== 0) {
        localStorage.setItem('accessTokenken', data.accessToken);
        const url = '/api/main';
        window.location.replace(url);
      }
    }
    )
}

