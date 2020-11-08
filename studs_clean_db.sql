--
-- PostgreSQL database dump
--

-- Dumped from database version 12.4
-- Dumped by pg_dump version 12.4

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

DROP INDEX public.idx_study_sessions_deleted_at;
DROP INDEX public.idx_profiles_deleted_at;
ALTER TABLE ONLY public.study_sessions DROP CONSTRAINT study_sessions_pkey;
ALTER TABLE ONLY public.profiles DROP CONSTRAINT profiles_pkey;
ALTER TABLE public.study_sessions ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.profiles ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE public.study_sessions_id_seq;
DROP TABLE public.study_sessions;
DROP SEQUENCE public.profiles_id_seq;
DROP TABLE public.profiles;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: profiles; Type: TABLE; Schema: public; Owner: studuser
--

CREATE TABLE public.profiles (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text,
    last_name text,
    authenticator_id text,
    major text,
    school text,
    studs character varying(64)[],
    picture_url text
);


ALTER TABLE public.profiles OWNER TO studuser;

--
-- Name: profiles_id_seq; Type: SEQUENCE; Schema: public; Owner: studuser
--

CREATE SEQUENCE public.profiles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.profiles_id_seq OWNER TO studuser;

--
-- Name: profiles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: studuser
--

ALTER SEQUENCE public.profiles_id_seq OWNED BY public.profiles.id;


--
-- Name: study_sessions; Type: TABLE; Schema: public; Owner: studuser
--

CREATE TABLE public.study_sessions (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    subject text,
    status text,
    lattitude numeric,
    longitude numeric,
    location text,
    buds character varying(64)[],
    duration text,
    is_private boolean,
    profile_id integer,
    scheduled_date text
);


ALTER TABLE public.study_sessions OWNER TO studuser;

--
-- Name: study_sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: studuser
--

CREATE SEQUENCE public.study_sessions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.study_sessions_id_seq OWNER TO studuser;

--
-- Name: study_sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: studuser
--

ALTER SEQUENCE public.study_sessions_id_seq OWNED BY public.study_sessions.id;


--
-- Name: profiles id; Type: DEFAULT; Schema: public; Owner: studuser
--

ALTER TABLE ONLY public.profiles ALTER COLUMN id SET DEFAULT nextval('public.profiles_id_seq'::regclass);


--
-- Name: study_sessions id; Type: DEFAULT; Schema: public; Owner: studuser
--

ALTER TABLE ONLY public.study_sessions ALTER COLUMN id SET DEFAULT nextval('public.study_sessions_id_seq'::regclass);


--
-- Data for Name: profiles; Type: TABLE DATA; Schema: public; Owner: studuser
--

COPY public.profiles (id, created_at, updated_at, deleted_at, first_name, last_name, authenticator_id, major, school, studs, picture_url) FROM stdin;
1	2020-11-07 19:09:17.015917-08	2020-11-07 19:09:17.015917-08	\N	Joshua	Magdaleno	2039193	Computer Science	CSUN	{Math,Physics}	\N
2	2020-11-07 19:31:16.861643-08	2020-11-07 19:31:16.861643-08	\N	Kaya	Mun	5839193	Computer Science	Northridge	{Math,English}	\N
3	2020-11-07 19:31:28.518639-08	2020-11-07 19:31:28.518639-08	\N	Deion	Shallenberger	9130192	Mechanical Engineer	UCLA	{English,Physics}	\N
4	2020-11-07 19:59:16.138444-08	2020-11-07 19:59:16.138444-08	\N	Chris	Von	42901	Computer Engineer	John Hopkins	{English,Physics,"Computer Science"}	\N
\.


--
-- Data for Name: study_sessions; Type: TABLE DATA; Schema: public; Owner: studuser
--

COPY public.study_sessions (id, created_at, updated_at, deleted_at, title, subject, status, lattitude, longitude, location, buds, duration, is_private, profile_id, scheduled_date) FROM stdin;
6	2020-11-08 02:10:51.12901-08	2020-11-08 02:10:51.12901-08	\N	Chemistry 102 Acids & Bases	Chemistry	Completed	34.27268	-118.50169	Trader Joes	{Tom,Tim,Bryan}	2 miles	f	3	11/04/2020
1	2020-11-07 22:46:23.143771-08	2020-11-07 22:46:23.143771-08	\N	Data Structures Exam	Computer Science	In Progress	34.27282	-118.50215	Chop Shop	{John,Bob}	5 miles	f	1	11/20/2020
2	2020-11-07 22:54:56.792782-08	2020-11-07 22:54:56.792782-08	\N	Calculus 2 Series Exam	Math	Completed	34.272873	-118.5037	Target	{Peter,Heather,Fernie}	8 miles	f	1	11/03/2020
3	2020-11-07 22:56:11.612009-08	2020-11-07 22:56:11.612009-08	\N	Chemistry 102 Acids & Bases	Chemistry	Completed	34.27268	-118.50169	Trader Joes	{Bonnie,Chloe,Bryan,Ronnie}	2 miles	f	1	11/01/2020
4	2020-11-07 22:58:11.565889-08	2020-11-07 22:58:11.565889-08	\N	Chemistry 102 Acids & Bases	Chemistry	Completed	34.27268	-118.50169	Trader Joes	{Zack,Zach,Zak}	2 miles	f	2	11/02/2020
5	2020-11-07 22:58:48.590981-08	2020-11-07 22:58:48.590981-08	\N	Chemistry 102 Acids & Bases	Chemistry	Completed	34.27268	-118.50169	Trader Joes	{Maria,Peter,Deion}	2 miles	f	3	11/02/2020
\.


--
-- Name: profiles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: studuser
--

SELECT pg_catalog.setval('public.profiles_id_seq', 4, true);


--
-- Name: study_sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: studuser
--

SELECT pg_catalog.setval('public.study_sessions_id_seq', 6, true);


--
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: studuser
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);


--
-- Name: study_sessions study_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: studuser
--

ALTER TABLE ONLY public.study_sessions
    ADD CONSTRAINT study_sessions_pkey PRIMARY KEY (id);


--
-- Name: idx_profiles_deleted_at; Type: INDEX; Schema: public; Owner: studuser
--

CREATE INDEX idx_profiles_deleted_at ON public.profiles USING btree (deleted_at);


--
-- Name: idx_study_sessions_deleted_at; Type: INDEX; Schema: public; Owner: studuser
--

CREATE INDEX idx_study_sessions_deleted_at ON public.study_sessions USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

