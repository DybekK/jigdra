DOCKER=$(docker ps | grep postgres | cut -f1 -d" ")
docker exec -i $DOCKER psql -d jigdra -U admin -W < init.sql