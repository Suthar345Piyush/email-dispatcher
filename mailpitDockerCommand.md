-> for windows 

docker run -d --restart unless-stopped --name mailpit -p 8025:8025 -p 1025:1025 axllent/mailpit


-> for mac and linux

docker run -d \
--restart unless-stopped \
--name=mailpit \
-p 8025:8025 \
-p 1025:1025 \
axllent/mailpit



