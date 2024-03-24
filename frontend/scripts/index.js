document.getElementById('loginForm').addEventListener('submit', handleLogin);

function handleLogin(event) {
    event.preventDefault(); // Предотвратить стандартное поведение формы
    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    // Сначала проверяем учетные данные пользователя
    fetch('/auth/credCheck', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({username: username, password: password}),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Ошибка при проверке учетных данных');
        }
        return response.json();
    })
    .then(data => {
        if (data.success) {
            // Учетные данные верны, теперь запрашиваем JWT токен
            return getJwtToken(username, password);
        } else {
            // Если учетные данные неверны, сообщаем об этом
            document.getElementById('response').textContent = 'Неверные данные.';
            throw new Error('Неверные учетные данные');
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

// function getJwtToken(username, password) {
//     return fetch('/auth/jwt', {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({username: username, password: password}),
//     })
//     .then(response => {
//         if (!response.ok) {
//             throw new Error('Ошибка аутентификации');
//         }
//         return response.json();
//     })
//     .then(data => {
//         console.log('Токен успешно получен:', data.jwt);
//         // Сохраняем токен для последующего использования
//         localStorage.setItem('jwtToken', data.jwt);
//         // Перенаправляем пользователя на защищенную страницу или обновляем UI
//         document.getElementById('response').textContent = 'Успешный вход!';
//     })
//     .catch(error => {
//         console.error('Произошла ошибка при получении токена:', error);
//         document.getElementById('response').textContent = 'Произошла ошибка при получении токена.';
//     });
// }


// const token = 'ваш_токен_здесь'; // Токен, который вы получили после аутентификации

// fetch('https://example.com/api/protected', {
//     method: 'GET', // или POST, PUT, DELETE и т.д.
//     headers: {
//         'Authorization': `Bearer ${token}`
//     }
// })
// .then(response => response.json())
// .then(data => console.log(data))
// .catch(error => console.error('Ошибка:', error));
