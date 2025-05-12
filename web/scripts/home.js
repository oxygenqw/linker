import { disableButtons } from './utils.js';

const mainContainer = document.querySelector('.main-container');
const profileButton = document.getElementById('profileButton');
const teachersButton = document.getElementById('teachersButton');
const studentsButton = document.getElementById('studentsButton');

const userRole = mainContainer.dataset.userRole;
const userId = mainContainer.dataset.userId;

const buttons = [profileButton, teachersButton, studentsButton];

profileButton.addEventListener('click', () => {
    disableButtons(buttons);
    if (userRole === 'student') {
        window.location.href = `/users/student/${userId}`;
    } else if (userRole === 'teacher') {
        window.location.href = `/users/teacher/${userId}`;
    }
});

teachersButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/teachers/${userId}/${userRole}`;
});

studentsButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/students/${userId}/${userRole}`;
});
