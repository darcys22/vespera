import './style.css';
import './app.css';

import {ToggleLight, ControllerStatus} from '../wailsjs/go/vespera/State';

document.querySelector('#app').innerHTML = `
    <div class="result" id="lightResult"></div>
    <button class="btn" onclick="toggleLight()">Toggle Light</button>
    <div class="result" id="xboxResult"></div>

    </div>
`;

let lightResultElement = document.getElementById("lightResult");
let xboxResultElement = document.getElementById("xboxResult");

// Setup the toggleLight function
window.toggleLight = function () {

    // Call State.ToggleLight()
    try {
        ToggleLight()
            .then((result) => {
                // Update result with data back from App.Greet()
                lightResultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

// Setup the toggleLight function
window.heartbeat = function () {

    // Call State.ControllerStatus()
    try {
        ControllerStatus()
            .then((result) => {
                // Update result with data back from App.Greet()
                xboxResultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
  setTimeout(window.heartbeat, 300)
};
window.heartbeat();
