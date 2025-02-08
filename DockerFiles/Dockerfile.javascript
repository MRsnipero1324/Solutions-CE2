# Dockerfile para Node.js
FROM node:20-slim

WORKDIR /app

COPY ../JavaScript/Script.js /app/

CMD ["node", "script.js"]
