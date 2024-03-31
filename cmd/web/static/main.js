document.getElementById("file").onchange = (event) => {
  let fileName;
  try {
    fileName = event.target.files[0].name;
  } catch (e) {
    console.log(e);
    return;
  }

  document.getElementById("info").style.display = "block";
  document.getElementById("selected-file").innerText = `"${fileName}"`;
};
