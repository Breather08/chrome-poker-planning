import "./style.css";

const root = document.querySelector<HTMLDivElement>("#app");

const input = document.createElement("input");
input.type = "file";

interface SendFileRequest {
  filename: string;
}

input.addEventListener("change", (e) => {
  const reader = new FileReader();
  const file = e.target?.files[0] as File;
  reader.readAsArrayBuffer(file);

  reader.onload = function () {
    console.log("Sending File", reader.result);
    const metadata: SendFileRequest = { filename: file.name };
    socket.send(JSON.stringify(metadata));

    socket.send(file);
  };

  reader.onerror = function () {
    console.error(reader.error);
  };
});

root?.append(input);

const socket = new WebSocket("ws://192.168.0.44:8080/ws");

socket.addEventListener("error", (e) => {
  // Send a message
  console.log("Error", e);
});

const filenames = document.createElement("div");
filenames.className = "filenames";
document.body.append(filenames);

let fileName = "";

socket.addEventListener("message", async (e) => {
  // Send a message
  console.log("Message", e.data);

  attachFile(e.data);
});

function attachFile(data: unknown) {
  if (typeof data === "string") {
    console.log("Filename Received");
    fileName = data;
  } else if (data instanceof Blob) {
    console.log("Blob Received");
    const url = URL.createObjectURL(data);
    const a = document.createElement("a");

    a.href = url;
    a.download = fileName; // Customize the file name
    document.body.appendChild(a);
    const btn = document.createElement("button");
    btn.textContent = fileName;
    btn.addEventListener("click", () => {
      a.click();
      setTimeout(() => {
        btn.remove();
        URL.revokeObjectURL(url);
      });
    });
    btn.textContent = fileName;
    filenames.append(btn);
  }
}
