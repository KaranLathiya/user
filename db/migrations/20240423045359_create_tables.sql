-- migrate:up

CREATE TYPE IF NOT EXISTS public.event_type AS ENUM (
	'signup',
	'login',
	'google_login',
	'organization_delete');

CREATE TYPE IF NOT EXISTS public.privacy AS ENUM (
	'public',
	'private');

CREATE TYPE IF NOT EXISTS public.signup_mode AS ENUM (
	'email',
	'google_login',
	'phone_number');

CREATE TABLE IF NOT EXISTS public.otp (
	phone_number VARCHAR(15) NULL,
	email VARCHAR(320) NULL,
	otp VARCHAR NOT NULL,
	expires_at TIMESTAMP NOT NULL,
	organization_id VARCHAR(20) NULL,
	id VARCHAR(20) NOT NULL DEFAULT unique_rowid(),
	event_type "user".public.event_type NOT NULL,
	country_code VARCHAR(5) NULL,
	CONSTRAINT pk_otp_id PRIMARY KEY (id ASC),
	CONSTRAINT cc_otp_email_and_phone_number CHECK (((email IS NULL) OR (phone_number IS NULL)) AND (NOT ((email IS NULL) AND (phone_number IS NULL))))
);

CREATE TABLE IF NOT EXISTS public.users (
	id VARCHAR(20) NOT NULL DEFAULT unique_rowid(),
	firstname VARCHAR(50) NOT NULL,
	lastname VARCHAR(50) NOT NULL,
	fullname VARCHAR(101) NOT NULL,
	email VARCHAR(320) NULL,
	username VARCHAR(30) NOT NULL,
	phone_number VARCHAR(15) NULL,
	country_code VARCHAR(5) NULL,
	created_at TIMESTAMP NOT NULL DEFAULT current_timestamp():::TIMESTAMP,
	updated_at TIMESTAMP NULL,
	privacy "user".public.privacy NOT NULL DEFAULT 'public':::"user".public.privacy,
	signup_mode "user".public.signup_mode NULL,
	CONSTRAINT pk_user_id PRIMARY KEY (id ASC),
	UNIQUE INDEX uk_users_email (email ASC),
	UNIQUE INDEX uk_users_phone_number (phone_number ASC),
	UNIQUE INDEX uk_users_username (username ASC),
	CONSTRAINT cc_users_email_and_phone_number CHECK (((email IS NULL) OR (phone_number IS NULL)) AND (NOT ((phone_number IS NULL) AND (email IS NULL))))
);

CREATE TABLE IF NOT EXISTS public.blocked_user (
	id VARCHAR(20) NOT NULL DEFAULT unique_rowid(),
	blocker VARCHAR(20) NOT NULL,
	blocked VARCHAR(20) NOT NULL,
	blocked_at TIMESTAMP NOT NULL DEFAULT current_timestamp():::TIMESTAMP,
	CONSTRAINT pk_blocked_user_id PRIMARY KEY (id ASC),
	CONSTRAINT fk_blocked_user_blocker FOREIGN KEY (blocker) REFERENCES public.users(id),
	CONSTRAINT fk_blocked_user_blocked FOREIGN KEY (blocked) REFERENCES public.users(id),
	UNIQUE INDEX uk_blocked_user_blocker_and_blocked (blocker ASC, blocked ASC)
);

-- migrate:down

DROP TABLE IF EXISTS public.otp;

DROP TABLE IF EXISTS public.users;

DROP TABLE IF EXISTS public.blocked_user;
