document.addEventListener("DOMContentLoaded", function() {
    const loginForm = document.getElementById("loginForm");
    // const message = document.getElementById("message");

    loginForm.addEventListener("submit", function(event) {
        event.preventDefault();
        
        const username = document.getElementById("usernameLogin").value;
        const password = document.getElementById("passwordLogin").value;

        // Kirim permintaan POST ke API
        fetch("https://us-central1-core-advice-401502.cloudfunctions.net/GCFPostHandler", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
        })
        .then(response => response.json())
        .then(data => {
            if (data.status === true) {
                // Pengolahan respons setelah login berhasil
                const token = data.token;
                // const welcomeMessage = data.message;
                // message.textContent = welcomeMessage;
                // message.style.color = "green";
                console.log(token);

                // Redirect ke halaman dashboard setelah login berhasil
                window.location.href = "../../public/pages/loginBerhasil.html";
            } else {
                // Pengolahan respons jika login gagal
                // message.textContent = "Login failed. Please check your username and password.";
                // message.style.color = "red";
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
    });
});