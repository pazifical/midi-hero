document.getElementById("files").onchange = (event) => {
  const fileNames = [];
  console.log(event.target.files);

  for (const f of event.target.files) {
    console.log(f.name);
    fileNames.push(f.name);
  }

  document.getElementById("submit").style.display = "inline-block";
  document.getElementById("info").style.display = "block";

  let selectedDiv = document.getElementById("selected-files");
  selectedDiv.innerHTML =
    "<ul>" + fileNames.map((f) => `<li>${f}</li>`).join(" ") + "</ul>";
};

async function openConverted() {
  await fetch("/api/open");
}
