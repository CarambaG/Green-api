// API endpoints
const API_BASE = '/api';

// Display loading state on button
function setButtonLoading(btnSelector, isLoading) {
    const btn = document.querySelector(btnSelector);
    if (btn) {
        btn.disabled = isLoading;
        btn.classList.toggle('loading', isLoading);
        if (isLoading) {
            btn.textContent = 'Loading...';
        }
    }
}

// Get credentials
function getCredentials() {
    const idInstance = document.getElementById('idInstance').value.trim();
    const apiToken = document.getElementById('apiToken').value.trim();

    if (!idInstance || !apiToken) {
        displayError('Please enter both idInstance and ApiTokenInstance');
        return null;
    }

    return { idInstance, apiToken };
}

// Display response
function displayResponse(data) {
    const responseEl = document.getElementById('response');
    try {
        responseEl.textContent = JSON.stringify(data, null, 2);
    } catch (e) {
        responseEl.textContent = String(data);
    }
}

// Display error
function displayError(message) {
    const responseEl = document.getElementById('response');
    responseEl.textContent = JSON.stringify({
        error: message,
        timestamp: new Date().toISOString()
    }, null, 2);
}

// Clear response
function clearResponse() {
    document.getElementById('response').textContent = '{}';
}

// Get Settings
async function getSettings() {
    const creds = getCredentials();
    if (!creds) return;

    setButtonLoading('button', true);
    try {
        const params = new URLSearchParams({
            idInstance: creds.idInstance,
            apiToken: creds.apiToken
        });

        const response = await fetch(`${API_BASE}/settings?${params}`);
        const data = await response.json();

        if (!response.ok) {
            displayError(`Error: ${data.message || response.statusText}`);
        } else {
            displayResponse(data);
        }
    } catch (error) {
        displayError(`Request failed: ${error.message}`);
    } finally {
        setButtonLoading('button', false);
    }
}

// Get State Instance
async function getStateInstance() {
    const creds = getCredentials();
    if (!creds) return;

    setButtonLoading('button', true);
    try {
        const params = new URLSearchParams({
            idInstance: creds.idInstance,
            apiToken: creds.apiToken
        });

        const response = await fetch(`${API_BASE}/state?${params}`);
        const data = await response.json();

        if (!response.ok) {
            displayError(`Error: ${data.message || response.statusText}`);
        } else {
            displayResponse(data);
        }
    } catch (error) {
        displayError(`Request failed: ${error.message}`);
    } finally {
        setButtonLoading('button', false);
    }
}

// Send Message
async function sendMessage() {
    const creds = getCredentials();
    if (!creds) return;

    const chatId = document.getElementById('chatId1').value.trim();
    const message = document.getElementById('message').value.trim();

    if (!chatId || !message) {
        displayError('Please enter Chat ID and Message');
        return;
    }

    setButtonLoading('button', true);
    try {
        const params = new URLSearchParams({
            idInstance: creds.idInstance,
            apiToken: creds.apiToken,
            chatId: chatId,
            message: message
        });

        const response = await fetch(`${API_BASE}/message?${params}`, {
            method: 'POST'
        });
        const data = await response.json();

        if (!response.ok) {
            displayError(`Error: ${data.message || response.statusText}`);
        } else {
            displayResponse(data);
        }
    } catch (error) {
        displayError(`Request failed: ${error.message}`);
    } finally {
        setButtonLoading('button', false);
    }
}

// Send File by URL
async function sendFileByUrl() {
    const creds = getCredentials();
    if (!creds) return;

    const chatId = document.getElementById('chatId2').value.trim();
    const fileUrl = document.getElementById('fileUrl').value.trim();
    const fileName = document.getElementById('fileName').value.trim();
    const caption = document.getElementById('caption').value.trim();

    if (!chatId || !fileUrl) {
        displayError('Please enter Chat ID and File URL');
        return;
    }

    setButtonLoading('button', true);
    try {
        const params = new URLSearchParams({
            idInstance: creds.idInstance,
            apiToken: creds.apiToken,
            chatId: chatId,
            fileUrl: fileUrl,
            fileName: fileName,
            caption: caption
        });

        const response = await fetch(`${API_BASE}/file?${params}`, {
            method: 'POST'
        });
        const data = await response.json();

        if (!response.ok) {
            displayError(`Error: ${data.message || response.statusText}`);
        } else {
            displayResponse(data);
        }
    } catch (error) {
        displayError(`Request failed: ${error.message}`);
    } finally {
        setButtonLoading('button', false);
    }
}

// Store credentials in localStorage for persistence
document.addEventListener('DOMContentLoaded', () => {
    const savedIdInstance = localStorage.getItem('green_api_idInstance');
    const savedApiToken = localStorage.getItem('green_api_apiToken');

    if (savedIdInstance) {
        document.getElementById('idInstance').value = savedIdInstance;
    }
    if (savedApiToken) {
        document.getElementById('apiToken').value = savedApiToken;
    }

    // Save credentials on input
    document.getElementById('idInstance').addEventListener('change', (e) => {
        localStorage.setItem('green_api_idInstance', e.target.value);
    });

    document.getElementById('apiToken').addEventListener('change', (e) => {
        localStorage.setItem('green_api_apiToken', e.target.value);
    });
});