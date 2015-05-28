git add -A
git status
::echo commit for change : 
@echo commit:
@set /p commit_for_change=
git commit -m "%commit_for_change%"
sleep 4