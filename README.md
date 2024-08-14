<p align="center">
    <img src="./front/src/assets/logo.svg" alt="logo" />
    <h1 align="center">Elite-network</h1>
</p>

<h2 align="center">About The Project</h2>

Welcome to our version of social-network Project!

 where we have to create a Facebook-like social network with features such as followers, profiles, posts, groups, notifications, and chats. This project involves both frontend and backend development using a JavaScript framework for the client side and Go for the server side. The backend handles user authentication, image management, WebSocket connections for real-time chats, and uses PostgreSQL for data storage with Docker containers for deployment.



## Table of Contents

-   [Getting Started](#getting-started)
-   [Usage](#usage)
-   [Directory Structure](#directory-structure)
-   [Authors](#authors)

## Getting Started

You can run the Lem-In project with the following command:

```console
git clone github.com/01Elites/social-network
cd social-network
```

## Usage

_make sure you are in project directory_

```
bash build.sh
```

1. Open http://localhost:3000/ in a browser .
2. Sign up to access the website; it's exclusively for registered users.

or vist: https://elite-network.extbh.dev/


### Directory Structure
<details>
<summary>tree map</summary>

```
── social-network
    ├── build.sh
    ├── Caddyfile
    ├── cmd
    │   └── socialNetwork
    │       └── main.go
    ├── docker-compose.yml
    ├── Dockerfile.caddy
    ├── front
    │   ├── index.html
    │   ├── package.json
    │   ├── package-lock.json
    │   ├── postcss.config.cjs
    │   ├── README.md
    │   ├── src
    │   │   ├── assets
    │   │   │   ├── bell.svg
    │   │   │   ├── favicon.ico
    │   │   │   ├── icons_svgs
    │   │   │   │   ├── follow.svg
    │   │   │   │   ├── github.svg
    │   │   │   │   └── globe.svg
    │   │   │   ├── logo.svg
    │   │   │   ├── reboot_01_logo.png
    │   │   │   ├── sample.avif
    │   │   │   └── svg-loaders
    │   │   │       └── tail-spin.svg
    │   │   ├── components
    │   │   │   ├── Chat
    │   │   │   │   ├── chatMessage.tsx
    │   │   │   │   └── index.tsx
    │   │   │   ├── core
    │   │   │   │   ├── navigation
    │   │   │   │   │   └── index.tsx
    │   │   │   │   ├── repeat
    │   │   │   │   │   └── index.tsx
    │   │   │   │   └── textbreaker
    │   │   │   │       └── index.tsx
    │   │   │   ├── EditProfileDialog
    │   │   │   │   └── index.tsx
    │   │   │   ├── Feed
    │   │   │   │   ├── FeedPostCellSkeleton.tsx
    │   │   │   │   ├── FeedPostCell.tsx
    │   │   │   │   ├── FeedPosts.tsx
    │   │   │   │   ├── index.tsx
    │   │   │   │   ├── NewGroupPostCell.tsx
    │   │   │   │   └── ...
    │   │   │   ├── HomeContacts
    │   │   │   │   └── index.tsx
    │   │   │   ├── HomeEvents
    │   │   │   │   └── index.tsx
    │   │   │   ├── LoginDialog
    │   │   │   │   └── index.tsx
    │   │   │   ├── PostAuthorCell
    │   │   │   │   └── index.tsx
    │   │   │   └── ui
    │   │   │       ├── aspect-ratio.tsx
    │   │   │       ├── avatar.tsx
    │   │   │       ├── button.tsx
    │   │   │       ├── card.tsx
    │   │   │       ├── grid.tsx
    │   │   │       ├── label.tsx
    │   │   │       ├── select.tsx
    │   │   │       ├── separator.tsx
    │   │   │       └── ...
    │   │   ├── config
    │   │   │   └── index.ts
    │   │   ├── contexts
    │   │   │   ├── NotificationsContext
    │   │   │   │   └── index.ts
    │   │   │   ├── UserDetailsContext
    │   │   │   │   └── index.ts
    │   │   │   └── WebSocketContext
    │   │   │       └── index.ts
    │   │   ├── extensions
    │   │   │   ├── arrays.ts
    │   │   │   ├── fetch.ts
    │   │   │   ├── File.ts
    │   │   │   └── index.ts
    │   │   ├── hooks
    │   │   │   ├── NotificationsHook
    │   │   │   │   └── index.ts
    │   │   │   ├── userDetails
    │   │   │   │   └── index.ts
    │   │   │   └── WebsocketHook
    │   │   │       └── index.ts
    │   │   ├── index.css
    │   │   ├── index.tsx
    │   │   ├── Layout.tsx
    │   │   ├── lib
    │   │   │   └── utils.ts
    │   │   ├── pages
    │   │   │   ├── events
    │   │   │   │   ├── eventsfeed.tsx
    │   │   │   │   └── index.tsx
    │   │   │   ├── friends
    │   │   │   │   ├── friendsFeed.tsx
    │   │   │   │   └── index.tsx
    │   │   │   ├── group
    │   │   │   │   ├── createevent.tsx
    │   │   │   │   ├── creatorsrequest.tsx
    │   │   │   │   ├── details.tsx
    │   │   │   │   ├── eventsfeed.tsx
    │   │   │   │   ├── groupcontacts.tsx
    │   │   │   │   └── ...
    │   │   │   ├── groups
    │   │   │   │   ├── groupsFeed.tsx
    │   │   │   │   └── index.tsx
    │   │   │   ├── home
    │   │   │   │   └── index.tsx
    │   │   │   ├── notifications
    │   │   │   │   ├── index.tsx
    │   │   │   │   └── notificationsfeed.tsx
    │   │   │   ├── profile
    │   │   │   │   ├── followRequest.tsx
    │   │   │   │   ├── index.tsx
    │   │   │   │   ├── proFeed.tsx
    │   │   │   │   ├── profileDetails.tsx
    │   │   │   │   └── style.css
    │   │   │   └── settings
    │   │   │       └── index.tsx
    │   │   └── types
    │   │       ├── Comment
    │   │       │   └── index.tsx
    │   │       ├── friends
    │   │       │   └── index.tsx
    │   │       └── ...
    │   ├── tailwind.config.js
    │   ├── tsconfig.json
    │   ├── ui.config.json
    │   └── vite.config.ts
    ├── go.mod
    ├── go.sum
    ├── group_update.json
    ├── internal
    │   ├── database
    │   │   ├── docker-compose.yml
    │   │   ├── Dockerfile
    │   │   ├── images
    │   │   │   ├── 000001.webp
    │   │   │   ├── 000002.webp
    │   │   │   └── serial.txt
    │   │   ├── migrations
    │   │   │   ├── 000001_initial_schema.down.sql
    │   │   │   ├── 000001_initial_schema.up.sql
    │   │   │   ├── 000002_update_schema.down.sql
    │   │   │   ├── 000002_update_schema.up.sql
    │   │   │   └── ...
    │   │   ├── querys
    │   │   │   ├── chat.go
    │   │   │   ├── comment.go
    │   │   │   ├── database.go
    │   │   │   ├── event.go
    │   │   │   ├── follow.go
    │   │   │   └── ...
    │   │   └── SN.sql
    │   ├── helpers
    │   │   ├── env_loader.go
    │   │   ├── http_response.go
    │   │   ├── image.go
    │   │   └── validators.go
    │   ├── models
    │   │   ├── friends.go
    │   │   ├── group.go
    │   │   ├── notification.go
    │   │   ├── post.go
    │   │   └── user.go
    │   └── views
    │       ├── auth
    │       │   ├── handlers.go
    │       │   ├── hash.go
    │       │   ├── routes.go
    │       │   └── validaters.go
    │       ├── follow
    │       │   ├── followhandlers.go
    │       │   └── routes.go
    │       ├── friends
    │       │   ├── handlers.go
    │       │   └── routes.go
    │       ├── group
    │       │   ├── groupeventshandler.go
    │       │   ├── grouphandler.go
    │       │   ├── groupinvitationhandler.go
    │       │   ├── grouprequesthandler.go
    │       │   └── routes.go
    │       ├── middleware
    │       │   └── middleware.go
    │       ├── pic
    │       │   ├── routes.go
    │       │   └── servesImage.go
    │       ├── post
    │       │   ├── comments_handler.go
    │       │   ├── likeshandler.go
    │       │   ├── post_handler.go
    │       │   └── routes.go
    │       ├── profile
    │       │   ├── handlers.go
    │       │   └── routes.go
    │       ├── routes.go
    │       ├── session
    │       │   └── session.go
    │       └── websocket
    │           ├── chat.go
    │           ├── client.go
    │           ├── event_processor.go
    │           ├── handlers.go
    │           ├── notification_processor.go
    │           ├── routes.go
    │           └── types
    │               ├── event
    │               │   └── events.go
    │               ├── group.go
    │               ├── message.go
    │               └── user.go
    ├── README.md
    ├── script.sh
    └── social-network.postman_collection.json

```
</details>

## Instructions

-   Just sign up and join the elites!

## Additional information

-   The backend of the project is written in Go.
-   The frontend is built using the SolidJS framework and TypeScript.
-   PostgreSQL is used as the database.

## Authors

-   [Eman](https://github.com/emahfoodh)
-   [Amjad](https://github.com/amali01)
-   [Ahmed](https://github.com/AhmedAlAli9402)
-   [Sameer](https://github.com/sahmedG)
-   [Natheer](https://github.com/extbh)
-   [Mohammed](https://github.com/MSK17A)
