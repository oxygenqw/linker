import { disableButtons, enableButtons } from './utils.js';

const container = document.querySelector('.main-container');
const telegramId = Number(container.dataset.telegramId);
const userName = container.dataset.userName;

const form = document.getElementById('registerForm');
const errorBlock = document.getElementById('form-error');
const submitButton = form.querySelector('button[type="submit"]');

form.addEventListener('submit', async (e) => {
    e.preventDefault();

    document.querySelectorAll('.input-error').forEach(el => el.classList.remove('input-error'));
    errorBlock.style.display = 'none';
    errorBlock.textContent = '';

    const firstName = form.first_name.value.trim();
    const lastName = form.last_name.value.trim();
    const middleName = form.middle_name.value.trim();
    const role = form.role.value;

    let errorMsg = '';
    let errorField = null;
    const nameRegex = /^[a-zA-Zа-яА-ЯёЁ\- ]+$/;

    if (!firstName) {
        errorMsg = 'Пожалуйста, введите имя';
        errorField = form.first_name;
    } else if (firstName.length > 50) {
        errorMsg = 'Имя не должно превышать 50 символов';
        errorField = form.first_name;
    } else if (!nameRegex.test(firstName)) {
        errorMsg = 'Имя может содержать только буквы';
        errorField = form.first_name;
    } else if (!lastName) {
        errorMsg = 'Пожалуйста, введите фамилию';
        errorField = form.last_name;
    } else if (lastName.length > 50) {
        errorMsg = 'Фамилия не должна превышать 50 символов';
        errorField = form.last_name;
    } else if (!nameRegex.test(lastName)) {
        errorMsg = 'Фамилия может содержать только буквы';
        errorField = form.last_name;
    } else if (!middleName) {
        errorMsg = 'Пожалуйста, введите отчество';
        errorField = form.middle_name;
    } else if (middleName.length > 50) {
        errorMsg = 'Отчество не должно превышать 50 символов';
        errorField = form.middle_name;
    } else if (!nameRegex.test(middleName)) {
        errorMsg = 'Отчество может содержать только буквы';
        errorField = form.middle_name;
    } else if (!role) {
        errorMsg = 'Пожалуйста, выберите роль';
        errorField = form.role;
    }

    if (errorMsg) {
        if (errorField) errorField.classList.add('input-error');
        errorBlock.textContent = errorMsg;
        errorBlock.style.display = 'block';
        if (errorField) errorField.focus();
        return;
    }

    let endpoint = '';
    if (role === 'student') {
        endpoint = `/users/student`;
    } else if (role === 'teacher') {
        endpoint = `/users/teacher`;
    }

    const data = {
        telegram_id: telegramId,
        user_name: userName,
        first_name: firstName,
        last_name: lastName,
        middle_name: middleName
    };

    disableButtons([submitButton]);

    try {
        const response = await fetch(endpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) {
            const error = await response.text();
            throw new Error(error || 'Ошибка сервера');
        }

        const result = await response.json();
        if (role === 'student') {
            window.location.href = `/home/${result.id}/student`;
        } else if (role === 'teacher') {
            window.location.href = `/home/${result.id}/teacher`;
        }
    } catch (error) {
        errorBlock.textContent = 'Ошибка: ' + error.message;
        errorBlock.style.display = 'block';
        enableButtons([submitButton]);
    }
});
