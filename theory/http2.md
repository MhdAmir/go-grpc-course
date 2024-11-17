# HTTP2 vs Htt1.1
### How HTTP1 Work
- **TCP connection per request**
- **Plaintext headers** : mengirim header dalam format teks biasa (plaintext) yang memakan lebih banyak bandwidth daripada format biner yang digunakan pada HTTP/2.
- **only accept request response**: 
    Pada HTTP/1.1, client perlu melakukan request ke server setiap beberapa detik untuk mengecek apakah ada pesan baru. Request ini membuat aplikasi kurang efisien karena banyak permintaan yang dilakukan hanya untuk memeriksa data baru, meskipun seringkali tidak ada update.
    
    **Solusi Lanjutan**
    
    Untuk mengatasi keterbatasan ini, beberapa teknologi berikut digunakan:
    - WebSocket: Membuat koneksi full-duplex yang memungkinkan komunikasi dua arah, sehingga server bisa mengirim data kapan saja tanpa perlu menunggu permintaan dari client.

    - HTTP/2 Server Push: HTTP/2 mendukung fitur "server push," yang memungkinkan server mengirim data ke client sebelum client memintanya.

### How to HTTP2 Works
- **one TCP Connection:** long lasting connection which will be shared by multiple request & response
- **Server Push:** only one request from client doesn't need ask for more data, it just wait the server push the data when data is ready
- **multiplexing**: server & client can push multiple messages in parallel on same TCP connection
- **Headers Binary**
- **more secure cause SSL**

 - less bandwith
 - More Security
 - Less Chatter