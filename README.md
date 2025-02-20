# portofolio-rest-api

cmd -> folder/package => tempat kita naruh file main.go -> buat web framework
internal -> untuk membuat rest api (semua yang berkaitan dengan api)
infrastructure -> konfigurasi third party. co:/konekin DB  

install gin framework
go get -u github.com/gin-gonic/gin

install gorm dan driver postgre
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

buat database namanya "portofolio"
buat enum manual di dbeaver
CREATE TYPE social_media_type AS ENUM ('primary', 'secondary', 'tertiary');
CREATE TYPE activity_type AS ENUM ('pelatihan', 'kompetisi', 'seminar', 'volunteer');

*pakai validator (belum di terapkan) = sudah

tabel yang belum dan urutannya
1. skills OK
2. experiences OK
3. activities
4. education
5. project_category
6. projects
7. admins
8. authentication (cari method yang paling mudah (JWT/Cookies/Session))
