/*
Package pocketlog exposes an API to your log work.

First, instantiate a logger with pocketlog.New giving it a threshold level.
Messages of lesser criticality won't be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to log messages on four levels:
  - Debug: mostly used to debug code, follow step-by-step processes
	- Info: valuable messages providing  insights to the milestones of a process
	- Warn: valuable messages providing  insights to the milestones of a process and warnings
  - Error: error messages to understand what went wrong
*/
package pocketlog