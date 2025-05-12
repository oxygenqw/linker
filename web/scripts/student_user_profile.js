import { disableButtons } from './utils.js';

const container = document.querySelector('.container');

const userId = container.dataset.userId;
const userRole = container.dataset.userRole;
const recipientID = container.dataset.recipientId;

const toRequestFormButton = document.getElementById('toRequestFormButton');
const homeButton = document.getElementById('homeButton');
const backButton = document.getElementById('backButton');

const buttons = [toRequestFormButton, homeButton, backButton];

toRequestFormButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/requests/${userId}/${recipientID}`;
});

homeButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/home/${userId}/${userRole}`;
});

backButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/students/${userId}/${userRole}`;
});
