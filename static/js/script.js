document.getElementById("generateBtn").addEventListener("click", async () => {
  try {
    const res = await fetch("http://localhost:8080/email", {
      method: "POST"
    });
    const data = await res.json();

    const emailID = data.id;
    const email = `${emailID}@tempdomain.com`;

    document.getElementById("emailAddress").textContent = email;
    document.getElementById("inboxLink").href = `/inbox.html?emailID=${emailID}`;
    document.getElementById("emailBox").classList.remove("hidden");
  } catch (err) {
    alert("Something went wrong! Make sure the backend is running.");
    console.error(err);
  }
});

