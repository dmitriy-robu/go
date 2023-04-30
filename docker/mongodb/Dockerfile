FROM mongo:latest

ENV MONGO_KEYFILE_PATH /etc/mongo/keyfile

COPY ./mongodb-keyfile ${MONGO_KEYFILE_PATH}

RUN chown mongodb:mongodb ${MONGO_KEYFILE_PATH} && \
    chmod 400 ${MONGO_KEYFILE_PATH}
