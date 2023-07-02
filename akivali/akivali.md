# Akivali (Name Pending) 

Akivali (name pending) is the all-in-one Testing Framework designed for SRSIM.

This framework is designed to be event-driven. We expect all important actions in the sim 
be traceable via Events. If a feature cannot be verified by Akivali, it cannot be audited via 
logs. 

The Aeon of Trailblaze (and Kyle, probably) has demanded that every Lightcone and Characters 
added to the sim be complete with a complimentary suite of tests to ensure 
that his beautiful sim does not get riddled with bugs. 

## Feature

### Currently Supported Features

- Customizing and starting Simulations
- Logging simulation events for debugging verification
- Generating test scenarios to verify various things
- A Turn Dictator for specifically testing stuff (Turn manager bypass)

### Pending features

- Missing certain core pkg features for full testing support. Turn/Combat/Ult/etc.
- Missing Codegen support for defining characters (traces in particular)

### Structure

- eventchecker: Provides functions for checking Events. Add new comparators to the correct Event folder. 
- testcase: Collection of all tests. 
  - basic: testcases for the framework itself
  - lightcone: testcases for LC logic
  - character: testcases for char logic (pending)
  - scenario: testcases for complicated scenarios (pending)
- testcfg: Collection of test configurations. 
  - testchar: character sample configs
  - testcone: lightcone sample configs
  - testeval: character skill evaluators, for skill/ult logic etc
- teststub: The test framework itself. You shouldn't need to access this package as a tester. 

## How to write a test

(WIP, will complete after finalizing v0.1 framework)