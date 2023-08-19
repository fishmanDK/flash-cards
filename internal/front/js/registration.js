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
            if (data && data.errors) {
                // Handle validation errors
                for (const field in data.errors) {
                    if (data.errors[field].length !== 0) {
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
                console.log("hola");
                window.location.replace('/auth/signIn');
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
};