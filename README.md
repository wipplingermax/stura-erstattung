# Stura Erstattung
Entwurf einer Web-App zur Rückerstattung der anteiligen Semesterbeiträge über den Zeitraum des 9€-Tickets

## Tech-Stack
- Frontend:
  - NextJS 13 (inkl. appDir beta)
  - Typescript
  - TailwindCSS

- Backend:
  - Golang
  - GORM
  - Gin

## Run Frontend-Dev-Container
**Achtung:** Hier wird bewusst der Ordner ``frontend`` live in den Container eingebunden. 
Werden im Docker-Container Dateien erstellt, so geschieht dies durch den User ``root``. 

Der Container kann via ``docker-compose`` gestartet werden
```
docker-compose -f docker-compose-frontend.yml up -d
```

## Run Database-Dev-Container
Enthaltenes Setup baut auf zwei Containern auf:

- postgres (Datenbank)
- pgadmin (Webanwendung zur Administration der Datenbank)

Der pgadmin Container wird bereits vorkonfiguriert zur Verfügung gestellt, sodass er bereits eine Verbindung mit der Datenbank herstellt. Hierzu werden in einem Dockerfile Konfiguationsdaten in das Image kopiert.

Die Services laufen durch Portfreigabe unter folgenden Adressen:

- Postgres: ``127.0.0.1:5432``
- PgAdmin: ``127.0.0.1:80`` 

Die Login-Daten für PgAdmin sind:
- Username: ``admin@example.com``
- Passwort: ``admin``

Die Container können via ``docker-compose`` gestartet werden
```
docker-compose -f docker-compose-database-dev.yml up -d
```
