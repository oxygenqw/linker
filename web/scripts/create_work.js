import { disableButtons, enableButtons } from './utils.js';

const container = document.querySelector('.work-container');
const userId = container.dataset.userId;
const role = container.dataset.role;

const form = document.getElementById('workForm');
const errorBlock = document.getElementById('form-error');
const backButton = document.getElementById('backButton');
const submitButton = form.querySelector('button[type="submit"]') || document.querySelector('button[form="workForm"]');

const buttons = [submitButton, backButton];

form.addEventListener('submit', async (e) => {
    e.preventDefault();

    document.querySelectorAll('.input-error').forEach(el => el.classList.remove('input-error'));
    errorBlock.style.display = 'none';
    errorBlock.textContent = '';

    const link = form.link.value.trim();
    const description = form.description.value.trim();

    if (!link) {
        errorBlock.textContent = 'Пожалуйста, введите ссылку на работу';
        errorBlock.style.display = 'block';
        form.link.focus();
        return;
    }

    if (!description) {
        errorBlock.textContent = 'Пожалуйста, введите описание работы';
        errorBlock.style.display = 'block';
        form.description.focus();
        return;
    }

    disableButtons(buttons);

    try {
        const response = await fetch(`/works/${userId}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ link, description }),
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(errorText || 'Ошибка сервера');
        }

        if (role === 'student') {
            window.location.href = `/users/student/${userId}`;
        } else if (role === 'teacher') {
            window.location.href = `/users/teacher/${userId}`;
        }
    } catch (error) {
        errorBlock.textContent = 'Ошибка: ' + error.message;
        errorBlock.style.display = 'block';

        enableButtons(buttons);
    }
});

backButton.addEventListener('click', () => {
    disableButtons(buttons);

    if (role === 'student') {
        window.location.href = `/users/student/${userId}`;
    } else if (role === 'teacher') {
        window.location.href = `/users/teacher/${userId}`;
    }
});
