REM execute "go install .", trying upto 5 times if it fails
set count=5
:DoWhile
    if %count%==0 goto EndDoWhile
    set /a count = %count% -1
    call go install .
    if %errorlevel%==0 goto EndDoWhile
    if %count% gtr 0 goto DoWhile
:EndDoWhile