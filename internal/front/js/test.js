console.log("hola hola")

const accessToken = localStorage.getItem('accessToken');

const headers = {
    'Authorization': `Bearer ${accessToken}`
};

fetch('/info-api/info-main-page', {
    method: 'GET',
    headers: headers
})
    .then(response => {
        if (response.status === 200) {
            return response.text();
        } else {
            throw new Error('请求错误');
        }
    })
    .then(html => {
        // 在/api/main上显示HTML
        document.querySelector('.main').innerHTML = html;
    })
    .catch(error => {
        console.error('错误:', error);
    });

