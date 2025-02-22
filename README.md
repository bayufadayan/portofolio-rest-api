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
3. activities OK
4. education OK
5. project_category OK
6. projects OK
7. admins 
8. authentication (cari method yang paling mudah (JWT/Cookies/Session))

deploy
git clone https://github.com/bayufadayan/portofolio-rest-api.git .
go mod tidy
go build -o app ./cmd/web
supervisorctl -c /home/bayufada/api.bayufadayan.my.id/conf/api.bayufadayan.my.id-gin-app.conf reread
supervisorctl -c /home/bayufada/api.bayufadayan.my.id/conf/api.bayufadayan.my.id-gin-app.conf update
supervisorctl -c /home/bayufada/api.bayufadayan.my.id/conf/api.bayufadayan.my.id-gin-app.conf restart api.bayufadayan.my.id-gin-app

CREATE TYPE social_media_type AS ENUM ('primary', 'secondary', 'tertiary');
CREATE TYPE activity_type AS ENUM ('pelatihan', 'kompetisi', 'seminar', 'volunteer');

cat /home/bayufada/logs/api.bayufadayan.my.id-gin-app.log
/home/bayufada/logs/api.bayufadayan.my.id-gin-app.log

command=/home/bayufada/api.bayufadayan.my.id/app
user=bayufada
directory=/home/bayufada/api.bayufadayan.my.id
environment=ENV_PATH="/home/bayufada/api.bayufadayan.my.id/.env"
autostart=true
autorestart=true
stderr_logfile=/home/bayufada/logs/api.bayufadayan.my.id-gin-app.err.log
stdout_logfile=/home/bayufada/logs/api.bayufadayan.my.id-gin-app.out.log

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=bayufada_bayufadayan
DB_PASSWORD=Myseruling1793*
DB_NAME=bayufada_portofolio

mv portofolio-rest-api/* portofolio-rest-api/.* .


