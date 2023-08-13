// document.querySelector('.butt').onclick = function (event) {
//     event.preventDefault();
//     let inputs = document.querySelectorAll(".inp > input");
//     let data = {};

//     for (i = 0; i < inputs.length; i++) {
//         data[inputs[i].name] = inputs[i].value;
//     }

//     if (data["repeat-password"] != data["password"]) {
//         console.log("Error!");
//     }

//     // let xhr = new XMLHttpRequest();

//     // xhr.open("POST", "/auth/signUp");
//     // xhr.onload = function (e) {
//     //     console.log(e)
//     // };
//     // xhr.send(JSON.stringify(data));

//     // xhr.onreadystatechange = function () {
//     //     if (xhr.readyState === XMLHttpRequest.DONE) {
//     //         if (xhr.status === 200) {
//     //             window.location.replace('/auth/signUp');
//     //         } else if (xhr.status === 400) {
//     //             for (const field in data.errors) {
//     //                 const errorBlock = document.createElement('div');
//     //                 errorBlock.classList.add('err');
//     //                 errorBlock.textContent = data.errors[field];
//     //                 document.querySelector(`input[name="${field}"]`).parentNode.appendChild(errorBlock);
//     //             }
//     //         }
//     //     }
//     // };
//     fetch('/auth/signUp', {
//         method: 'POST',
//         body: JSON.stringify(data),
//         // headers: {
//         //     'Content-Type': 'application/json'
//         // }
//     })
//     .then(response => response.json())
//     .then(data => {
//         if (data.errors) {
//             // Handle validation errors
//             for (const field in data.errors) {
//                 const errorBlock = document.createElement('div');
//                 errorBlock.classList.add('err');
//                 errorBlock.textContent = data.errors[field];
//                 console.log(field)
//                 document.querySelector(`input[name="${field}"]`).parentNode.appendChild(errorBlock);
//             }
//         } else {
//             // Handle successful response and redirect
//             window.location.href = '/auth/signIn';
//         }
//     })
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

    fetch('/auth/signUp', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => response.json())
    .then(data => {
        if (data.errors) {
            // Handle validation errors
            for (const field in data.errors) {
                if (data.errors[field].length !== 0){
                    console.log(field);
                    const errorBlock = document.createElement('div');
                    errorBlock.classList.add('err');
                    errorBlock.textContent = data.errors[field];
                    document.querySelector(`input[name="${field}"]`).insertAdjacentElement('afterend', errorBlock);
                    console.log(field);
                }
            }
        } else {
            // Handle successful response and redirect
            window.location.href = '/auth/signIn';
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
};