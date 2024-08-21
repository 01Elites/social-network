-- Up Migration: V2__add_engineer_and_alien_to_gender_type.sql

ALTER TYPE public.gender_type ADD VALUE 'engineer';
ALTER TYPE public.gender_type ADD VALUE 'alien';
