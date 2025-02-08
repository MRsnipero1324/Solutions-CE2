# Dockerfile para Node.js
FROM node:20-alpine

WORKDIR /app

COPY ../JavaScript/Script.js /app/

CMD ["node", "Script.js"]
