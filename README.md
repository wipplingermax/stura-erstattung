# Stura Erstattung
Entwurf einer Web-App zur Rückerstattung der anteiligen Semesterbeiträge über den Zeitraum des 9€-Tickets

## Tech-Stack
- Frontend:
  - NextJS 13 (inkl. appDir beta)
  - Typescript
  - TailwindCSS

- Backend:
  - Golang + Gin

## Run Frontend-Dev-Container
**Achtung:** Hier wird bewusst der Ordner ``frontend`` live in den Container eingebunden. 
Werden im Docker-Container Dateien erstellt, so geschieht dies durch den User ``root``. 

Frontend-Dev-Docker Container via ``docker-compose`` starten
```
docker-compose -f docker-compose-frontend.yml
```
