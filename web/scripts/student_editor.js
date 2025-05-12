import { disableButtons, enableButtons } from './utils.js';

const form = document.getElementById('profileForm');
const saveButton = document.getElementById('saveButton');
const backButton = document.getElementById('backButton');

const buttons = [saveButton, backButton];

const container = document.querySelector('.container');
const studentId = container.dataset.studentId;

backButton.addEventListener('click', () => {
    disableButtons(buttons);
    window.location.href = `/users/student/${studentId}`;
});

saveButton.addEventListener('click', async () => {
    document.querySelectorAll('.input-error').forEach(el => el.classList.remove('input-error'));
    const errorBlock = document.getElementById('form-error');
    errorBlock.style.display = 'none';
    errorBlock.textContent = '';

    const firstName = form.first_name.value.trim();
    const lastName = form.last_name.value.trim();
    const middleName = form.middle_name.value.trim();
    const university = form.university.value.trim();
    const faculty = form.faculty.value.trim();
    const specialty = form.specialty.value.trim();
    const github = form.github.value.trim();
    const job = form.job.value.trim();

    const nameRegex = /^[a-zA-Zа-яА-ЯёЁ\- ]+$/;
    let errorMsg = '';
    let errorField = null;

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
    } else if (middleName && (middleName.length > 50 || !nameRegex.test(middleName))) {
        errorMsg = middleName.length > 50
            ? 'Отчество не должно превышать 50 символов'
            : 'Отчество может содержать только буквы';
        errorField = form.middle_name;
    } else if (university.length > 100) {
        errorMsg = 'Название университета не должно превышать 100 символов';
        errorField = form.university;
    } else if (faculty.length > 100) {
        errorMsg = 'Название факультета не должно превышать 100 символов';
        errorField = form.faculty;
    } else if (github.length > 100) {
        errorMsg = 'GitHub не должен превышать 100 символов';
        errorField = form.github;
    } else if (job.length > 100) {
        errorMsg = 'Место работы не должно превышать 100 символов';
        errorField = form.job;
    }

    if (errorMsg) {
        if (errorField) errorField.classList.add('input-error');
        errorBlock.textContent = errorMsg;
        errorBlock.style.display = 'block';
        if (errorField) errorField.focus();
        return;
    }

    disableButtons(buttons);

    try {
        const response = await fetch(`/users/student`, {
            method: 'PATCH',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                id: studentId,
                first_name: firstName,
                last_name: lastName,
                middle_name: middleName,
                university: university,
                faculty: faculty,
                specialty: specialty,
                course: form.course.value,
                education: form.education.value,
                github: github,
                job: job,
                idea: form.idea.value,
                about: form.about.value
            })
        });

        if (response.ok) {
            window.location.href = `/users/student/${studentId}`;
        } else {
            const error = await response.text();
            throw new Error(error || 'Ошибка сервера');
        }
    } catch (error) {
        errorBlock.textContent = 'Ошибка: ' + error.message;
        errorBlock.style.display = 'block';
        enableButtons(buttons);
    }
});
