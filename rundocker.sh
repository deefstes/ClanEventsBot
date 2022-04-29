docker run --rm \
-p 8080:8080 \
-e CEB_TOKEN=NDI4ODMzNTUyOTIzMzYxMjgx.DZ417g.08RN1xgLPzrSbBqhchcKimI-Vlk \
-e CEB_CMDPREFIX="~" \
-e CEB_MONGODB=mongodb://host.docker.internal:27017 \
-e CEB_PORT=8080 \
-e CEB_APIKEY=c2e0dabe-bdb5-41ee-a2db-9140c5748473 \
-e CEB_SVCTIMER=60 \
-e CEB_LOGLEVEL=2 \
claneventsbot:latest