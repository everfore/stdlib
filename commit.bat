git add -A
git status
::echo commit for change : 
@echo commit:
@set /p commit_for_change=
git commit -m "%commit_for_change%"
::net stop server
ping /n 8 127.1 >nul
::net start server