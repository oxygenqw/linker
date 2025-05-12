import { disableButtons, enableButtons } from './utils.js';

const container = document.querySelector('.container');

const teacherId = container.dataset.teacherId;
const userName = container.dataset.username;
const telegramId = container.dataset.telegramId;

const editProfileButton = document.getElementById('editProfileButton');
const deleteProfileButton = document.getElementById('deleteProfileButton');
const addWorkButton = document.getElementById('addWorkButton');
const backButton = document.getElementById('backButton');

const buttons = [editProfileButton, deleteProfileButton, addWorkButton, backButton];

editProfileButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/teacher/edit/${teacherId}`;
});

addWorkButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/works/${teacherId}`;
});

backButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/home/${teacherId}/teacher`;
});

deleteProfileButton.addEventListener('click', async () => {
    disableButtons(buttons);
    try {
        const response = await fetch(`/users/teacher/${teacherId}`, { method: 'DELETE' });
        if (response.ok) {
            window.location.href = `/login/${userName}/${telegramId}`;
        } else {
            const text = await response.text();
            throw new Error(text || `Server error: ${response.status}`);
        }
    } catch (error) {
        alert('Не удалось удалить профиль: ' + error.message);
        enableButtons(buttons);
    }
});

document.addEventListener('click', async (e) => {
    if (e.target.classList.contains('work-delete-button')) {
        disableButtons(buttons);
        const workId = e.target.dataset.workId;
        try {
            const response = await fetch(`/works/${workId}`, { method: 'DELETE' });
            if (response.ok) {
                window.location.reload();
            } else {
                const error = await response.text();
                alert('Ошибка удаления: ' + (error || 'Ошибка сервера'));
                enableButtons(buttons);
            }
        } catch (error) {
            alert('Ошибка удаления: ' + error.message);
            enableButtons(buttons);
        }
    } else if (e.target.classList.contains('work-edit-button')) {
        alert('Редактирование пока недоступно');
    }
});