FROM node:18-alpine AS base
 
FROM base AS builder
RUN apk update
RUN apk add --no-cache libc6-compat
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
# Set working directory
WORKDIR /app
# Replace <your-major-version> with the major version installed in your repository. For example:
RUN pnpm install turbo --global
COPY . .
 
# Generate a partial monorepo with a pruned lockfile for a target workspace.
# Assuming "web" is the name entered in the project's package.json: { name: "web" }
RUN turbo prune @srsim/web --docker
 
# Add lockfile and package.json's of isolated subworkspace
FROM base AS installer
RUN apk update
RUN apk add --no-cache libc6-compat
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app
 
# First install the dependencies (as they change less often)
COPY --from=builder /app/out/json/ .
RUN pnpm install
 
# Build the project
COPY --from=builder /app/out/full/ .
RUN pnpm dlx turbo run build --filter=@srsim/web...

RUN ls apps
RUN ls -la apps/web
RUN ls -la apps/web/.next/
RUN ls -la apps/web/.next/standalone
RUN ls -la apps/web/public
 
FROM base AS runner
WORKDIR /app

ENV NODE_ENV production
ENV NEXT_TELEMETRY_DISABLED 1

# Don't run production as root
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

# Automatically leverage output traces to reduce image size
# https://nextjs.org/docs/advanced-features/output-file-tracing
COPY --from=installer --chown=nextjs:nodejs /app/apps/web/.next/standalone ./
COPY --from=installer --chown=nextjs:nodejs /app/apps/web/.next/static ./apps/web/.next/static
COPY --from=installer --chown=nextjs:nodejs /app/apps/web/public ./apps/web/public

RUN ls ./
RUN ls ./apps/web
RUN ls ./apps/web/.next
RUN ls ./apps/web/.next/static
RUN ls ./apps/web/public

USER nextjs

EXPOSE 3000

ENV PORT=3000
 
CMD HOSTNAME="0.0.0.0" node apps/web/server.js
