# Docker Exit Codes

## Setup and Questions

* Create three docker containers, Citrine, Druzy, and Eudialyte
* Spin up containers using docker-compose
* Have Citrine error out with error code 127 after 3 second delay. Druzy has no errors. What does the exit code message look like?
* Have Citrine and Eudialyte both error out. Citrine errors out with 127 after 3 second delay. Eudialyte does exit 1 with no deplay. What does the exit code message look like?
* How does a makefile receive these errors?

## Observations

* Test1: Citrine errors with code 127. Druzy doesn't error. Flag set is --exit-code-from citrine
  ```
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  make: *** [failed-citrine-send-exit-code] Error 127
  ```
  Citrine causes failure and container abort. Makefile receives Citrine's error code.

* Test2: Citrine errors with code 127. Druzy doesn't error. Flag set is --exit-code-from druzy
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
  We wanted exit code from Druzy, which did not error out. So, no error code was caught, even though Citrine errored out first.

* Test3: Citrine errors with code 127. Druzy doesn't error. No exit-code-from was set.
  ```
  citrine_1  | ./bashFail.sh: line 5: gibberish: command not found
  druzy_1    | /app/druzy_scripts:
  druzy_1    | 14 bashPass.sh
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  make: *** [citrine-fail-druzy-pass] Error 127
  ```
  Even though no exit code flag was set, make received Citrine's error. Was it because it caused the containers to abort?

* Test4: Citrine errors with code 127. Eudialyte sends exit 1 after 3 second delay. Flag set is --exit-code-from citrine
  ```
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  Stopping docker-exit-tests_eudialyte_1 ... done
  make: *** [two-fails-citrine-sends-exit-code] Error 127
  ```
  Citrine's exit caused container exit and make receives Citrine's error.

* Test5: Citrine errors with code 127. Eudialyte sends exit 1 after 3 second delay. Flag set is --exit-code-from eudialyte
  ```
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  Stopping docker-exit-tests_eudialyte_1 ... done
  make: *** [two-fails-delayed-one-sends-exit-code] Error 1
  ```
  Eudialyte's exit code was requested, so make received its error, even if it wasn't the cause of container abort.

* Test 6: Citrine errors with code 127. Eudialyte sends exit 1 after 3 second delay. No exit-code-from was set.
  ```
  docker-exit-tests_citrine_1 exited with code 127
  Aborting on container exit...
  Stopping docker-exit-tests_eudialyte_1 ... done
  make: *** [two-fails] Error 127
  ```
  Citrine's exit code was caught by make, even though none were specified. Could it be because its exit caused the container to abort?

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

## Conclusion
The error code that gets passed on to make is either (1) the container that causes "Abort on container exit..." or (2) the service container that was specified with the flag "--exit-code-from". 

If both happen by two different containers (see Tests #2 and #5), then the flagged service container takes priority. The system will wait for the container to send an error, if there is one. If no error, then the error code passed on is empty.

