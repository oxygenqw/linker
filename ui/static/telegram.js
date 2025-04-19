// Функция инициализации Telegram WebApp
function initTelegramWebApp() {
    try {
        // Проверяем наличие Telegram WebApp API
        const tg = window.Telegram?.WebApp;
        if (!tg) {
            console.warn("Telegram WebApp not detected - running in browser");
            return;
        }

        // Основные настройки WebApp

        //tg.expand();
        tg.setHeaderColor('#2a2a2a'); // Цвет верхней панели
        tg.setBackgroundColor('#222222'); // Цвет фона
        //tg.disableHeader(); 
        
        console.log("Telegram WebApp initialized successfully");
    } catch (e) {
        console.error("Error initializing Telegram WebApp:", e);
    }
}

// Инициализация при разных событиях для максимальной совместимости
document.addEventListener('DOMContentLoaded', initTelegramWebApp);
document.addEventListener('webappready', initTelegramWebApp);