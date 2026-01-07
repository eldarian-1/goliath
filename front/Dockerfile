FROM node:latest

WORKDIR /app

COPY package.json .

# Remove package-lock.json and node_modules if they exist, then install dependencies
# This fixes the issue with optional dependencies on ARM64 (rollup native modules)
RUN rm -rf package-lock.json node_modules && \
    npm install --include=optional

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev"]
