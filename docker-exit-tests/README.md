# Docker Exit Codes

## Setup and Questions

* Create three docker containers, Citrine, Druzy, and Eudialyte
* Spin up containers using docker-compose
* Have Citrine error out with error code 127 after delay. Druzy has no errors. What does the exit code message look like?
* Have Citrine and Eudialyte both error out. Citrine errors out with 127 right away. Eudialyte does exit 1 with a delay. What does the exit code message look like?
* How does a makefile receive these errors?

## Observations

* Test 1: Citrine errors with code 127. Druzy doesn't error. Flag set is --exit-code-from citrine
  ```
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  make: *** [failed-citrine-send-exit-code] Error 127
  ```
  Citrine causes failure and container abort. Makefile receives Citrine's error code.

* Test 2: Citrine errors with code 127. Druzy doesn't error. Flag set is --exit-code-from druzy
  ```
  citrine_1  | ./bashFail.sh: line 5: gibberish: command not found
  druzy_1    | Hi root. Here are the files and their line counts in pwd,
  druzy_1    | /app/druzy_scripts:
  druzy_1    | 14 bashPass.sh
  docker-exit-tests_druzy_1 exited with code 0
  Aborting on container exit...
  Removing docker-exit-tests_citrine_1 ... done
  Removing docker-exit-tests_druzy_1   ... done
  Removing network docker-exit-tests_default
  ```
  We wanted exit code from Druzy, which did not error out. So, no error code was caught (well, the error caught was exit 0), even though Citrine errored out first.

* Test 3: Citrine errors with code 127. Druzy doesn't error. No exit-code-from was set.
  ```
  citrine_1  | ./bashFail.sh: line 5: gibberish: command not found
  druzy_1    | /app/druzy_scripts:
  druzy_1    | 14 bashPass.sh
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  make: *** [citrine-fail-druzy-pass] Error 127
  ```
  This test can have two different outputs, depending on which service exits first. Since no exit-code-from was set, the exit code comes from the container that causes the abort. The above occurs if Citrine aborts first. The below occcurs if Druzy exits first (and sends over an exit code of 0):
  ```
  citrine_1  | ./bashFail.sh: line 5: gibberish: command not found
  druzy_1    | /app/druzy_scripts:
  druzy_1    | 14 bashPass.sh
  docker-exit-tests_druzy_1 exited with code 0
  Aborting on container exit...
  Removing docker-exit-tests_druzy_1   ... done
  Removing docker-exit-tests_citrine_1 ... done
  Removing network docker-exit-tests_default
  ```

* Test 4: Citrine errors with code 127. Eudialyte sends exit 1 or 137, depending on length of delay. Flag set is --exit-code-from citrine
  ```
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  Stopping docker-exit-tests_eudialyte_1 ... done
  make: *** [two-fails-citrine-sends-exit-code] Error 127
  ```
  Citrine's exit caused container exit and make receives Citrine's error.

* Test 5: Citrine errors with code 127. Eudialyte sends exit 1 or 137, depending on length of delay. Flag set is --exit-code-from eudialyte
  ```
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  Stopping docker-exit-tests_eudialyte_1 ... done
  make: *** [two-fails-delayed-one-sends-exit-code] Error 137
  ```
  Eudialyte's exit code was requested, so make received its error, even if it wasn't the cause of container abort.

* Test 6: Citrine errors with code 127. Eudialyte sends exit 1 or 137, depending on length of delay. No exit-code-from was set.
  ```
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  Stopping docker-exit-tests_eudialyte_1 ... done
  make: *** [two-fails] Error 127
  ```
  Citrine's exit code was caught by make, even though none were specified, but it was the cause of the container abort.

* Test 7: Druzy doesn't error out, while Eudialyte errors out after a delay. No exit-code-from was set.
  ```
  docker-exit-tests_druzy_1 exited with code 0
  Aborting on container exit...
  Stopping docker-exit-tests_eudialyte_1 ... done
  Removing orphan container "docker-exit-tests_eudialyte_1"
  Removing docker-exit-tests_citrine_1 ... done
  Removing docker-exit-tests_druzy_1   ... done
  Removing network docker-exit-tests_default
  ```
  No error code was received because the container that caused the abort did not fail, and no service was specified from which to get the error.

  * Special test. I added a long delay (25 s) in separate command, in addition to the 3 second one already in the bashDelayFail.sh file. Strangely, Eudialyte's exit code 1 is no longer received by Test #5, even though Eudialyte exit code  was specified. I see exit code 137. Before the addition, any time Eudialyte's exit code took priority, I saw exit code 1.

  The reason for this is explained in Docker [docs](https://success.docker.com/article/what-causes-a-container-to-exit-with-code-137):
  
  > The container received a docker stop, and the application didn't gracefully handle SIGTERM (kill -15) — whenever a SIGTERM has been issued, the docker daemon waits 10 seconds then issue a SIGKILL (kill -9) to guarantee the shutdown. To test whether your containerized application correctly handles SIGTERM, simply issue a docker stop against the container ID and check to see whether you get the "task: non-zero exit (137)"

  What that means is, Eudialyte is exiting with error 137 after 10 seconds of the first container exit. That is the exit code caught if it was the service from which the exit code would be from.

## Conclusion
The error code that gets passed on to make is either (1) the container that causes "Abort on container exit..." or (2) the service container that was specified with the flag "--exit-code-from".

If the first container to exit exited with no error, it appears there was no error because the exit code was 0.

If one container exited, but the other was specified to be where the exit code was from (see Tests #2 and #5), then the flagged service container takes priority. The system will wait for the container to send an error, if there is one. If no error, then the error code passed on is empty. The length of time for the second container to error out and send its error will be 10 seconds (as that is the amount of time for SIGTERM sent by Docker takes). Otherwise, the error code 137 will be sent by the late container.
