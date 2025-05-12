export function disableButtons(buttons) {
    buttons.forEach(button => {
        if ('disabled' in button) {
            button.disabled = true;
        }
    });
}
