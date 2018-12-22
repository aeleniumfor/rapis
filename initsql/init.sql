CREATE DATABASE domain;

\c domain

CREATE TABLE hosts (
  id                 SERIAL PRIMARY KEY NOT NULL,
  project_name       VARCHAR(256)       NOT NULL UNIQUE,
  r_urls             VARCHAR(256)       NOT NULL UNIQUE,
  t_urls             VARCHAR(256)       NOT NULL UNIQUE,
  created_at         TIMESTAMP          NOT NULL DEFAULT current_timestamp,
  updated_at         TIMESTAMP          NOT NULL DEFAULT current_timestamp
);
