## Membuat Go Module ##
1. Buat repo di GitHub/GitLab : go-say-hello
2. buat folder baru di local dan beri nama sama dengan repository yg telah di buat sebelumnya :go-say-hello
3. git init dan daftarkan remote repository ke repo local
4. commit dan push ke repository

## Rilis versi go module ##
```git tag v1.0.0
git push origin v1.0.0
```

## Menambah Dependency ke Project ##
1. masuk ke folder project golang
2. jalankan perintah : 
   ```go get <url github/GitLab module> ```
   contoh : 
   ```go get gitlab.com/pramonoagung/go-say-hello```
 
## Upgrade Module ##
cukup membuat tag baru
contoh :
ubah versi contoh v1.0.0 => v1.1.0 
```git tag v1.0.0
git push origin v1.0.0
```

## Upgrade Dependency ##
1. ubah versi dependency yg sudah terdaftar di go.mod
2. jalankan perintah : go get

## Major Upgrade ##
ubah nama module tanpa merusak backward compatibility yg sudah ada

gitlab.com/pramonoagung/app-say-hello/v2

lalu lakukan tag lagi dengan versi yg berbeda : 
```git tag v2.0.0
git push origin v2.0.0
```
