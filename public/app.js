const socket = new WebSocket("ws://localhost:4000");

socket.onopen = () => console.log("connected");

socket.onmessage = (event) => {
  const content = event.data;
  document.getElementById("chat").insertAdjacentHTML("beforeend", template("robot", content));
};

socket.onclose = (event) => console.log("disconnected", event);

socket.onerror = (event) => console.log("error", event);

document.getElementById("send").onclick = () => {
    const content = document.getElementById("content");
    socket.send(content.value);
    document.getElementById("chat").insertAdjacentHTML("beforeend", template("user", content.value));
    content.value = "";
};


const template = (user, content) => {
  const name = user == "robot" ? "Seu Robô" : "Eu Mesmo";
  return `
      <div class="bg-white shadow-md rounded-lg p-4 mt-4">
      <div class="flex ${user != "robot" ? "justify-end" : ""}">
      <img src="/public/${user}.png"  class="h-10 w-10 rounded-full ml-4"/>
      <div>
          <div class="font-semibold text-lg">${name}</div>
          <div class="text-gray-600">${content}</div>
      </div>
      </div>
      </div>
    `;
};