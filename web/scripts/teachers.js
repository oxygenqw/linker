import { disableButtons } from './utils.js';

const container = document.querySelector('.container');

const userId = container.dataset.userId;
const userRole = container.dataset.userRole;

const profileButton = document.getElementById('profile');
const backButton = document.getElementById('backButton');
const searchBtn = document.getElementById('searchBtn');
const teacherSearchInput = document.getElementById('teacherSearch');

const buttons = [profileButton, backButton, searchBtn];

profileButton.addEventListener('click', () => {
    disableButtons(buttons);
    if (userRole === 'student') {
        window.location.href = `/users/student/${userId}`;
    } else if (userRole === 'teacher') {
        window.location.href = `/users/teacher/${userId}`;
    }
});

backButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/home/${userId}/${userRole}`;
});

searchBtn.addEventListener('click', () => {
    const query = teacherSearchInput.value.trim();
    disableButtons(buttons);
    window.location.href = `/teachers/${userId}/${userRole}?search=${encodeURIComponent(query)}`;
});

document.querySelectorAll('.details-btn').forEach(btn => {
    btn.addEventListener('click', () => {
        disableButtons(buttons);
        const id = btn.getAttribute('data-id');
        const role = btn.getAttribute('data-role');
        const teacherId = btn.getAttribute('data-teacher-id');
        window.location.href = `/teacher/profile/${id}/${role}/${teacherId}`;
    });
});
