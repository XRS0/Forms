document.getElementById('loginForm').addEventListener('submit', handleLogin);
function handleLogin(event) {
    event.preventDefault();
    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    fetch('/auth/credCheck', { // Измените путь согласно новому эндпоинту
        method: 'POST',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify({username: username, password: password}),
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            document.getElementById('response').textContent = 'Успешный вход!';
        } else {
            document.getElementById('response').textContent = 'Неверные данные.';
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}
