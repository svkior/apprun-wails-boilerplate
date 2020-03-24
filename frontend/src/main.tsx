import app from "apprun"

import "core-js"
import * as Wails from '@wailsapp/runtime';

require('./app.pcss')


declare global {
	interface Window { wails: any; backend: any; }
}

window.wails = window.wails || {};
window.backend = window.backend || {};

const state = {
	status: "Initial Status"
};


const view = (state) => (
<div class="max-w-sm rounded overflow-hidden shadow-lg">
  <div class="px-6 py-4">
    <div class="font-bold text-xl mb-2">Current Status: {state.status}</div>
    <p class="text-gray-700 text-base">
      Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus quia, nulla! Maiores et perferendis eaque, exercitationem praesentium nihil.
    </p>
  </div>
  <div class="px-6 py-4">
  <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
  	Increment
  </button>
  <button 
	  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
	  onclick={() => app.run('clear_history')}
  >
  	Clear History
  </button>
  </div>
</div>
);

const STORAGE_KEY = "demoApp";

const update = {
	"clear_history": (state, al) => {
		localStorage.removeItem(STORAGE_KEY);
		document.location.reload();
	},
	"update_status": (state, newStatus) => ({
		...state,
		status: newStatus
	})
};

const rendered = state => localStorage.setItem(STORAGE_KEY, JSON.stringify(state));
const stored = localStorage.getItem(STORAGE_KEY);


Wails.Init(() => {

	app.start(document.body, stored ? JSON.parse(stored) : state, view, update, { history: true, rendered })

	Wails.Events.On("update_status", (newStatus) => {
		//console.log("On STATUS: ", newStatus)
		app.run("update_status", newStatus)
	  });

	for(var i=0; i < 100; i++){
		window.backend.MyBridge.SetStatus(`Count: ${i}`)
	}
});
