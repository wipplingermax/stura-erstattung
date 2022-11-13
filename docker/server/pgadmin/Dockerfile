FROM 'dpage/pgadmin4':latest

COPY ./config /var/lib/pgadmin

USER root

RUN sudo chown -R 5050:5050 /var/lib/pgadmin