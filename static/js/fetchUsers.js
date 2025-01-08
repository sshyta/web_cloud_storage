async function fetchUsers() {
    try {
        const response = await fetch('/user/list'); // API для получения списка пользователей
        if (response.ok) {
            const users = await response.json(); // Получение JSON-данных
            console.log('Fetched Users:', users);

            // Очистка текущего списка пользователей
            const userList = document.querySelector('.user-list'); // Найдите контейнер для списка пользователей
            userList.innerHTML = "";

            // Динамическое добавление пользователей в список
            users.forEach(user => {
                const userDiv = document.createElement('div');
                userDiv.className = 'user-item';
                userDiv.innerHTML = `
          <p><strong>Username:</strong> ${user.Username}</p>
          <p><strong>Login:</strong> ${user.Login}</p>
          <p><strong>Email:</strong> ${user.WorkingEmail}</p>
          <p><strong>Role ID:</strong> ${user.RolesID}</p>
        `;
                userList.appendChild(userDiv);
            });
        } else {
            console.error('Failed to fetch users:', response.status);
        }
    } catch (error) {
        console.error('Error fetching users:', error);
    }
}


// Вызов функции fetchUsers при загрузке страницы
window.onload = fetchUsers;