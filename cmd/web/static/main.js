document.getElementById("file").onchange = (event) => {
  let fileName;
  try {
    fileName = event.target.files[0].name;
  } catch (e) {
    console.log(e);
    return;
  }

  let info = document.getElementById("info");
  info.innerHTML = `Selected file: ${fileName}`;
  info.style.display = "block";
};
