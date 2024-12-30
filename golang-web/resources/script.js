// Menangkap elemen form
const form = document.getElementById('userForm');

// Menambahkan event listener pada submit
form.addEventListener('submit', function(event) {
    // Mendapatkan nilai input
    const nama = document.getElementById('nama').value.trim();
    const email = document.getElementById('email').value.trim();

    // Validasi sederhana
    if (!nama || !email) {
        // Mencegah pengiriman form jika input kosong
        event.preventDefault();
        alert('Harap isi semua field sebelum mengirim!');
    } else if (!validateEmail(email)) {
        // Validasi email
        event.preventDefault();
        alert('Harap masukkan alamat email yang valid!');
    }
});

// Fungsi untuk memvalidasi email menggunakan regex
function validateEmail(email) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
}
