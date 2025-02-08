# Dockerfile para Node.js
FROM node:18-alpine

WORKDIR /app

COPY ../JavaScript/Script.js /app/

CMD ["node", "Script.js"]
