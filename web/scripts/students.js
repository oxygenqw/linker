import { disableButtons } from './utils.js';

const container = document.querySelector('.container');

const userId = container.dataset.userId;
const userRole = container.dataset.userRole;

const profileButton = document.getElementById('profile');
const backButton = document.getElementById('backButton');
const searchBtn = document.getElementById('searchBtn');
const studentSearchInput = document.getElementById('studentSearch');

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
    const query = studentSearchInput.value.trim();
    disableButtons(buttons);
    window.location.href = `/students/${userId}/${userRole}?search=${encodeURIComponent(query)}`;
});

document.querySelectorAll('.details-btn').forEach(btn => {
    btn.addEventListener('click', () => {
        disableButtons(buttons);
        const id = btn.getAttribute('data-id');
        const role = btn.getAttribute('data-role');
        const studentId = btn.getAttribute('data-student-id');
        window.location.href = `/student/profile/${id}/${role}/${studentId}`;
    });
});
