# Dockerfile para Node.js
FROM node:20-slim

WORKDIR /app

COPY script.js /app/

CMD ["node", "script.js"]
