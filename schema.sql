--
-- PostgreSQL database dump
--

-- Dumped from database version 14.16 (Homebrew)
-- Dumped by pg_dump version 14.16 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

SET default_tablespace = '';
SET default_table_access_method = heap;

--
-- Drop existing tables if they exist
--

DROP TABLE IF EXISTS public.group_members CASCADE;
DROP TABLE IF EXISTS public.posts CASCADE;
DROP TABLE IF EXISTS public.votes CASCADE;
DROP TABLE IF EXISTS public.groups CASCADE;
DROP TABLE IF EXISTS public.users CASCADE;

--
-- Name: users; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name text,
    email text,
    password text,
    PRIMARY KEY (id),
    UNIQUE (email)
);

ALTER TABLE public.users OWNER TO agoravotedb;

--
-- Name: groups; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.groups (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name text NOT NULL,
    description text,
    picture text,
    is_private boolean NOT NULL,
    last_active text NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (id)
);

ALTER TABLE public.groups OWNER TO agoravotedb;

--
-- Name: group_members; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.group_members (
    group_id uuid NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4(),
    user_id uuid NOT NULL,
    role boolean,
    created_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (group_id) REFERENCES public.groups(id),
    FOREIGN KEY (user_id) REFERENCES public.users(id)
);

ALTER TABLE public.group_members OWNER TO agoravotedb;

--
-- Name: posts; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.posts (
    id bigint NOT NULL DEFAULT nextval('public.posts_id_seq'::regclass),
    group_id uuid NOT NULL,
    user_id uuid NOT NULL,
    content text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (group_id) REFERENCES public.groups(id),
    FOREIGN KEY (user_id) REFERENCES public.users(id)
);

ALTER TABLE public.posts OWNER TO agoravotedb;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: agoravotedb
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE public.posts_id_seq OWNER TO agoravotedb;

--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: agoravotedb
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;

--
-- Name: votes; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.votes (
    id bigint NOT NULL DEFAULT nextval('public.votes_id_seq'::regclass),
    group_id uuid NOT NULL,
    user_id uuid NOT NULL,
    value text,
    created_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (group_id) REFERENCES public.groups(id),
    FOREIGN KEY (user_id) REFERENCES public.users(id)
);

ALTER TABLE public.votes OWNER TO agoravotedb;

--
-- Name: votes_id_seq; Type: SEQUENCE; Schema: public; Owner: agoravotedb
--

CREATE SEQUENCE public.votes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE public.votes_id_seq OWNER TO agoravotedb;

--
-- Name: votes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: agoravotedb
--

ALTER SEQUENCE public.votes_id_seq OWNED BY public.votes.id;

--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);

--
-- Name: votes id; Type: DEFAULT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.votes ALTER COLUMN id SET DEFAULT nextval('public.votes_id_seq'::regclass);

--
-- Name: groups groups_id_unique; Type: CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_id_unique UNIQUE (id);

--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);

--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);

--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

--
-- Name: votes votes_pkey; Type: CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT votes_pkey PRIMARY KEY (id);

--
-- Name: group_members fk_groups_members; Type: FK CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.group_members
    ADD CONSTRAINT fk_groups_members FOREIGN KEY (group_id) REFERENCES public.groups(id);

--
-- Name: posts fk_posts_groups; Type: FK CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_posts_groups FOREIGN KEY (group_id) REFERENCES public.groups(id);

--
-- Name: posts fk_posts_users; Type: FK CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_posts_users FOREIGN KEY (user_id) REFERENCES public.users(id);

--
-- Name: votes fk_votes_groups; Type: FK CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT fk_votes_groups FOREIGN KEY (group_id) REFERENCES public.groups(id);

--
-- Name: votes fk_votes_users; Type: FK CONSTRAINT; Schema: public; Owner: agoravotedb
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT fk_votes_users FOREIGN KEY (user_id) REFERENCES public.users(id);

--
-- Alter existing tables
--

-- Example: Add a new column to an existing table
-- ALTER TABLE public.users ADD COLUMN new_column_name data_type;

-- Example: Drop a column from an existing table
-- ALTER TABLE public.users DROP COLUMN old_column_name;

-- Example: Add a foreign key constraint to an existing table
-- ALTER TABLE public.group_members ADD CONSTRAINT fk_group_members_group_id FOREIGN KEY (group_id) REFERENCES public.groups(id);

-- Example: Add a unique constraint to an existing table
-- ALTER TABLE public.users ADD CONSTRAINT unique_email UNIQUE (email);

--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: arnaudbrubacher
--

GRANT ALL ON SCHEMA public TO agoravotedb;

--
-- PostgreSQL database dump complete
--

