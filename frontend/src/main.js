import './style.css';
import bgImage from './assets/images/bg.gif';
import logoImage from './assets/images/logo-universal.png';

const translations = {
    en: {
        "Initializing...": "Initializing...",
        "Waiting for launcher to close...": "Waiting for launcher to close...",
        "Renaming old executable...": "Renaming old executable...",
        "Downloading update...": "Downloading update...",
        "Update finished!": "Update finished!",
        "No update required.": "No update required.",
        "Open Amatayakul": "Open Amatayakul",
        "Cancel Update?": "Cancel Update?",
        "Are you sure you want to cancel the update process?": "Are you sure you want to cancel the update process?",
        "No, resume": "No, resume",
        "Yes, cancel": "Yes, cancel"
    },
    es: {
        "Initializing...": "Iniciando...",
        "Waiting for launcher to close...": "Esperando a que se cierre el launcher...",
        "Renaming old executable...": "Renombrando ejecutable antiguo...",
        "Downloading update...": "Descargando actualización...",
        "Update finished!": "¡Actualización finalizada!",
        "No update required.": "No se requiere actualización.",
        "Open Amatayakul": "Abrir Amatayakul",
        "Cancel Update?": "¿Cancelar actualización?",
        "Are you sure you want to cancel the update process?": "¿Estás seguro de que quieres cancelar el proceso de actualización?",
        "No, resume": "No, continuar",
        "Yes, cancel": "Sí, cancelar"
    },
    pt: {
        "Initializing...": "Inicializando...",
        "Waiting for launcher to close...": "Aguardando o launcher fechar...",
        "Renaming old executable...": "Renomeando executável antigo...",
        "Downloading update...": "Baixando atualização...",
        "Update finished!": "Atualização finalizada!",
        "No update required.": "Nenhuma atualização necessária.",
        "Open Amatayakul": "Abrir Amatayakul",
        "Cancel Update?": "Cancelar Atualização?",
        "Are you sure you want to cancel the update process?": "Tem certeza que deseja cancelar o processo de atualização?",
        "No, resume": "Não, continuar",
        "Yes, cancel": "Sim, cancelar"
    }
};

let currentLang = 'en';

function t(key) {
    if (translations[currentLang] && translations[currentLang][key]) {
        return translations[currentLang][key];
    }
    return key;
}

document.querySelector('#app').innerHTML = `
    <div class="bg-wrapper">
        <img src="${bgImage}" alt="" class="bg-image">
        <div class="bg-dim"></div>
        <div class="noise"></div>
    </div>

    <!-- Window Bar -->
    <header class="app-header" style="--wails-draggable:drag">
        <div class="header-brand"></div>
        <div class="header-controls" style="--wails-draggable:no-drag">
            <div class="window-controls">
                <button class="ctrl-btn" id="btnMinimize">
                    <svg width="12" height="12" viewBox="0 0 24 24"><path d="M19 13H5v-2h14v2z" fill="currentColor"/></svg>
                </button>
                <button class="ctrl-btn ctrl-close" id="btnClose">
                    <svg width="12" height="12" viewBox="0 0 24 24"><path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z" fill="currentColor"/></svg>
                </button>
            </div>
        </div>
    </header>

    <div class="updater-container">
        <div class="card-logo-container">
            <img src="${logoImage}" alt="Amatayakul" class="main-logo" />
        </div>
        
        <div id="content">
            <div id="statusMessage" class="status-msg">Initializing...</div>
            
            <div class="progress-bar">
                <div id="progressFill" class="progress-fill" style="width: 0%"></div>
            </div>
            
            <div class="stats-row" id="statsRow" style="display: none;">
                <span id="speedTxt">0 MB/s</span>
                <span id="sizeTxt">0 / 0 MB</span>
                <span id="timeTxt">Estimating...</span>
            </div>
<<<<<<< HEAD
            
            <button id="btnDone" class="btn" style="display: none;">Open Amatayakul</button>
=======
>>>>>>> 1e7144e (1.0.0)
        </div>
    </div>

    <!-- Confirmation Modal -->
    <div id="closeModal" class="modal-overlay" style="display: none;">
        <div class="modal-box">
            <h3 id="modalTitle">Cancel Update?</h3>
            <p id="modalDesc">Are you sure you want to cancel the update process?</p>
            <div class="modal-actions">
                <button id="btnModalNo" class="btn btn-secondary">No, resume</button>
                <button id="btnModalYes" class="btn btn-danger">Yes, cancel</button>
            </div>
        </div>
    </div>
`;

let statusMessage = document.getElementById('statusMessage');
let progressFill = document.getElementById('progressFill');
let statsRow = document.getElementById('statsRow');
let speedTxt = document.getElementById('speedTxt');
let sizeTxt = document.getElementById('sizeTxt');
let timeTxt = document.getElementById('timeTxt');
<<<<<<< HEAD
let btnDone = document.getElementById('btnDone');
=======
>>>>>>> 1e7144e (1.0.0)

let btnMinimize = document.getElementById('btnMinimize');
let btnClose = document.getElementById('btnClose');
let closeModal = document.getElementById('closeModal');
let btnModalYes = document.getElementById('btnModalYes');
let btnModalNo = document.getElementById('btnModalNo');
let modalTitle = document.getElementById('modalTitle');
let modalDesc = document.getElementById('modalDesc');

let finalPath = "";

btnMinimize.addEventListener('click', () => {
    window.runtime.WindowMinimise();
});

btnClose.addEventListener('click', () => {
    closeModal.style.display = 'flex';
});

btnModalNo.addEventListener('click', () => {
    closeModal.style.display = 'none';
});

btnModalYes.addEventListener('click', () => {
    window.go.main.App.Exit();
});

function applyTranslations() {
    statusMessage.textContent = t("Initializing...");
<<<<<<< HEAD
    btnDone.textContent = t("Open Amatayakul");
=======
>>>>>>> 1e7144e (1.0.0)
    modalTitle.textContent = t("Cancel Update?");
    modalDesc.textContent = t("Are you sure you want to cancel the update process?");
    btnModalNo.textContent = t("No, resume");
    btnModalYes.textContent = t("Yes, cancel");
}

async function initUpdater() {
    try {
        const args = await window.go.main.App.GetArgs();
        
        if (args.lang && translations[args.lang]) {
            currentLang = args.lang;
        }
        applyTranslations();
        
        if (!args.isUpdate) {
            statusMessage.textContent = t("No update required.");
            return;
        }

        window.runtime.EventsOn("update:status", (msg) => {
            statusMessage.textContent = t(msg);
        });

        window.runtime.EventsOn("update:progress", (data) => {
            statsRow.style.display = 'flex';
            progressFill.style.width = `${data.percent}%`;
            
            speedTxt.textContent = `${data.mbps.toFixed(2)} Mbps`;
            sizeTxt.textContent = `${data.downloadedMB.toFixed(2)} / ${data.totalMB.toFixed(2)} MB`;
            timeTxt.textContent = `${Math.ceil(data.remainingSec)}s`;
        });

        window.runtime.EventsOn("update:error", (msg) => {
            statusMessage.textContent = msg; // Dynamic error, don't translate fully
            if (msg.startsWith("Failed to")) {
               statusMessage.textContent = "Error: " + msg;
            }
            statusMessage.style.color = "#ff4444";
            statsRow.style.display = 'none';
        });

        window.runtime.EventsOn("update:done", (path) => {
<<<<<<< HEAD
            statusMessage.textContent = t("Update finished!");
            statusMessage.style.color = "#44ff44";
            statsRow.style.display = 'none';
            btnDone.style.display = 'block';
            finalPath = path;
        });

        btnDone.addEventListener('click', () => {
            window.go.main.App.LaunchAndExit(finalPath);
=======
            statusMessage.style.color = "#44ff44";
            statsRow.style.display = 'none';
            finalPath = path;
            
            let countdown = 3;
            const tick = () => {
                if (countdown > 0) {
                    statusMessage.textContent = t("Update finished!") + " (" + countdown + ")";
                    countdown--;
                    setTimeout(tick, 1000);
                } else {
                    statusMessage.textContent = t("Update finished!");
                    window.go.main.App.LaunchAndExit(finalPath);
                }
            };
            tick();
>>>>>>> 1e7144e (1.0.0)
        });

        window.go.main.App.RunUpdate();
        
    } catch (e) {
        statusMessage.textContent = "Error: " + e;
    }
}

// Give wails bindings a moment to load before calling
setTimeout(initUpdater, 500);
