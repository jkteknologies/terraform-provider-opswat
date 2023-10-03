CALL :try 5 "go install ."
GOTO :EOF


:try
SET /A tries=%1

:loop
IF %tries% LEQ 0 GOTO return

SET /A tries-=1
EVAL %2 && (GOTO return) || (GOTO loop)

:return
EXIT /B