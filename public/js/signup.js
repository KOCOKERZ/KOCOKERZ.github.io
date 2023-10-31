document.getElementById('signupForm').addEventListener('submit', function (e) {
    e.preventDefault();

    const username = document.getElementById('usernameSignup').value;
    const password = document.getElementById('passwordSignup').value;

    const userData = {
        Username: username,
        Password: password
    };

    fetch('https://us-central1-core-advice-401502.cloudfunctions.net/REGIS-GIS', {
        method: 'POST',
        body: JSON.stringify(userData),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Pendaftaran gagal');
        }
    })
    .then(data => {
        // Handle respons dari server (misalnya, tampilkan pesan sukses)
        console.log('Pendaftaran berhasil:', data);
        // Redirect ke halaman login
        window.location.href = '../../public/pages/signupBerhasil.html'; // Ganti dengan nama file halaman login yang sesuai
    })
    .catch(error => {
        // Handle kesalahan (misalnya, tampilkan pesan kesalahan)
        console.error('Pendaftaran error:', error);
    });
});