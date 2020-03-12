SINGLE_FAIL_CITRINE_EXIT_CODE = docker-compose -f docker-compose-single-fail.yml up --build --abort-on-container-exit --exit-code-from citrine
SINGLE_FAIL_DRUZY_EXIT_CODE = docker-compose -f docker-compose-single-fail.yml up --build --abort-on-container-exit --exit-code-from druzy
SINGLE_FAIL_ALL_EXIT_CODE  = docker-compose -f docker-compose-single-fail.yml up --build --abort-on-container-exit
SINGLE_DELAY_FAIL_ALL_EXIT_CODE  = docker-compose -f docker-compose-single-delay-fail.yml up --build --abort-on-container-exit

DOUBLE_FAIL_CITRINE_EXIT_CODE = docker-compose -f docker-compose-double-fail.yml up --build --abort-on-container-exit --exit-code-from citrine
DOUBLE_FAIL_EUDIALYTE_EXIT_CODE = docker-compose -f docker-compose-double-fail.yml up --build --abort-on-container-exit --exit-code-from eudialyte
DOUBLE_FAIL_ALL_EXIT_CODE = docker-compose -f docker-compose-double-fail.yml up --build --abort-on-container-exit

SINGLE_FAIL_DOWN = docker-compose -f docker-compose-single-fail.yml down --remove-orphans
DOUBLE_FAIL_DOWN = docker-compose -f docker-compose-double-fail.yml down --remove-orphans

# Test 1
# Citrine fails and Druzy passes.
# The Makefile receives exit code from Citrine only.
failed-citrine-send-exit-code: 
	@$(SINGLE_FAIL_CITRINE_EXIT_CODE)
	@$(SINGLE_FAIL_DOWN)

# Test 2
# Citrine fails and Druzy passes.
# Makefile receives exit code from Druzy only.
passing-druzy-send-exit-code: 
	@$(SINGLE_FAIL_DRUZY_EXIT_CODE)
	@$(SINGLE_FAIL_DOWN)

# Test 3
# Citrine fails and Druzy passes.
# No service specified for exit code.
citrine-fail-druzy-pass: 
	@$(SINGLE_FAIL_ALL_EXIT_CODE)
	@$(SINGLE_FAIL_DOWN)

# Test 4
# Citrine fails with error 127. Eudialyte fails after 3 second delay with exit 1.
# Makefile retrieves error code from Citrine only.
two-fails-citrine-sends-exit-code: 
	@$(DOUBLE_FAIL_CITRINE_EXIT_CODE)
	@$(DOUBLE_FAIL_DOWN)

# Test 5
# Citrine fails with error 127. Eudialyte fails after 3 second delay with exit 1.
# Makefile only gets exit code from Eudialyte.
two-fails-delayed-one-sends-exit-code: 
	@$(DOUBLE_FAIL_EUDIALYTE_EXIT_CODE)
	@$(DOUBLE_FAIL_DOWN)

# Test 6
# Citrine fails with error 127. Eudialyte fails after 3 second delay with exit 1.
# Makefile does not specify where to get exit code.
two-fails: 
	@$(DOUBLE_FAIL_ALL_EXIT_CODE)
	@$(DOUBLE_FAIL_DOWN)

# Druzy passes. Eudialyte fails after 3 second delay.
# Make file does not specify where to get exit code.
delay-failed:
	@$(SINGLE_DELAY_FAIL_ALL_EXIT_CODE)
	@$(SINGLE_FAIL_DOWN)
