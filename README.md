# Stura Rückerstattung
Entwurf einer Web-App zur Rückerstattung der anteiligen Semesterbeiträge über den Zeitraum des 9€-Tickets

## Tech-Stack
- Frontend:
  - NextJS 13 (inkl. appDir beta)
  - Typescript
  - TailwindCSS

- Backend:
  - Golang + Gin

## Run Frontend-Dev-Container
**Achtung:** Hier wird bewusst der Ordner ``frontend`` live in den Container eingebunden

Damit keine Probleme mit den Dateiberechtigungen entstehen, ist es sinnvoll der Docker-Entwicklungsumgebung den aktuellen Nutzer mitzugeben.
Die docker-compose-frontend.yml Datei wurde entsprechend angepasst

Um nextJS im Dev-Modus im Container zu starten sollten folgende Schritte erfolgen: 

```
DOCKER_USER="$(id -u):$(id -g)"
```

```
docker-compose -f docker-compose-frontend.yml
```
