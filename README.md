# go_restapi

Чтобы создать таблиц sql :
	migrate create -ext sql -dir migration create_user
	
Будут созданы 2 slq файла в которых мы должны прописать поля для базы дынных

Команда миграции :
	migrate -path path -database "username:password@host:port/dbname?param1=true&param2=false" up/down

