while [ "$(docker inspect -f {{.State.Running}} chat-postgresql)" != "true" ]; do sleep 2; done
