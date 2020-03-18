# Abort on Container Exit VS Remove Orphans

## Background information

* _--abort-on-container-exit_ Stops all containers if any container was stopped. Incompatible with -d
* _--remove-orphans_ Remove containers for services not defined in the Compose file.

## Assumptions and Questions

* I believe the _--abort-on-container-exit_ was introduced because the database persisted in between tests. The reason for this is because the db container was not shutdown when a test errored out. By aborting the db container when the test container aborted, we can ensure a fresh test every time.
* After everything with docker-compose down is complete, is the final effect of abort on container exit and remove orphans the same?
* Do all the containers get removed in the end?

## Setup

* Have three containers running, Fluorite, and two versions of Garnet
* Fluorite runs forever (simulate or run a server)
* The Garnet containers do tasks and finish (simulate a test that calls the server)
* 