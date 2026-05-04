# Team: Romanshk Volkov - https://github.com/RomanshkVolkov
# Team: Diegode - https://github.com/diegode-tsx
# Team: Alexandergv2117 - https://github.com/Alexandergv2117

FROM node:lts-alpine AS builder
WORKDIR /app

# Install packages
COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

# NODE ADAPTER SVELTE
FROM node:lts-alpine AS runner
WORKDIR /app

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 svelte

COPY --from=builder /app/build/ .
COPY --from=builder /app/node_modules /app/node_modules

# Set ownership for the svelte user
RUN chown -R svelte:nodejs /app

USER svelte

EXPOSE 3000

CMD ["node", "index.js"]
