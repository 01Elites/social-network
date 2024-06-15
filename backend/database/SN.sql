CREATE SCHEMA IF NOT EXISTS "public";

CREATE TYPE public.user_type AS ENUM ('private','public');

CREATE TYPE public.status_type AS ENUM ('pending', 'accepted', 'rejected');

CREATE TYPE post_privacy AS ENUM ('public', 'private', 'almost_private', 'group');

CREATE TYPE public.role_type AS ENUM ('admin', 'member');

CREATE TYPE public.chat_type AS ENUM ('private', 'group');

CREATE TYPE public.notification_type AS ENUM (
    'follow_request',
    'group_invite',
    'join_request',
    'event_notification',
    'post_notification',
    'comment_notification'
);

CREATE TABLE public.user (
    user_id        SERIAL PRIMARY KEY,
    user_name      VARCHAR(100),
    email          VARCHAR NOT NULL,
    "password"     VARCHAR NOT NULL,
    first_name     VARCHAR(100),
    last_name      VARCHAR(100),
    gender         VARCHAR,
    date_of_birth  DATE,
    image          VARCHAR,
    session_uuid   UUID,
    user_type      public.user_type DEFAULT 'public'::user_type NOT NULL,
    CONSTRAINT unq_tbl UNIQUE (user_name, email)
);
 
CREATE TABLE public.follower (
    follower_id    INTEGER NOT NULL,
    followed_id    INTEGER NOT NULL,
    PRIMARY KEY (follower_id, followed_id),
    CONSTRAINT fk_follower_user FOREIGN KEY (follower_id) REFERENCES public.user (user_id) ON DELETE CASCADE,
    CONSTRAINT fk_follower_user_followed FOREIGN KEY (followed_id) REFERENCES public.user (user_id) ON DELETE CASCADE
);

CREATE TABLE public.Follow_Requests (
    request_id    SERIAL PRIMARY KEY,
    sender_id     INTEGER NOT NULL,
    receiver_id   INTEGER NOT NULL,
    status        public.status_type NOT NULL DEFAULT 'pending',
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES public.user (user_id),
    FOREIGN KEY (receiver_id) REFERENCES public.user (user_id)
);

CREATE TABLE public.group (
    group_id     SERIAL PRIMARY KEY,
    title        VARCHAR(255),
    description  TEXT,
    creator_id   INTEGER NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (creator_id) REFERENCES public.user (user_id)
);

CREATE TABLE public.group_invitations (
    invitation_id  SERIAL PRIMARY KEY,
    group_id       INTEGER NOT NULL,
    sender_id      INTEGER NOT NULL,
    receiver_id    INTEGER NOT NULL,
    status         public.status_type NOT NULL DEFAULT 'pending',
    sent_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (group_id) REFERENCES public.group (group_id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES public.user (user_id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES public.user (user_id) ON DELETE CASCADE
);

CREATE TABLE public.post (
    post_id       SERIAL PRIMARY KEY,
    title         VARCHAR(255),
    content       VARCHAR(255) NOT NULL,
    privacy_type  post_privacy NOT NULL DEFAULT 'public',
    user_id       INTEGER NOT NULL,
    image         VARCHAR(255),
    group_id      INTEGER,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.user(user_id),
    FOREIGN KEY (group_id) REFERENCES public.group(group_id)
);

CREATE TABLE public.post_user (
    post_id          INTEGER NOT NULL,
    allowed_user_id  INTEGER NOT NULL,
    FOREIGN KEY (post_id) REFERENCES post(post_id),
    FOREIGN KEY (allowed_user_id) REFERENCES public.user(user_id)
);

CREATE TABLE public.comment (
    comment_id   SERIAL PRIMARY KEY,
    post_id      INTEGER NOT NULL,
    user_id      INTEGER NOT NULL,
    content      VARCHAR(255) NOT NULL,
    image        VARCHAR(255), 
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES public.post(post_id),
    FOREIGN KEY (user_id) REFERENCES public.user(user_id)
);

CREATE TABLE public.chat (
    chat_id       SERIAL PRIMARY KEY,
    chat_type     public.chat_type NOT NULL,
    group_id      INTEGER,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (group_id) REFERENCES public.group (group_id) ON DELETE SET NULL
);

CREATE TABLE public.participant (
    user_id       INTEGER NOT NULL,
    chat_id       INTEGER NOT NULL,
    role          public.role_type NOT NULL,
    joined_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (chat_id, user_id),
    FOREIGN KEY (chat_id) REFERENCES public.chat (chat_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES public.user (user_id) ON DELETE CASCADE
);

CREATE TABLE public.group_requests (
    request_id     SERIAL PRIMARY KEY,
    group_id       INTEGER NOT NULL,
    requester_id   INTEGER NOT NULL,
    status         public.status_type NOT NULL DEFAULT 'pending',
    requested_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (group_id) REFERENCES public.group (group_id) ON DELETE CASCADE,
    FOREIGN KEY (requester_id) REFERENCES public.user (user_id) ON DELETE CASCADE
);

CREATE TABLE public.event (
    event_id     SERIAL PRIMARY KEY,
    group_id     INTEGER NOT NULL,
    title        VARCHAR(255) NOT NULL,
    description  TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (group_id) REFERENCES public.group (group_id) ON DELETE CASCADE
);

CREATE TABLE public.event_option (
    option_id    SERIAL PRIMARY KEY,
    event_id     INTEGER NOT NULL,
    name         VARCHAR(255) NOT NULL,
    FOREIGN KEY (event_id) REFERENCES public.event (event_id) ON DELETE CASCADE
);

CREATE TABLE public.user_choice (
    event_id     INTEGER NOT NULL,
    user_id      INTEGER NOT NULL,
    option_id    INTEGER NOT NULL,
    PRIMARY KEY (event_id, user_id),
    FOREIGN KEY (event_id) REFERENCES public.event (event_id),
    FOREIGN KEY (user_id) REFERENCES public.user (user_id),
    FOREIGN KEY (option_id) REFERENCES public.event_option (option_id)
);

CREATE TABLE public.notifications (
    notification_id  SERIAL PRIMARY KEY,
    user_id          INTEGER NOT NULL,
    type             public.notification_type NOT NULL,
    status           public.status_type,
    related_id       INTEGER,  -- This can store IDs related to the notification (e.g., group_id, event_id)
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    read             BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES public.user (user_id) ON DELETE CASCADE
);



