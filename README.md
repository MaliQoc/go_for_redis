# go_for_redis

Bu proje, Go dilinde Redis ile temel veri yapılarını kullanmayı gösteren bir örnektir. Aşağıdaki Redis veri yapılarını kullanarak temel işlemleri gerçekleştirir:

- String: Bir anahtar-değer çifti olarak veri saklama.
- Hash: Bir anahtar altında birden fazla alan ve değeri saklama.
- List: Sıralı bir veri listesi oluşturma.
- Set: Benzersiz elemanlardan oluşan bir küme oluşturma.
- Sorted Set: Skor (puan) ile sıralı bir küme oluşturma.


1) Kod Açıklaması
   
- Bağlantı Kurulumu: Redis sunucusuna bağlanılır.
- Ping: Redis sunucusuna bağlantıyı doğrulayan bir PING komutu gönderilir.
- String Operasyonları: name anahtarına değer atanır ve okunur.
- Hash Operasyonları: user:1 anahtarında bir hash oluşturulur ve okunur.
- Liste Operasyonları: my_list anahtarına elemanlar eklenir ve okunur.
- Küme Operasyonları: my_set anahtarına elemanlar eklenir ve okunur.
- Sıralı Küme Operasyonları: my_zset anahtarına sıralı elemanlar eklenir ve okunur.


2) Çalıştırma aşamaları

- Terminalde go get github.com/go-redis/redis/v8 yazın ve çalıştırın. Github dan redis kütüphanesini çekmiş olacaksınız.
- Kodu çalıştırırken go run main.go yazıp çalıştıracaksınız.

ÖNEMLİ! Redis Windows'ta desteklenmemektedir. Docker Desktop indirip kurmanız gerekiyor -> [Docker Linki] [https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe?utm_source=docker&utm_medium=webreferral&utm_campaign=dd-smartbutton&utm_location=module]

Redisi Docker içerisinde kurduktan sonra örneğin docker run -d --name myredis -p 6380:6379 redis gibi bir kod girebilirsiniz. Bu komutta, ana makinedeki 6380 numaralı portu Redis konteynerinin 6379 numaralı portuna yönlendirir. Bu durumda, Redis sunucusuna erişmek için ana makinenizin localhost:6380 adresini kullanmanız gerekir. 

myredis konteynerın adıdır. Docker'da konteynerlar, izole edilmiş çalışma ortamları sağlar ve her bir konteyner kendi adını taşır. Bu isim, konteyneri Docker içinde tanımlamak ve yönetmek için kullanılır.

docker ps komutu çalışan konteynerları listeler.


3) Kaydettiğiniz verileri görme


- Visual Studio Code terminalinde docker exec -it myredis redis-cli yazıp ardından

1- GET name
  
2- HGETALL user:1

3- LRANGE my_list 0 -1

4- SMEMBERS my_set

5- ZRANGE my_zset 0 -1 gibi komutlarla yazdığınız verileri türlerine göre listeleyebilirsiniz.

Komutların çoğunu https://www.yusufsezer.com.tr/redis/ linkinde bulabilirsiniz.


- Docker içerisinde ise data kısmında .rdb uzantısıyla kaydedilen dosyayı açarak verinizi görebilirsiniz. 

