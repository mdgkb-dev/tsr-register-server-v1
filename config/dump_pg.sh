#!/bin/sh
PGPASSWORD=$1 dropdb -Umdgkb -hlocalhost tsr
PGPASSWORD=$1 createdb -Umdgkb tsr
ssh root@45.67.57.208 "pg_dump -C -h 45.67.57.208 -d tsr -U mdgkb -Fc --password" | PGPASSWORD=$1 pg_restore -Umdgkb -hlocalhost --format=c -dtsr