
# 📝 TO-DO App (Backend)

Bu proje, staj değerlendirme süreci kapsamında geliştirilen bir **TO-DO uygulamasının backend** servisidir. Uygulama, kullanıcıların yapılacaklar listesi oluşturup bu listelere maddeler (adımlar) eklemesine olanak tanır. Geliştirme dili olarak **GoLang** seçilmiş, web çatısı olarak ise **Gin Gonic Framework** kullanılmıştır.

## 🚀 Özellikler

- ✅ JWT (JSON Web Token) ile kimlik doğrulama
- ✅ İki farklı kullanıcı tipi: **admin** ve **user**
- ✅ TO-DO listesi CRUD işlemleri
- ✅ Liste içi madde (adım) CRUD işlemleri
- ✅ Soft delete (silinme tarihi ile)
- ✅ Güncelleme tarihi kaydı
- ✅ Mock servis üzerinden veri saklama (veritabanı bağımsız)
- ✅ Yetki kontrolü: kullanıcılar sadece kendi verilerini yönetebilir, admin tüm verileri görebilir


## ✍️ Notlar

- Proje Clean Architecture ve MVC mimarisine uygun olarak yapılandırılmıştır.
- Veriler geçici olarak bellekte (mock veri) tutulur, veritabanı bağlantısı içermez.
- Silme işlemleri kalıcı değildir, sadece `DeletedAt` tarihi atanır.


## ℹ️ Not: Commit Sayısı Hakkında

Proje geliştirme sürecinde local ortamda düzenli çalıştım ve birçok aşamada ilerleme kaydettim. Ancak Git versiyonlamasına başta yeterince dikkat etmediğim için, bazı geliştirme adımlarını `commit` ile kaydetmeyi atladım. Bu yüzden son aşamada, projenin tüm fonksiyonel gereksinimlerini tamamlayıp tek seferde versiyonladım. Commit sayısının az olmasından dolayı bir açıklama gereği hissettim.
 
Teşekkür ederim.


## 👤 Kullanıcılar (Ön Tanımlı)

| Kullanıcı Adı | Şifre     | Rol   |
|---------------|-----------|-------|
| `admin`       | `admin123`| admin |
| `user`        | `user123` | user  |

Bu kullanıcılar `mock_data.go` dosyasında tanımlıdır.

## 🧪 API Endpointleri

- `/login` dışında tüm `GET`, `POST`, `PUT`, `DELETE` istekleri JWT token gerektirir.

### 🔐 /login
Kullanıcı giriş yapar ve geçerli bilgilerle JWT token alır.
```json
POST /login
{
  "username": "admin",
  "password": "admin123"
}
```


### 📋 TO-DO Listeleri

#### ➕ Liste Oluştur
```http
POST /createTodo
Body: { "name": "Alışveriş Listesi" }
```


#### 📖 Liste Görüntüle
```http
GET /getTodos
```


#### ✏️ Liste Güncelle
```http
PUT /updateTodo/:id
Body: { "name": "Yeni Liste Adı" }
```


#### 🗑️ Liste Sil
```http
DELETE /deleteTodo/:id
```

---


### ✅ Liste Maddeleri (Items)

#### ➕ Madde Ekle
```http
POST /lists/:id/addItem
Body: { "content": "Süt al" }
```


#### 📖 Maddeleri Listele
```http
GET /lists/:id/getItems
```

#### ✏️ Madde Güncelle
```http
PUT /updateItem/:id
Body: { "content": "Süt ve ekmek al", "isDone": true }
```


#### 🗑️ Madde Sil
```http
DELETE /deleteItem/:id
```


## 📁 Proje Yapısı

```bash
todo-app/
│
├── main.go               # Uygulama giriş noktası
├── routes/               # Tüm route tanımlamaları
├── controllers/          # HTTP endpoint işlemleri
├── middleware/           # JWT kontrolü
├── models/               # Struct tanımları
├── data/                 # Mock veri işlemleri
└── utils/                # JWT üretim/fonksiyonları
```


## 📦 Kullanılan Teknolojiler

- GoLang
- Gin Gonic Framework
- JWT (github.com/golang-jwt/jwt/v4)


## 🛠️ Kurulum ve Çalıştırma

```bash
git clone https://github.com/furkanyld/TO-DO-APP.git
cd TO-DO-APP
go run main.go
```

Uygulama `http://localhost:8080` üzerinden çalışır.


## 🔒 Yetki Kuralları

- **admin** → Tüm kullanıcıların listelerini ve maddelerini görebilir/güncelleyebilir/silebilir.
- **user** → Sadece kendi oluşturduğu liste ve maddeleri görebilir/güncelleyebilir/silebilir.


## 🌐 Yayın

Proje şu adresten erişilebilir:  
👉 [https://github.com/furkanyld/TO-DO-APP](https://github.com/furkanyld/TO-DO-APP) *(private repo)*


## 📩 Geliştirici

**Furkan Yıldız**  
[www.linkedin.com/in/furkan-yıldız-584383254]
[github.com/furkanyld]

