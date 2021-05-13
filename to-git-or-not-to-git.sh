#!/bin/bash
curl -s https://ytrack.learn.ynov.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"ABOYER1\"}}){id}}"}' | tr -dc '0-9'