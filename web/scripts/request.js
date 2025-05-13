import { disableButtons, enableButtons } from './utils.js';

const container = document.querySelector('.container');
const senderId = container.dataset.senderId;
const recipientId = container.dataset.recipientId;
const userRole = container.dataset.userRole;

const form = document.getElementById('requestForm');
const errorBlock = document.getElementById('form-error');
const backButton = document.getElementById('backButton');
const submitButton = form.querySelector('button[type="submit"]');

const buttons = [submitButton, backButton];

form.addEventListener('submit', async (e) => {
    e.preventDefault();

    document.querySelectorAll('.input-error').forEach(el => el.classList.remove('input-error'));
    errorBlock.style.display = 'none';
    errorBlock.textContent = '';

    const message = form.message.value.trim();

    if (!message) {
        errorBlock.textContent = 'Пожалуйста, введите сообщение';
        errorBlock.style.display = 'block';
        form.message.focus();
        return;
    }

    disableButtons(buttons);

    try {
        const response = await fetch(`/requests/${senderId}/${recipientId}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ message }),
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.message || 'Ошибка сервера');
        }

        const result = await response.json();
        alert(result.message || 'Запрос отправлен');
        window.history.back();

    } catch (error) {
        errorBlock.textContent = 'Ошибка: ' + error.message;
        errorBlock.style.display = 'block';
        enableButtons(buttons);
        console.error('Ошибка отправки запроса:', error);
    }
});

backButton.addEventListener('click', () => {
    disableButtons(buttons);

    if (userRole === 'student') {
        window.location.href = `/student/profile/${senderId}/${userRole}/${recipientId}`;
    } else if (userRole === 'teacher') {
        window.location.href = `/teacher/profile/${senderId}/${userRole}/${recipientId}`;
    } else {
        window.history.back();
    }
});
