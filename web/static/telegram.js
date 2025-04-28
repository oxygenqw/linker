function initTelegramWebApp() {
    try {
        const tg = window.Telegram?.WebApp;
        if (!tg) {
            console.warn("Telegram WebApp not detected - running in browser");
            return;
        }

        tg.expand();
        tg.setHeaderColor('#2a2a2a');
        tg.setBackgroundColor('#222222');
        
        console.log("Telegram WebApp initialized successfully");
    } catch (e) {
        console.error("Error initializing Telegram WebApp:", e);
    }
}

document.addEventListener('DOMContentLoaded', initTelegramWebApp);
document.addEventListener('webappready', initTelegramWebApp);