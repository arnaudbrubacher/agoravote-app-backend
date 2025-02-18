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
-- Name: group_members; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.group_members (
    group_id uuid NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4(),
    user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.group_members OWNER TO agoravotedb;

--
-- Name: groups; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.groups (
    name text NOT NULL,
    description text,
    picture text,
    is_private boolean NOT NULL,
    last_active text NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4()
);


ALTER TABLE public.groups OWNER TO agoravotedb;

--
-- Name: posts; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.posts (
    id bigint NOT NULL,
    group_id uuid NOT NULL,
    content text NOT NULL,
    created_at text NOT NULL
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
-- Name: users; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.users (
    id text NOT NULL,
    name text,
    email text,
    password text,
    role text
);


ALTER TABLE public.users OWNER TO agoravotedb;

--
-- Name: votes; Type: TABLE; Schema: public; Owner: agoravotedb
--

CREATE TABLE public.votes (
    id bigint NOT NULL,
    group_id uuid NOT NULL,
    value text,
    user_id text NOT NULL,
    created_at text NOT NULL
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
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: arnaudbrubacher
--

GRANT ALL ON SCHEMA public TO agoravotedb;


--
-- PostgreSQL database dump complete
--

