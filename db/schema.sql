-- 
-- Name: user_role; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.user_role AS ENUM (
    'customer',
    'admin'
);

--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email text,
    username text NOT NULL,
    role public.user_role DEFAULT 'user'::public.user_role,
    password text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

--
-- Name: expenses; Type: TABLE; Schema: public; Owner: -
-- 

CREATE TABLE public.expenses (
    id integer NOT NULL,
    user_id integer NOT NULL,
    description text NOT NULL,
    amount numeric(10, 2) NOT NULL,
    date timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);