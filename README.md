# Boşluk Kontrolü API

Bu basit Go programı, verilen bir metnin boşluk karakterleri içerip içermediğini kontrol etmek için bir HTTP endpoint'i sağlar.

## Nasıl Kullanılır

1. Programı çalıştırmak için bir Go geliştirme ortamınızın kurulu olduğundan emin olun.

2. Proje dizinine gidin ve HTTP sunucusunu başlatmak için `go run main.go` komutunu çalıştırın.

3. Metni isteğin gövdesinde bulunan bir POST isteği göndererek `http://localhost:8080/whiteSpaceController` adresine gönderin.

4. API, metinde boşluk karakterleri bulunuyorsa "Metin boşluk içeriyor", bulunmuyorsa "Metin boşluk içermiyor" yanıtını verecektir.

## Örnek

### İstek
```plaintext
POST /whiteSpaceController HTTP/1.1
Content-Type: text/plain

Bu bir örnek metindir
