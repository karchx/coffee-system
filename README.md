# Coffee-system starter

This is an official coffe-system repo.

__coffee-sytem is a shop coffee based in microservices__

## What's inside?
It includes the following packages/apps:

### Apps and Packages

- `docs`: a [Next.js](https://nextjs.org/) app with documentation project
- `web`: another [Next.js](https://nextjs.org/) main web system
- `ui`: a stub Frontend component library shared by both `web` and `docs` applications
- `backend`: a [Go]() app with microservices
- `eslint-config-custom`: `eslint` configurations (includes `eslint-config-next` and `eslint-config-prettier`)
- `tsconfig`: `tsconfig.json`s used throughout the monorepo

Each package/app is 100% [TypeScript](https://www.typescriptlang.org/).

### Docker
```
docker-compose up --build -d
```

### Utilities

This turborepo has some additional tools already setup for you:

- [TypeScript](https://www.typescriptlang.org/) for static type checking
- [ESLint](https://eslint.org/) for code linting
- [Prettier](https://prettier.io) for code formatting
- [Docker](https://www.docker.com/) containers
- [Go](https://go.dev/) for backend language
